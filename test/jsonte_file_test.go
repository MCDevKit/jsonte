package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func evalJsonteFile(t *testing.T, script string, scope *types.JsonObject) *types.JsonObject {
	t.Helper()
	if err := jsonte.EvalJsonteFile(script, scope, "#"); err != nil {
		t.Fatal(err)
	}
	return scope
}

// Mutating an existing scope key persists without $scope.
func TestJsonteMutateExistingKey(t *testing.T) {
	scope := types.AsObject(utils.ToNavigableMap("x", types.AsNumber(10)))
	evalJsonteFile(t, `x = x * 2;`, scope)
	if types.AsNumber(scope.Get("x")).IntValue() != 20 {
		t.Fatalf("expected x=20, got %s", scope.Get("x").StringValue())
	}
}

// Mutating a nested existing object persists.
func TestJsonteMutateNestedKey(t *testing.T) {
	inner := types.NewJsonObject()
	inner.Put("a", types.AsNumber(1))
	scope := types.AsObject(utils.ToNavigableMap("obj", inner))
	evalJsonteFile(t, `obj.b = 99;`, scope)
	obj := scope.Get("obj").(*types.JsonObject)
	if types.AsNumber(obj.Get("b")).IntValue() != 99 {
		t.Fatalf("expected obj.b=99, got %s", obj.Get("b").StringValue())
	}
}

// Adding a new top-level key via $scope persists.
func TestJsonteScopeNewTopLevelKey(t *testing.T) {
	scope := types.NewJsonObject()
	evalJsonteFile(t, `$scope.newKey = 42;`, scope)
	if types.AsNumber(scope.Get("newKey")).IntValue() != 42 {
		t.Fatalf("expected newKey=42, got %s", scope.Get("newKey").StringValue())
	}
}

// Adding a new nested object via $scope.
func TestJsonteScopeNewObject(t *testing.T) {
	scope := types.NewJsonObject()
	evalJsonteFile(t, `$scope.cfg = {"enabled": true, "count": 5};`, scope)
	cfg, ok := scope.Get("cfg").(*types.JsonObject)
	if !ok {
		t.Fatalf("expected cfg to be an object, got %T", scope.Get("cfg"))
	}
	if !cfg.Get("enabled").BoolValue() {
		t.Fatalf("expected cfg.enabled=true")
	}
	if types.AsNumber(cfg.Get("count")).IntValue() != 5 {
		t.Fatalf("expected cfg.count=5")
	}
}

// Lambda stored via $scope is callable.
func TestJsonteScopeLambda(t *testing.T) {
	scope := types.NewJsonObject()
	evalJsonteFile(t, `$scope.double = x => x * 2;`, scope)
	fn, ok := scope.Get("double").(*types.JsonLambda)
	if !ok {
		t.Fatalf("expected double to be a lambda, got %T", scope.Get("double"))
	}
	result, err := fn.Call(types.AsNumber(7))
	if err != nil {
		t.Fatal(err)
	}
	if types.AsNumber(result).IntValue() != 14 {
		t.Fatalf("expected 14, got %s", result.StringValue())
	}
}

// Script-local variables (e.g. loop counters) are NOT persisted to scope.
func TestJsonteLocalVarsNotPersisted(t *testing.T) {
	scope := types.NewJsonObject()
	evalJsonteFile(t, `i = 0; for i in 1..5 { }`, scope)
	if scope.ContainsKey("i") {
		t.Fatalf("expected i to not be in scope, but it is: %s", scope.Get("i").StringValue())
	}
	if scope.ContainsKey("$scope") {
		t.Fatalf("$scope pseudo-variable should not be persisted to scope")
	}
}

// $scope is not available outside .jsonte context.
func TestScopePseudoVarOnlyInJsonteContext(t *testing.T) {
	// In normal expression evaluation, $scope resolves to null (no global scope set).
	eval := evaluate(t, `$scope`)
	assertNull(t, eval)
}

// Multiple operations in one script, mixing mutation and $scope.
func TestJsonteCombined(t *testing.T) {
	inner := types.NewJsonObject()
	inner.Put("base", types.AsNumber(10))
	scope := types.AsObject(utils.ToNavigableMap("data", inner))
	evalJsonteFile(t, `
		data.derived = data.base * 3;
		$scope.total = data.base + data.derived;
	`, scope)
	data := scope.Get("data").(*types.JsonObject)
	if types.AsNumber(data.Get("derived")).IntValue() != 30 {
		t.Fatalf("expected data.derived=30")
	}
	if types.AsNumber(scope.Get("total")).IntValue() != 40 {
		t.Fatalf("expected total=40, got %s", scope.Get("total").StringValue())
	}
}
