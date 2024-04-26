package test

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"testing"
)

func evaluateScript(t *testing.T, text string) jsonte.Result {
	scope := deque.Deque[*types.JsonObject]{}
	eval, err := jsonte.EvalScript(text, scope, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func assertScriptError(t *testing.T, text string, error []string) {
	scope := deque.Deque[*types.JsonObject]{}
	_, err := jsonte.EvalScript(text, scope, "#")
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if burrito.IsBurritoError(err) {
		split := burrito.GetAllMessages(err)
		if len(split) != len(error) {
			for i := 0; i < len(split); i++ {
				t.Logf("Line %d: %s", i, split[i])
			}
			t.Fatalf("Error is not correct (expected %d lines, got %d)", len(error), len(split))
		}
		for i := 0; i < len(split); i++ {
			if split[i] != error[i] {
				for i := 0; i < len(split); i++ {
					t.Logf("Line %d: %s", i, split[i])
				}
				t.Fatalf("Error is not correct (expected %s, got %s)", error[i], split[i])
			}
		}
	} else {
		t.Fatalf("Error is not a burrito error (%s)", err.Error())
	}
}

func assertReturn(t *testing.T, eval jsonte.Result, value types.JsonType) {
	if (eval.Value == nil || types.IsNull(eval.Value)) && value != nil && value != types.Null {
		t.Fatalf("Return value is null")
	}
	if !eval.Value.Equals(value) {
		t.Fatalf("Return value is incorrect")
	}
}

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

func TestForLoop(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "for i in 0..10 { testCall('loop', i); }")
	AssertTestFunctionCalledNTimes(t, "loop", 11)
	for i := 0; i <= 10; i++ {
		AssertTestFunctionCalledWith(t, "loop", i, types.AsNumber(i))
	}
}

func TestForLoopBreak(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "for i in 0..10 { testCall('loop', i); if i == 5 { break; } }")
	AssertTestFunctionCalledNTimes(t, "loop", 6)
	for i := 0; i <= 5; i++ {
		AssertTestFunctionCalledWith(t, "loop", i, types.AsNumber(i))
	}
}

func TestForLoopContinue(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "for i in 0..10 { if i == 5 { continue; } testCall('loop', i); }")
	AssertTestFunctionCalledNTimes(t, "loop", 10)
	index := 0
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		AssertTestFunctionCalledWith(t, "loop", index, types.AsNumber(i))
		index++
	}
}

func TestWhileLoop(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "i = 0; while i < 10 { testCall('loop', i); i = i + 1; }")
	AssertTestFunctionCalledNTimes(t, "loop", 10)
	for i := 0; i <= 9; i++ {
		AssertTestFunctionCalledWith(t, "loop", i, types.AsNumber(i))
	}
}

func TestDoWhileLoop(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "i = 0; do { testCall('loop', i); i = i + 1; } while i < 10;")
	AssertTestFunctionCalledNTimes(t, "loop", 10)
	for i := 0; i <= 9; i++ {
		AssertTestFunctionCalledWith(t, "loop", i, types.AsNumber(i))
	}
}

func TestEarlyReturn(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "testCall('before'); return; testCall('after');")
	AssertTestFunctionCalled(t, "before")
	AssertTestFunctionNotCalled(t, "after")
}

func TestValueReturn(t *testing.T) {
	eval := evaluateScript(t, "return 'hello world';")
	assertReturn(t, eval, &types.JsonString{Value: "hello world"})
}

func TestReturnInLoop(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "for i in 0..10 { testCall('loop', i); return; }")
	AssertTestFunctionCalledNTimes(t, "loop", 1)
	AssertTestFunctionCalledWith(t, "loop", 0, types.AsNumber(0))
}

func TestLooseBreak(t *testing.T) {
	ClearTestCalls()
	assertScriptError(t, "break; testCall('after');", []string{"break outside of loop"})
	AssertTestFunctionNotCalled(t, "after")
}

func TestLooseContinue(t *testing.T) {
	ClearTestCalls()
	assertScriptError(t, "continue; testCall('after');", []string{"continue outside of loop"})
	AssertTestFunctionNotCalled(t, "after")
}

func TestIfStatement(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "if 1 == 1 { testCall('true'); }")
	AssertTestFunctionCalled(t, "true")
}

func TestIfElseStatement(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "if 1 == 0 { testCall('true'); } else { testCall('false'); }")
	AssertTestFunctionNotCalled(t, "true")
	AssertTestFunctionCalled(t, "false")
}

func TestIfElseIfStatement(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "if 1 == 0 { testCall('true'); } else if 1 == 1 { testCall('false'); }")
	AssertTestFunctionNotCalled(t, "true")
	AssertTestFunctionCalled(t, "false")
}

func TestIfElseIfElseStatement(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "if 1 == 0 { testCall('true'); } else if 1 == 0 { testCall('false'); } else { testCall('else'); }")
	AssertTestFunctionNotCalled(t, "true")
	AssertTestFunctionNotCalled(t, "false")
	AssertTestFunctionCalled(t, "else")
}

func TestIfMultipleElseIfStatement(t *testing.T) {
	ClearTestCalls()
	evaluateScript(t, "if 1 == 0 { testCall('true'); } else if 1 == 0 { testCall('false'); } else if 1 == 0 { testCall('else'); }")
	AssertTestFunctionNotCalled(t, "true")
	AssertTestFunctionNotCalled(t, "false")
	AssertTestFunctionNotCalled(t, "else")
}
