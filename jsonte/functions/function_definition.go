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
	Value      reflect.Value
	WithError  bool
	IsInstance bool
	IsUnsafe   bool
	IsVarArgs  bool
	Deprecated bool
	Docs       Docs

	// argMatchers are precomputed per-parameter type checks used by checkParams.
	argMatchers []argMatcher
	// invoke is a typed, reflection-free invoker for common signatures. It is nil when the
	// signature is not covered, in which case Value.Call is used.
	invoke invoker
}

type invoker = func(args []types.JsonType) (types.JsonType, error)

// argMatcher checks whether a (non-nil) argument can be passed to a parameter.
type argMatcher struct {
	typ reflect.Type
	// any is set when the parameter is types.JsonType, which every argument implements.
	any bool
	// exactOnly is set when only the identical type is assignable (unnamed pointer types).
	exactOnly bool
}

func (m *argMatcher) matches(arg types.JsonType) bool {
	if m.any {
		return true
	}
	t := reflect.TypeOf(arg)
	if t == m.typ {
		return true
	}
	if m.exactOnly {
		return false
	}
	return t.AssignableTo(m.typ)
}

// functionTable holds overloads keyed by lower-cased name, plus an index by every exact
// spelling that was registered, so that lookups can usually skip strings.ToLower.
type functionTable struct {
	lower    map[string][]JsonFunction
	exact    map[string][]JsonFunction
	variants map[string][]string
}

func newFunctionTable(lower map[string][]JsonFunction) *functionTable {
	return &functionTable{
		lower:    lower,
		exact:    make(map[string][]JsonFunction),
		variants: make(map[string][]string),
	}
}

func (t *functionTable) add(fn JsonFunction) {
	lower := strings.ToLower(fn.Name)
	t.lower[lower] = append(t.lower[lower], fn)
	names := t.variants[lower]
	found := false
	for _, n := range names {
		if n == fn.Name {
			found = true
			break
		}
	}
	if !found {
		names = append(names, fn.Name)
		t.variants[lower] = names
	}
	// Keep every exact spelling pointing at the complete overload list.
	for _, n := range names {
		t.exact[n] = t.lower[lower]
	}
}

func (t *functionTable) lookup(name string) ([]JsonFunction, bool) {
	if fns, ok := t.exact[name]; ok {
		return fns, true
	}
	fns, ok := t.lower[strings.ToLower(name)]
	return fns, ok
}

var groups = map[string]Group{}
var functions = make(map[string][]JsonFunction)
var functionsTable = newFunctionTable(functions)
var functionNames = make([]string, 0)
var instanceFunctions = make(map[reflect.Type]*functionTable)

var jsonTypeType = reflect.TypeOf((*types.JsonType)(nil)).Elem()
var jsonTypeSliceType = reflect.TypeOf((*[]types.JsonType)(nil)).Elem()

var initialized = false

var cacheAll = false
var cacheAllBucket = "all-cache"

var SafeMode = false
var ServerMode = false

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
		RegisterScriptFunctions()
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
	fn.Value = of
	if of.Type().NumIn() == 0 && fn.IsInstance {
		utils.BadDeveloperError("Registered instance function doesn't have an instance parameter!")
	}
	for i := 0; i < of.Type().NumIn(); i++ {
		typeSanityCheck(of.Type().In(i))
		fn.Args = append(fn.Args, of.Type().In(i))
		if of.Type().In(i).AssignableTo(jsonTypeSliceType) {
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
	fn.argMatchers = make([]argMatcher, len(fn.Args))
	for i, t := range fn.Args {
		fn.argMatchers[i] = argMatcher{
			typ:       t,
			any:       t == jsonTypeType,
			exactOnly: t.Kind() == reflect.Pointer && t.Name() == "",
		}
	}
	fn.invoke = makeInvoker(fn.Body)
	functionsTable.add(fn)
	functionNames = append(functionNames, fn.Name)
	if fn.IsInstance {
		table, ok := instanceFunctions[fn.Args[0]]
		if !ok {
			table = newFunctionTable(make(map[string][]JsonFunction))
			instanceFunctions[fn.Args[0]] = table
		}
		table.add(fn)
	}
}

func typeSanityCheck(in reflect.Type) {
	if in.AssignableTo(jsonTypeType) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*types.JsonLambda)(nil)).Elem()) {
		return
	}
	if in.AssignableTo(jsonTypeSliceType) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*error)(nil)).Elem()) {
		return
	}
	utils.BadDeveloperError(fmt.Sprintf("Function body must only take parameters of types JsonType, error or JsonLambda! Got %s", in.String()))
}

func HasInstanceFunction(t reflect.Type, name string) bool {
	table, ok := instanceFunctions[t]
	if !ok {
		return false
	}
	_, ok = table.lookup(name)
	return ok
}

func HasFunction(name string) bool {
	_, ok := functionsTable.lookup(name)
	return ok
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
	var fns []JsonFunction
	ok := false
	if table, found := instanceFunctions[reflect.TypeOf(instance)]; found {
		fns, ok = table.lookup(name)
	}
	if !ok {
		find := FindMisspelling(name)
		if find != nil {
			return nil, burrito.WrappedErrorf("Instance function '%s' not found, did you mean '%s'?", name, *find)
		}
		return nil, burrito.WrappedErrorf("Instance function '%s' not found", name)
	}
	a := make([]types.JsonType, 1+len(args))
	a[0] = instance
	copy(a[1:], args)
	return callFunctionImpl(name, fns, a, resolveFunc)
}

func CallFunction(name string, args []types.JsonType, resolveFunc func(string) types.JsonType) (types.JsonType, error) {
	fns, ok := functionsTable.lookup(name)
	if !ok {
		find := FindMisspelling(name)
		if find != nil {
			return nil, burrito.WrappedErrorf("Function '%s' not found, did you mean '%s'?", name, *find)
		}
		return nil, burrito.WrappedErrorf("Function '%s' not found", name)
	}
	return callFunctionImpl(name, fns, args, resolveFunc)
}

func sizeMatches(fn *JsonFunction, args []types.JsonType) bool {
	return len(fn.Args) == len(args) || (fn.IsVarArgs && len(fn.Args)-1 <= len(args))
}

func callFunctionImpl(name string, fns []JsonFunction, args []types.JsonType, resolveFunc func(string) types.JsonType) (types.JsonType, error) {
	var match *JsonFunction
	sizeCount, matchCount := 0, 0
	for i := range fns {
		fn := &fns[i]
		if !sizeMatches(fn, args) {
			continue
		}
		sizeCount++
		if checkParams(fn, args) {
			matchCount++
			if match == nil {
				match = fn
			}
		}
	}
	if sizeCount == 0 {
		return nil, burrito.WrappedErrorf("Incorrect number of arguments for function \"%s\"", name)
	}
	if matchCount != 1 {
		sizeMatching := make([]*JsonFunction, 0, sizeCount)
		matching := make([]*JsonFunction, 0, matchCount)
		for i := range fns {
			fn := &fns[i]
			if !sizeMatches(fn, args) {
				continue
			}
			sizeMatching = append(sizeMatching, fn)
			if checkParams(fn, args) {
				matching = append(matching, fn)
			}
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
		}
		matched := make([]string, 0)
		for _, fn := range matching {
			matched = append(matched, paramsToString(args, fn.Args))
		}
		return nil, burrito.WrapErrorf(nil, "Ambiguous function call for \"%s\". Matched: %s", name, strings.Join(matched, ", "))
	}
	fn := match
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
	if fn.invoke != nil {
		result, err := fn.invoke(args)
		if err != nil {
			return nil, err
		}
		if cacheAll {
			utils.PutCache(cacheAllBucket, key, result)
		}
		return result, nil
	}
	vArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		vArgs[i] = reflect.ValueOf(arg)
	}
	call := fn.Value.Call(vArgs)
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

func checkParams(fn *JsonFunction, args []types.JsonType) bool {
	for i, arg := range args {
		if arg == nil {
			return false
		}
		if fn.IsVarArgs && i >= len(fn.Args)-1 {
			return true
		}
		if !fn.argMatchers[i].matches(arg) {
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

func GetGroups() map[string]Group {
	return groups
}

func GetFunctions() map[string][]JsonFunction {
	return functions
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
