//go:build js && wasm

package main

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/fatih/color"
	"github.com/gammazero/deque"
	"go.uber.org/zap/zapcore"
	"syscall/js"
)

var (
	commit      string
	version     = "0.0.0"
	date        string
	buildSource = "DEV"
)

func main() {
	color.NoColor = true
	utils.InitLogging(zapcore.InfoLevel)
	functions.Init()
	fmt.Println("jsonte WebAssembly")
	js.Global().Set("jsonte", js.ValueOf(map[string]interface{}{
		"version": version,
		"eval":    evalWrapper(),
	}))
	<-make(chan bool)
}

func evalWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		expression := args[0].String()
		scope := utils.NewNavigableMap[string, interface{}]()
		if len(args) > 1 && !args[1].IsNull() {
			if args[1].Type() != js.TypeObject {
				return toError(utils.WrappedErrorf("Scope is not an object!"))
			}
			scope = toGo(args[1]).(utils.NavigableMap[string, interface{}])
		}
		s := deque.Deque[interface{}]{}
		s.PushBack(scope)
		result, err := jsonte.Eval(expression, s, "#")
		if err != nil {
			return toError(err)
		}
		return js.ValueOf(result)
	})
}

func toError(err error) js.Value {
	return js.ValueOf(map[string]interface{}{
		"error": err.Error(),
	})
}

func toGo(value js.Value) interface{} {
	switch value.Type() {
	case js.TypeBoolean:
		return value.Bool()
	case js.TypeNumber:
		return value.Float()
	case js.TypeString:
		return value.String()
	case js.TypeObject:
		if value.Get("length").Type() == js.TypeNumber {
			arr := make([]interface{}, value.Get("length").Int())
			for i := 0; i < len(arr); i++ {
				arr[i] = toGo(value.Index(i))
			}
			return arr
		}
		obj := utils.NewNavigableMap[string, interface{}]()
		js.Global().Get("Object").Call("entries", value).Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			obj.Put(args[0].String(), toGo(args[1]))
			return nil
		}))
		return obj
	default:
		return nil
	}
}
