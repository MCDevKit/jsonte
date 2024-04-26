package functions

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/paul-mannino/go-fuzzywuzzy"
	"reflect"
	"strings"
)

type JsonFunction struct {
	Group      string
	Name       string
	Args       []reflect.Type
	Body       interface{}
	WithError  bool
	IsInstance bool
	IsUnsafe   bool
	IsVarArgs  bool
	Deprecated bool
	Docs       Docs
}

var groups = map[string]Group{}
var functions = make(map[string][]JsonFunction)
var functionNames = make([]string, 0)
var instanceFunctions = make(map[string]map[string][]JsonFunction)

var initialized = false

var cacheAll = false
var cacheAllBucket = "all-cache"

var SafeMode = false

func Init() {
	if !initialized {
		RegisterMathFunctions()
		RegisterStringFunctions()
		RegisterArrayFunctions()
		RegisterImageFunctions()
		RegisterAudioFunctions()
		RegisterColorFunctions()
		RegisterMinecraftFunctions()
		RegisterFileFunctions()
		RegisterSemverFunctions()
		RegisterTypeFunctions()
		RegisterJsonPathFunctions()
		initialized = true
	}
}

func SetCacheAll(cache bool) {
	cacheAll = cache
	if cacheAll {
		utils.CreateCacheBucket(cacheAllBucket)
	}
}

func RegisterGroup(group Group) {
	groups[group.Name] = group
}

func RegisterFunction(fn JsonFunction) {
	of := reflect.ValueOf(fn.Body)
	if of.Kind() != reflect.Func {
		utils.BadDeveloperError("Function body must be a function!")
	}
	if of.Type().NumIn() == 0 && fn.IsInstance {
		utils.BadDeveloperError("Registered instance function doesn't have an instance parameter!")
	}
	for i := 0; i < of.Type().NumIn(); i++ {
		typeSanityCheck(of.Type().In(i))
		fn.Args = append(fn.Args, of.Type().In(i))
		if of.Type().In(i).AssignableTo(reflect.TypeOf((*[]types.JsonType)(nil)).Elem()) {
			fn.IsVarArgs = true
		}
	}
	if (of.Type().NumOut() > 2) || (of.Type().NumOut() == 2 && of.Type().Out(1).String() != "error") {
		utils.BadDeveloperError("Function body must return only one value and can return an error!")
	}
	for i := 0; i < of.Type().NumOut(); i++ {
		typeSanityCheck(of.Type().Out(i))
	}
	if (of.Type().NumOut() == 2) || (of.Type().NumOut() == 1 && of.Type().Out(0).String() == "error") {
		fn.WithError = true
	}
	functions[strings.ToLower(fn.Name)] = append(functions[strings.ToLower(fn.Name)], fn)
	functionNames = append(functionNames, fn.Name)
	if fn.IsInstance {
		if _, ok := instanceFunctions[fn.Args[0].String()]; !ok {
			instanceFunctions[fn.Args[0].String()] = make(map[string][]JsonFunction)
		}
		instanceFunctions[fn.Args[0].String()][strings.ToLower(fn.Name)] = append(instanceFunctions[fn.Args[0].String()][strings.ToLower(fn.Name)], fn)
	}
}

func typeSanityCheck(in reflect.Type) {
	test := reflect.TypeOf((*types.JsonType)(nil))
	if in.AssignableTo(test.Elem()) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*types.JsonLambda)(nil)).Elem()) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*[]types.JsonType)(nil)).Elem()) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*error)(nil)).Elem()) {
		return
	}
	utils.BadDeveloperError(fmt.Sprintf("Function body must only take parameters of types JsonType, error or JsonLambda! Got %s", in.String()))
}

func HasInstanceFunction(t reflect.Type, name string) bool {
	if _, ok := instanceFunctions[t.String()]; !ok {
		return false
	}
	if _, ok := instanceFunctions[t.String()][strings.ToLower(name)]; !ok {
		return false
	}
	return true
}

func HasFunction(name string) bool {
	if _, ok := functions[strings.ToLower(name)]; !ok {
		return false
	}
	return true
}

func FindMisspelling(name string) *string {
	find, err := fuzzy.Extract(name, functionNames, 5)
	if err != nil {
		return nil
	}
	if find != nil {
		for i := 0; i < find.Len(); i++ {
			if len(name)-len(find[i].Match) > 2 {
				continue
			}
			return &find[i].Match
		}
	}
	return nil
}

func CallInstanceFunction(name string, instance types.JsonType, args []types.JsonType, resolveFunc func(string) types.JsonType) (types.JsonType, error) {
	fns, ok := instanceFunctions[reflect.TypeOf(instance).String()][strings.ToLower(name)]
	if !ok {
		find := FindMisspelling(name)
		if find != nil {
			return nil, burrito.WrappedErrorf("Instance function '%s' not found, did you mean '%s'?", name, *find)
		}
		return nil, burrito.WrappedErrorf("Instance function '%s' not found", name)
	}
	a := make([]types.JsonType, 0)
	a = append(a, instance)
	a = append(a, args...)
	return callFunctionImpl(name, fns, a, resolveFunc)
}

func CallFunction(name string, args []types.JsonType, resolveFunc func(string) types.JsonType) (types.JsonType, error) {
	fns, ok := functions[strings.ToLower(name)]
	if !ok {
		find := FindMisspelling(name)
		if find != nil {
			return nil, burrito.WrappedErrorf("Function '%s' not found, did you mean '%s'?", name, *find)
		}
		return nil, burrito.WrappedErrorf("Function '%s' not found", name)
	}
	return callFunctionImpl(name, fns, args, resolveFunc)
}

func callFunctionImpl(name string, fns []JsonFunction, args []types.JsonType, resolveFunc func(string) types.JsonType) (types.JsonType, error) {
	sizeMatching := make([]JsonFunction, 0)
	for _, fn := range fns {
		if len(fn.Args) == len(args) || (fn.IsVarArgs && len(fn.Args)-1 <= len(args)) {
			sizeMatching = append(sizeMatching, fn)
		}
	}
	if len(sizeMatching) == 0 {
		return nil, burrito.WrappedErrorf("Incorrect number of arguments for function \"%s\"", name)
	}
	matching := make([]JsonFunction, 0)
	for _, fn := range sizeMatching {
		if !checkParams(fn, args) {
			continue
		}
		matching = append(matching, fn)
	}
	if len(matching) == 0 {
		expected := make([]string, 0)
		for _, fn := range sizeMatching {
			expected = append(expected, paramsToString(nil, fn.Args))
		}
		argTypes := make([]reflect.Type, 0)
		for _, arg := range args {
			argTypes = append(argTypes, reflect.TypeOf(arg))
		}
		return nil, burrito.WrappedErrorf("Incorrect argument types for function \"%s\". Expected: %s, got: %s", name, strings.Join(expected, ", "), paramsToString(args, argTypes))
	} else if len(matching) > 1 {
		matched := make([]string, 0)
		for _, fn := range matching {
			matched = append(matched, paramsToString(args, fn.Args))
		}
		return nil, burrito.WrapErrorf(nil, "Ambiguous function call for \"%s\". Matched: %s", name, strings.Join(matched, ", "))
	} else {
		fn := matching[0]
		if SafeMode && fn.IsUnsafe {
			return nil, burrito.WrappedErrorf("Function \"%s\" is marked as unsafe and is disabled in safe mode", name)
		}
		key := ""
		if cacheAll {
			keyObject := types.AsObject(utils.ToNavigableMap(
				"f", fn.Name,
				"args", &types.JsonArray{Value: args},
			))
			// Need to take care of lambdas, that use outside variables
			scope := types.NewJsonObject()
			for _, arg := range args {
				if lambda, ok := arg.(*types.JsonLambda); ok && len(lambda.UsedVariables) > 0 {
				outer:
					for _, variable := range lambda.UsedVariables {
						if lambda.Arguments != nil {
							for _, argument := range lambda.Arguments {
								if argument == variable {
									continue outer
								}
							}
						}
						scope.Put(variable, resolveFunc(variable))
					}
				}
			}
			keyObject.Put("scope", scope)
			key = keyObject.StringValue()
			cached := utils.GetCache(cacheAllBucket, key)
			if cached != nil {
				//utils.Logger.Debugf("Using cached function call: %s", key)
				return (*cached).(types.JsonType), nil
			}
		}
		vArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			vArgs[i] = reflect.ValueOf(arg)
		}
		call := reflect.ValueOf(fn.Body).Call(vArgs)
		if fn.WithError && call[len(call)-1].Interface() != nil {
			return nil, call[len(call)-1].Interface().(error)
		}
		if fn.WithError && len(call) == 1 {
			return nil, nil
		}
		if cacheAll {
			//utils.Logger.Debugf("Caching function call: %s", key)
			//utils.Logger.Debugf("Caching function call result: %s", call[0].Interface().(types.JsonType).StringValue())
			utils.PutCache(cacheAllBucket, key, call[0].Interface())
		}
		return call[0].Interface().(types.JsonType), nil
	}
}

func checkParams(fn JsonFunction, args []types.JsonType) bool {
	for i, arg := range args {
		if arg == nil {
			return false
		}
		if fn.IsVarArgs && i >= len(fn.Args)-1 {
			return true
		}
		if !reflect.TypeOf(arg).AssignableTo(fn.Args[i]) {
			return false
		}
	}
	return true
}

func paramsToString(args []types.JsonType, argTypes []reflect.Type) string {
	join := make([]string, 0)
	for i, arg := range argTypes {
		if args != nil {
			join = append(join, fmt.Sprintf("%s (%s)", types.ToString(args[i]), typeToString(arg)))
		} else {
			join = append(join, typeToString(arg))
		}
	}
	return "(" + strings.Join(join, ", ") + ")"
}

func paramsForLambda(params []interface{}) []types.JsonType {
	args := make([]types.JsonType, len(params))
	for i, param := range params {
		args[i] = types.Box(param)
	}
	return args
}

func typeToString(t reflect.Type) string {
	if t == nil {
		return "null"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonNumber{})) {
		return "number"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonString{})) {
		return "string"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonBool{})) {
		return "bool"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonArray{})) {
		return "array"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonObject{})) {
		return "object"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonNull{})) {
		return "null"
	}
	if t.AssignableTo(reflect.TypeOf(&types.Semver{})) {
		return "semver"
	}
	if t.AssignableTo(reflect.TypeOf(&types.JsonLambda{})) {
		return "semver"
	}

	return "Unsupported (report bug): " + t.Name()
}
