package test

import (
	"github.com/gammazero/deque"
	"jsonte/jsonte"
	"jsonte/jsonte/utils"
	"reflect"
	"testing"
)

func assertAction(t *testing.T, eval jsonte.Result, action utils.JsonAction) {
	t.Helper()
	if eval.Action != action {
		t.Errorf("Action is not %d", action)
	}
}

func assertArray(t *testing.T, eval jsonte.Result, expected utils.JsonArray) utils.JsonArray {
	t.Helper()
	assertAction(t, eval, utils.Value)
	if !utils.IsArray(eval.Value) {
		t.Fatalf("Result is not an array (%s)", reflect.TypeOf(eval.Value).Name())
	}
	array, ok := eval.Value.(utils.JsonArray)
	if !ok {
		t.Fatalf("Result is not a JSON array (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if len(array) != len(expected) {
		t.Fatalf("Array length is not correct (expected %d, got %d)", len(expected), len(array))
	}
	for i := 0; i < len(expected); i++ {
		if array[i] != expected[i] {
			t.Errorf("Array element %d is not correct (expected %s, got %s)", i, utils.ToString(array[i]), utils.ToString(expected[i]))
		}
	}
	return array
}

func assertNumber(t *testing.T, eval jsonte.Result, expected float64) {
	t.Helper()
	assertAction(t, eval, utils.Value)
	if !utils.IsNumber(eval.Value) {
		t.Fatalf("Result is not a number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	number, ok := eval.Value.(utils.JsonNumber)
	if !ok {
		t.Fatalf("Result is not a JSON number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if number.FloatValue() != expected {
		t.Fatalf("Result is not correct (expected %f, got %f)", expected, number.FloatValue())
	}
}

func assertBool(t *testing.T, eval jsonte.Result, expected bool) {
	t.Helper()
	assertAction(t, eval, utils.Value)
	if _, ok := eval.Value.(bool); !ok {
		t.Fatalf("Result is not a boolean (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if eval.Value.(bool) != expected {
		t.Fatalf("Result is not correct (expected %t, got %t)", expected, eval.Value)
	}
}

func evaluate(t *testing.T, text string) jsonte.Result {
	t.Helper()
	eval, err := jsonte.QuickEval(text, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func evaluateWithScope(t *testing.T, text string, scope utils.JsonObject) jsonte.Result {
	t.Helper()
	s := deque.Deque[interface{}]{}
	s.PushBack(scope)
	eval, err := jsonte.Eval(text, s, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func TestRangeOperator(t *testing.T) {
	eval := evaluate(t, "1..10")
	assertArray(t, eval, utils.JsonArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestAddition(t *testing.T) {
	eval := evaluate(t, "1 + 2")
	assertNumber(t, eval, 3)
}

func TestSubtraction(t *testing.T) {
	eval := evaluate(t, "1 - 2")
	assertNumber(t, eval, -1)
}

func TestMultiplication(t *testing.T) {
	eval := evaluate(t, "2 * 3")
	assertNumber(t, eval, 6)
}

func TestIntegerDivision(t *testing.T) {
	eval := evaluate(t, "6 / 4")
	assertNumber(t, eval, 1)
}

func TestFloatDivision(t *testing.T) {
	eval := evaluate(t, "6.0 / 4")
	assertNumber(t, eval, 1.5)
}

func TestOperationOrder(t *testing.T) {
	eval := evaluate(t, "2 + 2 * 2")
	assertNumber(t, eval, 6)
}

func TestOperationOrder2(t *testing.T) {
	eval := evaluate(t, "2 * 2 + 2")
	assertNumber(t, eval, 6)
}

func TestOperationOrder3(t *testing.T) {
	eval := evaluate(t, "2 * (2 + 2)")
	assertNumber(t, eval, 8)
}

func TestEquality(t *testing.T) {
	eval := evaluate(t, "1 == 1")
	assertBool(t, eval, true)
}

func TestInequality(t *testing.T) {
	eval := evaluate(t, "1 != 1")
	assertBool(t, eval, false)
}

func TestLessThan(t *testing.T) {
	eval := evaluate(t, "1 < 2")
	assertBool(t, eval, true)
}

func TestLessThanOrEqual(t *testing.T) {
	eval := evaluate(t, "1 <= 2")
	assertBool(t, eval, true)
}

func TestGreaterThan(t *testing.T) {
	eval := evaluate(t, "2 > 1")
	assertBool(t, eval, true)
}

func TestGreaterThanOrEqual(t *testing.T) {
	eval := evaluate(t, "2 >= 1")
	assertBool(t, eval, true)
}

func TestAnd(t *testing.T) {
	eval := evaluate(t, "true && true")
	assertBool(t, eval, true)
}

func TestOr(t *testing.T) {
	eval := evaluate(t, "true || false")
	assertBool(t, eval, true)
}

func TestNot(t *testing.T) {
	eval := evaluate(t, "!true")
	assertBool(t, eval, false)
}

func TestNot2(t *testing.T) {
	eval := evaluate(t, "!false")
	assertBool(t, eval, true)
}

func TestNot3(t *testing.T) {
	eval := evaluate(t, "!!true")
	assertBool(t, eval, true)
}

func TestNot4(t *testing.T) {
	eval := evaluate(t, "!!false")
	assertBool(t, eval, false)
}

func TestArrayAccess(t *testing.T) {
	eval := evaluate(t, "[1, 2, 3][1]")
	assertNumber(t, eval, 2)
}

func TestObjectAccess(t *testing.T) {
	eval := evaluate(t, `{"a": 1, "b": 2}["b"]`)
	assertNumber(t, eval, 2)
}

func TestObjectAccess2(t *testing.T) {
	eval := evaluate(t, `{"a": 1, "b": 2}.b`)
	assertNumber(t, eval, 2)
}

func TestObjectAccess3(t *testing.T) {
	eval := evaluateWithScope(t, `b`, utils.JsonObject{"a": utils.ToNumber(1), "b": utils.ToNumber(2)})
	assertNumber(t, eval, 2)
}
