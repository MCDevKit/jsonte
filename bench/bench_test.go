package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
)

var (
	benchmarkEvalResult    jsonte.Result
	benchmarkProcessResult utils.NavigableMap[string, types.JsonType]
	benchmarkModuleResult  jsonte.JsonModule
)

const (
	benchmarkSimpleExpression      = "1 + 2 * 3"
	benchmarkScopedArrayExpression = "numbers[10]"
	benchmarkSimpleTemplate        = `{
		"$template": {
			"values": "{{=1..40}}",
			"static": 42,
			"{{#1..40}}": {
				"item{{index}}": {
					"original": "{{=value}}",
					"double": "{{=value * 2}}",
					"isEven": "{{=mod(value, 2) == 0}}"
				}
			}
		}
	}`
	benchmarkModuleSource = `{
		"$module": "bench-module",
		"$scope": {
			"base": 10,
			"values": [1,2,3,4,5,6,7,8,9,10]
		},
		"$template": {
			"{{#values}}": {
				"value{{index}}": "{{=value + base}}",
				"base": "{{=base}}"
			}
		}
	}`
	benchmarkModuleTemplate = `{
		"$extend": "bench-module",
		"$template": {
			"{{#values}}": {
				"square{{index}}": "{{=value * value}}",
				"value{{index}}": "{{=value + base + index}}"
			},
			"base": "{{=base + 1}}"
		}
	}`
)

func BenchmarkEvalSimpleExpression(b *testing.B) {
	scope := deque.Deque[*types.JsonObject]{}
	b.ReportAllocs()
	b.ResetTimer()

	var result jsonte.Result
	for i := 0; i < b.N; i++ {
		r, err := jsonte.Eval(benchmarkSimpleExpression, scope, "#")
		if err != nil {
			b.Fatal(err)
		}
		result = r
	}
	benchmarkEvalResult = result
}

func BenchmarkEvalWithScope(b *testing.B) {
	numbers := make([]interface{}, 25)
	for i := range numbers {
		numbers[i] = i
	}
	scopeObject := types.AsObject(map[string]interface{}{
		"numbers": numbers,
	})

	b.ReportAllocs()
	b.ResetTimer()

	var result jsonte.Result
	for i := 0; i < b.N; i++ {
		scope := deque.Deque[*types.JsonObject]{}
		scope.PushBack(scopeObject)
		r, err := jsonte.Eval(benchmarkScopedArrayExpression, scope, "#")
		if err != nil {
			b.Fatal(err)
		}
		result = r
	}
	benchmarkEvalResult = result
}

func BenchmarkProcessSimpleTemplate(b *testing.B) {
	modules := map[string]jsonte.JsonModule{}
	globalScope := types.NewJsonObject()

	b.ReportAllocs()
	b.ResetTimer()

	var result utils.NavigableMap[string, types.JsonType]
	for i := 0; i < b.N; i++ {
		r, err := jsonte.Process("benchmark", benchmarkSimpleTemplate, globalScope, modules, -1)
		if err != nil {
			b.Fatal(err)
		}
		result = r
	}
	benchmarkProcessResult = result
}

func BenchmarkProcessTemplateWithModule(b *testing.B) {
	globalScope := types.NewJsonObject()
	module, err := jsonte.LoadModule(benchmarkModuleSource, globalScope, -1)
	if err != nil {
		b.Fatal(err)
	}
	modules := map[string]jsonte.JsonModule{
		module.Name.StringValue(): module,
	}

	b.ReportAllocs()
	b.ResetTimer()

	var result utils.NavigableMap[string, types.JsonType]
	for i := 0; i < b.N; i++ {
		r, err := jsonte.Process("benchmark", benchmarkModuleTemplate, globalScope, modules, -1)
		if err != nil {
			b.Fatal(err)
		}
		result = r
	}
	benchmarkProcessResult = result
}

func BenchmarkLoadModule(b *testing.B) {
	globalScope := types.NewJsonObject()

	b.ReportAllocs()
	b.ResetTimer()

	var module jsonte.JsonModule
	for i := 0; i < b.N; i++ {
		m, err := jsonte.LoadModule(benchmarkModuleSource, globalScope, -1)
		if err != nil {
			b.Fatal(err)
		}
		module = m
	}
	benchmarkModuleResult = module
}
