package test

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

// assertExpressionError checks that an expression returns an error with the given messages,
// without enforcing the action type (useful for action-prefix expressions like #, ?, =).
func assertExpressionError(t *testing.T, text string, expected []string) {
	t.Helper()
	_, err := jsonte.QuickEval(text, "#")
	if err == nil {
		t.Fatalf("Expected error for %q, got none", text)
	}
	if burrito.IsBurritoError(err) {
		split := burrito.GetAllMessages(err)
		if len(split) != len(expected) {
			for i, msg := range split {
				t.Logf("Line %d: %s", i, msg)
			}
			t.Fatalf("Error line count mismatch: expected %d, got %d", len(expected), len(split))
		}
		for i := range split {
			if split[i] != expected[i] {
				t.Fatalf("Error line %d mismatch: expected %q, got %q", i, expected[i], split[i])
			}
		}
	} else {
		t.Fatalf("Error is not a burrito error: %s", err.Error())
	}
}

// --- valid uses: lambdas stored and called, not returned as top-level expression ---

func TestLambdaAsObjectFieldCalled(t *testing.T) {
	// Lambda stored as object value and then immediately called
	eval := evaluate(t, `{"fn": x => x * 2}.fn(5)`)
	assertNumber(t, eval, 10)
}

func TestLambdaMultiArgAsObjectField(t *testing.T) {
	eval := evaluate(t, `{"fn": (a, b) => a + b}.fn(3, 4)`)
	assertNumber(t, eval, 7)
}

func TestLambdaInArrayCalled(t *testing.T) {
	// Lambda stored in array element and then called
	eval := evaluate(t, `[x => x + 1][0](3)`)
	assertNumber(t, eval, 4)
}

func TestLambdaNoArgInObject(t *testing.T) {
	eval := evaluate(t, `{"fn": () => 42}.fn()`)
	assertNumber(t, eval, 42)
}

func TestLambdaInTernaryCalled(t *testing.T) {
	// Ternary choosing between two lambdas, then calling the result
	eval := evaluate(t, `(true ? (x => x * 2) : (x => x * 3))(5)`)
	assertNumber(t, eval, 10)
}

func TestLambdaInTernaryFalseCalled(t *testing.T) {
	eval := evaluate(t, `(false ? (x => x * 2) : (x => x * 3))(5)`)
	assertNumber(t, eval, 15)
}

func TestLambdaViaNullCoalescingCalled(t *testing.T) {
	// Null coalescing produces a lambda, then it's called
	eval := evaluate(t, `(null ?? (x => x * 2))(5)`)
	assertNumber(t, eval, 10)
}

func TestLambdaInScriptAssignAndCall(t *testing.T) {
	// In a script, assign lambda to variable then call it
	eval := evaluateScript(t, `fn = x => x + 1; return fn(5);`)
	assertReturn(t, eval, types.AsNumber(6))
}

func TestLambdaInScriptReturned(t *testing.T) {
	// Script can return a lambda as its result; unwrap the return signal to get the lambda
	eval := evaluateScript(t, `fn = x => x * 3; return fn;`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	l, ok := signal.Value.(*types.JsonLambda)
	if !ok {
		t.Fatalf("Expected lambda inside signal, got %T", signal.Value)
	}
	result, err := l.Call(types.AsNumber(4))
	if err != nil {
		t.Fatal(err)
	}
	if types.AsNumber(result).IntValue() != 12 {
		t.Fatalf("Expected 12, got %s", result.StringValue())
	}
}

func TestLambdaStoredAndPassedToMap(t *testing.T) {
	// Lambda stored in variable, passed to map function
	eval := evaluateScript(t, `fn = x => x * 2; return (1..5).map(fn);`)
	signal, ok := eval.Value.(*types.JsonSignal)
	if !ok {
		t.Fatalf("Expected return signal, got %T", eval.Value)
	}
	assertArray(t, jsonte.Result{Value: signal.Value, Action: types.Value}, types.AsArray([]interface{}{2, 4, 6, 8, 10}))
}

// --- invalid uses ---

func TestLambdaAdditionError(t *testing.T) {
	assertError(t, `(x => x) + 1`, []string{
		"Cannot use a lambda in an addition or concatenation",
	})
}

func TestLambdaAdditionError2(t *testing.T) {
	assertError(t, `1 + (x => x)`, []string{
		"Cannot use a lambda in an addition or concatenation",
	})
}

func TestLambdaConcatenationError(t *testing.T) {
	assertError(t, `"hello" + (x => x)`, []string{
		"Cannot use a lambda in an addition or concatenation",
	})
}

func TestLambdaSubtractError(t *testing.T) {
	assertError(t, `(x => x) - 1`, []string{
		"Cannot use a lambda in an arithmetic operation",
	})
}

func TestLambdaMultiplyError(t *testing.T) {
	assertError(t, `(x => x) * 2`, []string{
		"Cannot use a lambda in an arithmetic operation",
	})
}

func TestLambdaDivideError(t *testing.T) {
	assertError(t, `(x => x) / 2`, []string{
		"Cannot use a lambda in an arithmetic operation",
	})
}

func TestLambdaNegationError(t *testing.T) {
	assertError(t, `-(x => x)`, []string{
		"Cannot negate a lambda",
	})
}

func TestLambdaRangeError(t *testing.T) {
	assertError(t, `(x => x)..(y => y)`, []string{
		"Cannot use a lambda in a range expression",
	})
}

func TestLambdaIndexAsObjectError(t *testing.T) {
	assertError(t, `(x => x)[0]`, []string{
		"Cannot index a lambda",
	})
}

func TestLambdaIndexAsIndexError(t *testing.T) {
	assertError(t, `[1, 2, 3][x => x]`, []string{
		"Cannot use a lambda as an index",
	})
}

func TestLambdaLessError(t *testing.T) {
	assertError(t, `(x => x) < 1`, []string{
		"Cannot compare a lambda",
	})
}

func TestLambdaGreaterError(t *testing.T) {
	assertError(t, `(x => x) > 1`, []string{
		"Cannot compare a lambda",
	})
}

func TestLambdaLessEqualError(t *testing.T) {
	assertError(t, `(x => x) <= 1`, []string{
		"Cannot compare a lambda",
	})
}

func TestLambdaGreaterEqualError(t *testing.T) {
	assertError(t, `(x => x) >= 1`, []string{
		"Cannot compare a lambda",
	})
}

func TestLambdaIterationActionError(t *testing.T) {
	assertExpressionError(t, `#(x => x)`, []string{
		"Cannot iterate over a lambda",
	})
}

func TestLambdaPredicateActionError(t *testing.T) {
	assertExpressionError(t, `?(x => x)`, []string{
		"Cannot use a lambda as a predicate",
	})
}

func TestLambdaLiteralActionError(t *testing.T) {
	assertExpressionError(t, `=(x => x)`, []string{
		"Cannot use a lambda as a literal value",
	})
}

func TestLambdaValueActionError(t *testing.T) {
	assertError(t, `x => x`, []string{
		"Cannot use a lambda as a template value",
	})
}
