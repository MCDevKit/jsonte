package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func assertVariable(t *testing.T, eval jsonte.Result, path *types.JsonPath, value types.JsonType) {
	if eval.VariableScope == nil {
		t.Fatalf("Variable scope is nil")
	}
	index, err := path.Get(eval.VariableScope)
	if err != nil {
		t.Fatalf("Error occured while accessing %s: %s", path.StringValue(), err.Error())
	}
	if index == nil || types.IsNull(index) {
		t.Fatalf("Variable scope does not contain %s", path.StringValue())
	}
	if !index.Equals(value) {
		t.Fatalf("Variable scope value %s is incorrect", path.StringValue())
	}
}

func assertScope(t *testing.T, eval jsonte.Result, path *types.JsonPath, value types.JsonType) {
	scopeObj := &types.JsonObject{
		StackValue:  &eval.Scope,
		StackTarget: eval.VariableScope,
	}
	index, err := path.Get(scopeObj)
	if err != nil {
		t.Fatalf("Error occured while accessing %s: %s", path.StringValue(), err.Error())
	}
	if index == nil || types.IsNull(index) {
		t.Fatalf("Variable scope does not contain %s", path.StringValue())
	}
	if !index.Equals(value) {
		t.Fatalf("Variable scope value %s is incorrect", path.StringValue())
	}
}

func TestAssignment(t *testing.T) {
	eval := evaluate(t, "hello = 'world'")
	path, err := types.ParseJsonPath("#/hello")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertVariable(t, eval, path, &types.JsonString{Value: "world"})
}

func TestReassignment(t *testing.T) {
	eval := evaluateWithScope(t, "hello = 'world'", utils.ToNavigableMap("hello", types.AsNumber(1)))
	path, err := types.ParseJsonPath("#/hello")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertScope(t, eval, path, &types.JsonString{Value: "world"})
}

func TestAssignmentInObject(t *testing.T) {
	object, err := types.ParseJsonObject([]byte("{\"foo\": {\"bar\": \"baz\"}}"))
	if err != nil {
		t.Fatalf("Failed to parse object: %s", err.Error())
	}
	eval := evaluateWithObjectScope(t, "foo.bar = 'hello world'", object)
	path, err := types.ParseJsonPath("#/foo/bar")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertScope(t, eval, path, &types.JsonString{Value: "hello world"})
}

func TestAssignmentInArray(t *testing.T) {
	object, err := types.ParseJsonObject([]byte("{\"foo\": [0]}"))
	if err != nil {
		t.Fatalf("Failed to parse object: %s", err.Error())
	}
	eval := evaluateWithObjectScope(t, "foo[0] = 'hello world'", object)
	path, err := types.ParseJsonPath("#/foo[0]")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertScope(t, eval, path, &types.JsonString{Value: "hello world"})
}

func TestAssignmentInArrayWithVariable(t *testing.T) {
	object, err := types.ParseJsonObject([]byte("{\"foo\": [0], \"bar\": 1}"))
	if err != nil {
		t.Fatalf("Failed to parse object: %s", err.Error())
	}
	eval := evaluateWithObjectScope(t, "foo[0] = bar", object)
	path, err := types.ParseJsonPath("#/foo[0]")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertScope(t, eval, path, types.AsNumber(1))
}

func TestAssignmentInObjectByPath(t *testing.T) {
	object, err := types.ParseJsonObject([]byte("{\"foo\": {\"bar\": {\"baz\": 1}}}"))
	if err != nil {
		t.Fatalf("Failed to parse object: %s", err.Error())
	}
	eval := evaluateWithObjectScope(t, "foo[jsonPath('#/bar/baz')] = 'hello world'", object)
	path, err := types.ParseJsonPath("#/foo/bar/baz")
	if err != nil {
		t.Errorf("Failed to parse path: %s", err.Error())
	}
	assertScope(t, eval, path, &types.JsonString{Value: "hello world"})
}
