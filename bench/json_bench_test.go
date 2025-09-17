package bench

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	processTemplateSimple = `{
		"$template": {
			"metadata": {
				"generated": "{{='benchmark'}}",
				"count": "{{=(1..64).Count()}}"
			},
			"entries": {
				"{{#1..64}}": {
					"id": "entry_{{index}}",
					"value": "{{=value}}",
					"double": "{{=value * 2}}",
				"even": "{{=mod(value, 2) == 0}}"
				}
			}
		}
	}`

	processModuleSource = `{
		"$module": "bench-module",
		"$scope": {
			"base": 100,
			"numbers": "{{=1..64}}",
			"suffix": "mod"
		},
		"$template": {
			"{{#numbers}}": {
				"value{{index}}": "{{=value + base}}",
				"tag": "{{='tag-' + suffix}}"
			}
		}
	}`

	processModuleTemplate = `{
		"$extend": "bench-module",
		"$template": {
			"summary": {
				"total": "{{=numbers.Count()}}",
				"max": "{{=max(numbers)}}"
			}
		}
	}`

	processAssertionsInput = `[
		"1 + 1 == 2",
		"'bench' == 'bench'",
		"(1..10).Count() == 10"
	]`

	processLangInput = "bench.value=##{'bench-'}##{'case'}\nbench.flag=##{true}"

	processStringInput = "Name: #{name}, Score: #{score}, Double: #{score * 2}"
)

var (
	processModuleOnce sync.Once
	processModuleMap  map[string]jsonte.JsonModule
	processScope      *types.JsonObject
	processModulesRes utils.NavigableMap[string, types.JsonType]
	processStringRes  string
	findAnyCaseObject *types.JsonObject
)

func BenchmarkJSONProcessSimpleTemplate(b *testing.B) {
	ensureProcessBenchData()
	runProcessBenchmark(b, "bench_simple", processTemplateSimple, types.NewJsonObject(), nil)
}

func BenchmarkJSONProcessTemplateWithModule(b *testing.B) {
	ensureProcessBenchData()
	runProcessBenchmark(b, "bench_module", processModuleTemplate, types.NewJsonObject(), processModuleMap)
}

func BenchmarkJSONProcessAssertionsFile(b *testing.B) {
	ensureProcessBenchData()
	globalScope := types.NewJsonObject()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := jsonte.ProcessAssertionsFile("bench", processAssertionsInput, globalScope, -1); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONProcessLangFile(b *testing.B) {
	ensureProcessBenchData()
	scope := types.NewJsonObject()
	scope.Put("case", types.AsString("lang"))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := jsonte.ProcessLangFile(processLangInput, scope)
		if err != nil {
			b.Fatal(err)
		}
		processStringRes = res
	}
}

func BenchmarkJSONProcessString(b *testing.B) {
	ensureProcessBenchData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := jsonte.ProcessString(processStringInput, processScope, "#", "")
		if err != nil {
			b.Fatal(err)
		}
		processStringRes = res
	}
}

func BenchmarkJSONFindAnyCaseLegacy(b *testing.B) {
	ensureProcessBenchData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := legacyFindAnyCase[*types.JsonString](findAnyCaseObject, "file", "name"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSONFindAnyCaseOptimized(b *testing.B) {
	ensureProcessBenchData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := jsonte.FindAnyCase[*types.JsonString](findAnyCaseObject, "file", "name"); err != nil {
			b.Fatal(err)
		}
	}
}

func runProcessBenchmark(b *testing.B, name, template string, scope *types.JsonObject, modules map[string]jsonte.JsonModule) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := jsonte.Process(name, template, scope, modules, -1)
		if err != nil {
			b.Fatal(err)
		}
		processModulesRes = res
	}
}

func ensureProcessBenchData() {
	processModuleOnce.Do(func() {
		scope := types.NewJsonObject()
		scope.Put("name", types.AsString("jsonte"))
		scope.Put("score", types.AsNumber(42))
		processScope = scope

		module, err := jsonte.LoadModule(processModuleSource, types.NewJsonObject(), -1)
		if err != nil {
			panic(err)
		}
		processModuleMap = map[string]jsonte.JsonModule{
			module.Name.StringValue(): module,
		}

		obj, err := types.ParseJsonObject([]byte(`{
			"FileName": "example.json",
			"metadata": {
				"file_name": "alt.json",
				"camel": "FileName"
			}
		}`))
		if err != nil {
			panic(err)
		}
		findAnyCaseObject = obj
	})
}

func legacyFindAnyCase[T types.JsonType](o *types.JsonObject, key ...string) (*T, error) {
	if key == nil || len(key) == 0 {
		return nil, errors.New("missing key")
	}
	if len(key) > 1 {
		snakeCase := strings.Join(key, "_")
		for _, k := range o.Keys() {
			if strings.EqualFold(k, snakeCase) {
				return legacyTypedResult[T](k, o.Get(k))
			}
		}
	}
	camelCase := legacyCamelCase(key...)
	for _, k := range o.Keys() {
		if strings.EqualFold(k, camelCase) {
			return legacyTypedResult[T](k, o.Get(k))
		}
	}
	return nil, errors.New("not found")
}

func legacyTypedResult[T types.JsonType](key string, value types.JsonType) (*T, error) {
	if typed, ok := value.(T); ok {
		return &typed, nil
	}
	return nil, errors.New("type mismatch")
}

var legacyCaser = cases.Title(language.Und)

func legacyCamelCase(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}
	var sb strings.Builder
	for _, v := range parts {
		sb.WriteString(legacyCaser.String(v))
	}
	return sb.String()
}
