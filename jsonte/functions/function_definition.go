package functions

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
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
	Deprecated bool
	Docs       Docs
}

var groups = map[string]Group{}
var functions = make(map[string][]JsonFunction)
var functionNames = make([]string, 0)
var instanceFunctions = make(map[string]map[string][]JsonFunction)

var initialized = false

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
		initialized = true
	}
}

func RegisterGroup(group Group) {
	groups[group.Name] = group
}

func RegisterFunction(fn JsonFunction) {
	of := reflect.ValueOf(fn.Body)
	if of.Kind() != reflect.Func {
		panic("Function body must be a function!")
	}
	if of.Type().NumIn() == 0 && fn.IsInstance {
		panic("Registered instance function doesn't have an instance parameter!")
	}
	for i := 0; i < of.Type().NumIn(); i++ {
		typeSanityCheck(of.Type().In(i))
		fn.Args = append(fn.Args, of.Type().In(i))
	}
	if (of.Type().NumOut() > 2) || (of.Type().NumOut() == 2 && of.Type().Out(1).String() != "error") {
		panic("Function body must return only one value and can return an error!")
	}
	for i := 0; i < of.Type().NumOut(); i++ {
		typeSanityCheck(of.Type().Out(i))
	}
	if (of.Type().NumOut() == 2) || (of.Type().NumOut() == 1 && of.Type().Out(0).String() == "error") {
		fn.WithError = true
	}
	functions[fn.Name] = append(functions[fn.Name], fn)
	functionNames = append(functionNames, fn.Name)
	if fn.IsInstance {
		if _, ok := instanceFunctions[fn.Args[0].String()]; !ok {
			instanceFunctions[fn.Args[0].String()] = make(map[string][]JsonFunction)
		}
		instanceFunctions[fn.Args[0].String()][fn.Name] = append(instanceFunctions[fn.Args[0].String()][fn.Name], fn)
	}
}

func typeSanityCheck(in reflect.Type) {
	if in.AssignableTo(reflect.TypeOf((*types.JsonType)(nil)).Elem()) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*types.JsonLambda)(nil)).Elem()) {
		return
	}
	if in.AssignableTo(reflect.TypeOf((*error)(nil)).Elem()) {
		return
	}
	panic(fmt.Sprintf("Function body must only take parameters of types JsonType, error or JsonLambda! Got %s", in.String()))
}

func HasInstanceFunction(t reflect.Type, name string) bool {
	if _, ok := instanceFunctions[t.String()]; !ok {
		return false
	}
	if _, ok := instanceFunctions[t.String()][name]; !ok {
		return false
	}
	return true
}

func HasFunction(name string) bool {
	if _, ok := functions[name]; !ok {
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

func CallInstanceFunction(name string, instance types.JsonType, args []types.JsonType) (types.JsonType, error) {
	fns, ok := instanceFunctions[reflect.TypeOf(instance).String()][name]
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
	return callFunctionImpl(name, fns, a)
}

func CallFunction(name string, args []types.JsonType) (types.JsonType, error) {
	fns, ok := functions[name]
	if !ok {
		find := FindMisspelling(name)
		if find != nil {
			return nil, burrito.WrappedErrorf("Function '%s' not found, did you mean '%s'?", name, *find)
		}
		return nil, burrito.WrappedErrorf("Function '%s' not found", name)
	}
	return callFunctionImpl(name, fns, args)
}

func callFunctionImpl(name string, fns []JsonFunction, args []types.JsonType) (types.JsonType, error) {
	sizeMatching := make([]JsonFunction, 0)
	for _, fn := range fns {
		if len(fn.Args) == len(args) {
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
		return call[0].Interface().(types.JsonType), nil
	}
}

func checkParams(fn JsonFunction, args []types.JsonType) bool {
	for i, arg := range args {
		if arg == nil || !reflect.TypeOf(arg).AssignableTo(fn.Args[i]) {
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
	if t.AssignableTo(reflect.TypeOf(types.JsonNumber{})) {
		return "number"
	}
	if t.AssignableTo(reflect.TypeOf(types.JsonString{})) {
		return "string"
	}
	if t.AssignableTo(reflect.TypeOf(types.JsonBool{})) {
		return "bool"
	}
	if t.AssignableTo(reflect.TypeOf(types.JsonArray{})) {
		return "array"
	}
	if t.AssignableTo(reflect.TypeOf(types.JsonObject{})) {
		return "object"
	}
	if t.AssignableTo(reflect.TypeOf(types.JsonNull{})) {
		return "null"
	}
	if t.AssignableTo(reflect.TypeOf(types.Semver{})) {
		return "semver"
	}

	switch t.Kind() {
	//case reflect.Bool:
	//	return "boolean"
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//case reflect.Float32, reflect.Float64:
	//	return "number"
	//case reflect.String:
	//	return "string"
	//case reflect.Slice, reflect.Array:
	//	return "array"
	//case reflect.Map:
	//	return "object"
	//case reflect.Struct:
	//	return "object"
	//case reflect.Ptr:
	//	return typeToString(t.Elem())
	case reflect.Func:
		return "lambda"
	}
	return "Unsupported (report bug): " + t.Name()
}
