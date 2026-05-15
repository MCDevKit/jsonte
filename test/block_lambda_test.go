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
