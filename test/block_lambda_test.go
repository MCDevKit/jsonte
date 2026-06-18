package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func TestBlockLambdaBasic(t *testing.T) {
	eval := evaluateScript(t, `
		fn = (x) => { return x * 2; };
		return fn(5);
	`)
	assertReturn(t, eval, types.AsNumber(10))
}

func TestBlockLambdaMultiStatement(t *testing.T) {
	eval := evaluateScript(t, `
		fn = (x) => {
			y = x + 1;
			return y * 2;
		};
		return fn(4);
	`)
	assertReturn(t, eval, types.AsNumber(10))
}

func TestBlockLambdaNoArguments(t *testing.T) {
	eval := evaluateScript(t, `
		fn = () => { return 42; };
		return fn();
	`)
	assertReturn(t, eval, types.AsNumber(42))
}

func TestAddAssignString(t *testing.T) {
	eval := evaluateScript(t, `
		result = 'hello';
		result += ' world';
		return result;
	`)
	assertReturn(t, eval, types.AsString("hello world"))
}

func TestAddAssignNumber(t *testing.T) {
	eval := evaluateScript(t, `
		n = 10;
		n += 5;
		return n;
	`)
	assertReturn(t, eval, types.AsNumber(15))
}

func TestJoinNoSeparator(t *testing.T) {
	eval := evaluate(t, `['a', 'b', 'c'].join()`)
	assertString(t, eval, "abc")
}

func TestJoinWithSeparator(t *testing.T) {
	eval := evaluate(t, `['a', 'b', 'c'].join(', ')`)
	assertString(t, eval, "a, b, c")
}

func TestJoinNumberElements(t *testing.T) {
	eval := evaluate(t, `(1..3).join('-')`)
	assertString(t, eval, "1-2-3")
}

func TestBlockLambdaWithAddAssignAndJoin(t *testing.T) {
	eval := evaluateScript(t, `
		fn = (items) => {
			result = 'start;';
			result += items.map(x => x + ';').join('');
			return result;
		};
		return fn((1..3));
	`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	expected := "start;1;2;3;"
	if signal.Value.StringValue() != expected {
		t.Fatalf("Expected %q, got %q", expected, signal.Value.StringValue())
	}
}

func TestBlockLambdaReturnFromMap(t *testing.T) {
	eval := evaluateScript(t, `
		double = x => { return x * 2; };
		return (1..5).map(double);
	`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	assertArray(t, jsonte.Result{Value: signal.Value, Action: types.Value}, types.AsArray([]interface{}{2, 4, 6, 8, 10}))
}

func TestLambdaLocalVarsIsolatedPerCall(t *testing.T) {
	// Internal variables assigned inside a lambda body must not leak between calls
	// or into variableScope.
	eval := evaluateScript(t, `
		makeRect = (i) => { rx = 'r' + i; ry = 's' + i; return [rx, ry]; };
		usePoint = (rx, ry) => [rx, ry];
		makeRect(0);
		makeRect(1);
		makeRect(2);
		return usePoint('0', '0');
	`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	expected := types.AsArray([]interface{}{"0", "0"})
	if !signal.Value.Equals(expected) {
		t.Fatalf("Expected ['0', '0'], got %s", signal.Value.StringValue())
	}
}

func TestLambdaLocalVarsNotVisibleOutside(t *testing.T) {
	// A variable assigned inside a lambda must not be visible in the outer script scope.
	eval := evaluateScript(t, `
		fn = () => { inner = 42; };
		fn();
		return inner;
	`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	if !types.IsNull(signal.Value) {
		t.Fatalf("Expected null (inner not visible outside lambda), got %s", signal.Value.StringValue())
	}
}

func TestLambdaCanReadScriptLevelVar(t *testing.T) {
	// A lambda should be able to read variables defined at script level (in variableScope).
	eval := evaluateScript(t, `
		multiplier = 3;
		fn = (x) => x * multiplier;
		return fn(7);
	`)
	assertReturn(t, eval, types.AsNumber(21))
}

func TestLambdaThisWritesToScriptScope(t *testing.T) {
	// this.x = value inside a lambda is the explicit escape hatch for writing to script scope.
	eval := evaluateScript(t, `
		fn = () => { this.exported = 99; };
		fn();
		return exported;
	`)
	assertReturn(t, eval, types.AsNumber(99))
}

func TestLineComment(t *testing.T) {
	eval := evaluateScript(t, `
		// this is a comment
		x = 1 + 2; // inline comment
		return x;
	`)
	assertReturn(t, eval, types.AsNumber(3))
}

func TestBlockComment(t *testing.T) {
	eval := evaluateScript(t, `
		/* block
		   comment */
		x = 10;
		return x /* trailing */ + 1;
	`)
	assertReturn(t, eval, types.AsNumber(11))
}
