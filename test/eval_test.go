package test

import (
	"jsonte/jsonte"
	"jsonte/jsonte/utils"
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
		t.Fatal("Result is not an array")
	}
	array, ok := eval.Value.(utils.JsonArray)
	if !ok {
		t.Fatal("Result is not a JSON array")
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
		t.Fatal("Result is not a number")
	}
	number, ok := eval.Value.(utils.JsonNumber)
	if !ok {
		t.Fatal("Result is not a JSON number")
	}
	if number.FloatValue() != expected {
		t.Fatalf("Result is not correct (expected %f, got %f)", expected, number.FloatValue())
	}
}

func evaluate(t *testing.T, text string) jsonte.Result {
	t.Helper()
	eval, err := jsonte.QuickEval(text)
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
