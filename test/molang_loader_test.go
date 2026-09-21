package test

import (
	"strings"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
)

// setupMolangLoader serves files from a fake file system and enables loadMolang with the given scope.
func setupMolangLoader(t *testing.T, files map[string]interface{}, scope *types.JsonObject, scopeFiles []string, validate bool) {
	t.Helper()
	safeio.Resolver = safeio.CreateFakeFS(files, false)
	jsonte.EnableMolangLoading(scope, scopeFiles, validate)
	t.Cleanup(func() {
		safeio.Resolver = CacheFS
		functions.SetMolangLoader(nil, nil)
	})
}

func TestLoadMolangWithParams(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/rect.molang": "# rect index\nvariable.i = #{index};",
	}, types.NewJsonObject(), nil, true)
	eval := evaluate(t, "loadMolang('loader/rect.molang', {'index': 3})")
	assertString(t, eval, "v.i=3;")
	// The cache is per parameters
	eval = evaluate(t, "loadMolang('loader/rect.molang', {'index': 4})")
	assertString(t, eval, "v.i=4;")
}

func TestLoadMolangParamsOverrideScope(t *testing.T) {
	scope := types.AsObject(utils.ToNavigableMap("speed", 1, "scale", 2))
	setupMolangLoader(t, map[string]interface{}{
		"loader/speed.molang": "variable.speed = #{speed * scale};",
	}, scope, nil, true)
	eval := evaluateWithObjectScope(t, "loadMolang('loader/speed.molang')", scope)
	assertString(t, eval, "v.speed=2;")
	eval = evaluateWithObjectScope(t, "loadMolang('loader/speed.molang', {'speed': 5})", scope)
	assertString(t, eval, "v.speed=10;")
}

func TestLoadTextProcessesScopeMolangFiles(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/scope.molang": "variable.x = #{1 + 1};",
		"loader/other.molang": "variable.x = #{1 + 1};",
	}, types.NewJsonObject(), []string{"loader\\scope.molang"}, true)
	eval := evaluate(t, "loadText('loader/scope.molang')")
	assertString(t, eval, "v.x=2;")
	// Files outside scope directories are returned unchanged
	eval = evaluate(t, "loadText('loader/other.molang')")
	assertString(t, eval, "variable.x = #{1 + 1};")
}

func TestLoadMolangValidation(t *testing.T) {
	files := map[string]interface{}{
		"loader/invalid.molang": "variable.i = __INDEX__;",
	}
	setupMolangLoader(t, files, types.NewJsonObject(), nil, true)
	assertError(t, "loadMolang('loader/invalid.molang')", []string{
		"Error calling function 'loadMolang'",
		"Failed to process the molang file loader/invalid.molang",
		"unknown identifier '__INDEX__' — expected math./query./variable./temp./context. (line 1, column 14 of the processed file: variable.i = __INDEX__;)",
	})

	setupMolangLoader(t, files, types.NewJsonObject(), nil, false)
	eval := evaluate(t, "loadMolang('loader/invalid.molang')")
	assertString(t, eval, "v.i=__INDEX__;")
}

func TestLoadMolangNested(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/outer.molang": "variable.a = 1; #{loadMolang('loader/inner.molang', {'value': 2})}",
		"loader/inner.molang": "variable.b = #{value};",
	}, types.NewJsonObject(), nil, true)
	eval := evaluate(t, "loadMolang('loader/outer.molang')")
	assertString(t, eval, "v.a=1;v.b=2;")
}

func TestLoadMolangInJsonteScript(t *testing.T) {
	scope := types.AsObject(utils.ToNavigableMap("prefix", "demo"))
	setupMolangLoader(t, map[string]interface{}{
		"loader/script.molang": "variable.id = '#{prefix}_#{index}';",
	}, scope, nil, true)
	if err := jsonte.EvalJsonteFile(`$scope.effect = (i) => loadMolang('loader/script.molang', {'index': i}); $scope.first = effect(1);`, scope, "#"); err != nil {
		t.Fatal(err)
	}
	if scope.Get("first").StringValue() != "v.id='demo_1';" {
		t.Fatalf("Unexpected result: %s", scope.Get("first").StringValue())
	}
}

func TestLoadMolangWithoutLoader(t *testing.T) {
	functions.SetMolangLoader(nil, nil)
	_, err := jsonte.QuickEval("loadMolang('loader/any.molang')", "#")
	if err == nil || !strings.Contains(err.Error(), "Molang files can only be loaded when compiling with a scope") {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// Expressions that do not mention a parameter are evaluated once; the others for every set of parameters
func TestLoadMolangCachesParameterIndependentExpressions(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/cached.molang": "variable.a = #{[testCall('once'), 1][1]}; variable.b = #{[testCall('each'), index][1]};",
	}, types.NewJsonObject(), nil, true)
	ClearTestCalls()
	eval := evaluate(t, "loadMolang('loader/cached.molang', {'index': 2})")
	assertString(t, eval, "v.a=1;v.b=2;")
	eval = evaluate(t, "loadMolang('loader/cached.molang', {'index': 3})")
	assertString(t, eval, "v.a=1;v.b=3;")
	AssertTestFunctionCalledNTimes(t, "once", 1)
	AssertTestFunctionCalledNTimes(t, "each", 2)
}

func TestLoadMolangCycle(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/cycle_a.molang": "variable.a = 1; #{loadMolang('loader/cycle_b.molang')}",
		"loader/cycle_b.molang": "variable.b = 1; #{loadMolang('loader/cycle_a.molang')}",
	}, types.NewJsonObject(), nil, true)
	_, err := jsonte.QuickEval("loadMolang('loader/cycle_a.molang')", "#")
	if err == nil || !strings.Contains(err.Error(), "Circular loadMolang: loader/cycle_a.molang -> loader/cycle_b.molang -> loader/cycle_a.molang") {
		t.Fatalf("Unexpected error: %v", err)
	}
	// A failed load must not leave the chain behind
	_, err = jsonte.QuickEval("loadMolang('loader/cycle_b.molang')", "#")
	if err == nil || !strings.Contains(err.Error(), "Circular loadMolang: loader/cycle_b.molang -> loader/cycle_a.molang -> loader/cycle_b.molang") {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestLoadMolangEndlessRecursionWithParams(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/endless.molang": "#{loadMolang('loader/endless.molang', {'depth': depth + 1})}",
	}, types.NewJsonObject(), nil, true)
	_, err := jsonte.QuickEval("loadMolang('loader/endless.molang', {'depth': 0})", "#")
	if err == nil || !strings.Contains(err.Error(), "loadMolang is nested more than 100 levels deep") {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestLoadMolangFiniteRecursion(t *testing.T) {
	setupMolangLoader(t, map[string]interface{}{
		"loader/countdown.molang": "#{depth > 0 ? loadMolang('loader/countdown.molang', {'depth': depth - 1}) : 'variable.x = 1;'}",
	}, types.NewJsonObject(), nil, true)
	eval := evaluate(t, "loadMolang('loader/countdown.molang', {'depth': 5})")
	assertString(t, eval, "v.x=1;")
}
