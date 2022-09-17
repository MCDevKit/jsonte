package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
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
		fn.Args = append(fn.Args, of.Type().In(i))
	}
	if (of.Type().NumOut() > 2) || (of.Type().NumOut() == 2 && of.Type().Out(1).String() != "error") {
		panic("Function body must return only one value and can return an error!")
	}
	if (of.Type().NumOut() == 2) || (of.Type().NumOut() == 1 && of.Type().Out(0).String() == "error") {
		fn.WithError = true
	}
	functions[fn.Name] = append(functions[fn.Name], fn)
	if fn.IsInstance {
		if _, ok := instanceFunctions[fn.Args[0].String()]; !ok {
			instanceFunctions[fn.Args[0].String()] = make(map[string][]JsonFunction)
		}
		instanceFunctions[fn.Args[0].String()][fn.Name] = append(instanceFunctions[fn.Args[0].String()][fn.Name], fn)
	}
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

func CallInstanceFunction(name string, instance interface{}, args []interface{}) (interface{}, error) {
	fns, ok := instanceFunctions[reflect.TypeOf(instance).String()][name]
	if !ok {
		return nil, utils.WrappedErrorf("Instance function \"%s\" not found", name)
	}
	a := make([]interface{}, 0)
	a = append(a, instance)
	a = append(a, args...)
	return callFunctionImpl(name, fns, a)
}

func CallFunction(name string, args []interface{}) (interface{}, error) {
	fns, ok := functions[name]
	if !ok {
		return nil, utils.WrappedErrorf("Function \"%s\" not found", name)
	}
	return callFunctionImpl(name, fns, args)
}

func callFunctionImpl(name string, fns []JsonFunction, args []interface{}) (interface{}, error) {
	sizeMatching := make([]JsonFunction, 0)
	for _, fn := range fns {
		if len(fn.Args) == len(args) {
			sizeMatching = append(sizeMatching, fn)
		}
	}
	if len(sizeMatching) == 0 {
		return nil, utils.WrappedErrorf("Incorrect number of arguments for function \"%s\"", name)
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
			expected = append(expected, paramsToString(fn.Args))
		}
		argTypes := make([]reflect.Type, 0)
		for _, arg := range args {
			argTypes = append(argTypes, reflect.TypeOf(arg))
		}
		return nil, utils.WrappedErrorf("Incorrect argument types for function \"%s\". Expected: %s, got: %s", name, strings.Join(expected, ", "), paramsToString(argTypes))
	} else if len(matching) > 1 {
		matched := make([]string, 0)
		for _, fn := range matching {
			matched = append(matched, paramsToString(fn.Args))
		}
		return nil, utils.WrapErrorf(nil, "Ambiguous function call for \"%s\". Matched: %s", name, strings.Join(matched, ", "))
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
		return call[0].Interface(), nil
	}
}

func checkParams(fn JsonFunction, args []interface{}) bool {
	for i, arg := range args {
		if arg == nil || !reflect.TypeOf(arg).AssignableTo(fn.Args[i]) {
			return false
		}
	}
	return true
}

func paramsToString(args []reflect.Type) string {
	join := make([]string, 0)
	for _, arg := range args {
		join = append(join, typeToString(arg))
	}
	return "(" + strings.Join(join, ", ") + ")"
}

func typeToString(t reflect.Type) string {
	if t == nil {
		return "null"
	}
	if t.AssignableTo(reflect.TypeOf(utils.JsonNumber{})) {
		return "number"
	}
	switch t.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.String:
		return "string"
	case reflect.Slice, reflect.Array:
		return "array"
	case reflect.Map:
		return "object"
	case reflect.Struct:
		return "object"
	case reflect.Ptr:
		return typeToString(t.Elem())
	case reflect.Func:
		return "lambda"
	}
	return "unknown"
}
