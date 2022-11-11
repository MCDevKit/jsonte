package test

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"reflect"
	"strings"
	"testing"
)

func assertAction(t *testing.T, eval jsonte.Result, action types.JsonAction) {
	t.Helper()
	if eval.Action != action {
		t.Errorf("Action is not %d", action)
	}
}

func assertArray(t *testing.T, eval jsonte.Result, expected types.JsonArray) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsArray(eval.Value) {
		t.Fatalf("Result is not an array (%s)", reflect.TypeOf(eval.Value).Name())
	}
	array, ok := eval.Value.(types.JsonArray)
	if !ok {
		t.Fatalf("Result is not a JSON array (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if len(array.Value) != len(expected.Value) {
		t.Fatalf("Array length is not correct (expected %d, got %d)", len(expected.Value), len(array.Value))
	}
	for i := 0; i < len(expected.Value); i++ {
		if types.IsNumber(expected.Value[i]) && types.IsNumber(array.Value[i]) {
			if types.AsNumber(expected.Value[i]).FloatValue() != types.AsNumber(array.Value[i]).FloatValue() {
				t.Fatalf("Array element %d is not correct (expected %f, got %f)", i, expected.Value[i], array.Value[i])
			}
		} else if types.IsObject(expected.Value[i]) {
			if array.Value[i] == nil {
				t.Fatalf("Array element %d is null", i)
			}
			if !types.IsObject(array.Value[i]) {
				t.Fatalf("Array element %d is not an object (%s)", i, reflect.TypeOf(array.Value[i]).Name())
			}
			compareJsonObject(t, expected.Value[i].(types.JsonObject), array.Value[i].(types.JsonObject), fmt.Sprintf("#[%d]", i), true)
		} else if types.IsArray(expected.Value[i]) {
			if array.Value[i] == nil {
				t.Fatalf("Array element %d is null", i)
			}
			if !types.IsArray(array.Value[i]) {
				t.Fatalf("Array element %d is not an array (%s)", i, reflect.TypeOf(array.Value[i]).Name())
			}
			compareJsonArray(t, expected.Value[i].(types.JsonArray), array.Value[i].(types.JsonArray), fmt.Sprintf("#[%d]", i))
		} else if !array.Value[i].Equals(expected.Value[i]) {
			t.Errorf("Array element %d is not correct (expected %s, got %s)", i, types.ToString(expected.Value[i]), types.ToString(array.Value[i]))
		}
	}
}

func assertObject(t *testing.T, eval jsonte.Result, expected types.JsonObject) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsObject(eval.Value) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	obj, ok := eval.Value.(types.JsonObject)
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	compareJsonObject(t, expected, obj, "#", true)
}

func assertObjectContains(t *testing.T, eval jsonte.Result, expected types.JsonObject) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsObject(eval.Value) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	obj, ok := eval.Value.(types.JsonObject)
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	compareJsonObject(t, expected, obj, "#", false)
}

func assertNumber(t *testing.T, eval jsonte.Result, expected float64) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsNumber(eval.Value) {
		t.Fatalf("Result is not a number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	number, ok := eval.Value.(types.JsonNumber)
	if !ok {
		t.Fatalf("Result is not a JSON number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if number.FloatValue() != expected {
		t.Fatalf("Result is not correct (expected %f, got %f)", expected, number.FloatValue())
	}
}

func assertSemver(t *testing.T, eval jsonte.Result, expected types.Semver) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsSemver(eval.Value) {
		t.Fatalf("Result is not a semver (%s)", reflect.TypeOf(eval.Value).Name())
	}
	semver := types.AsSemver(eval.Value)
	if semver.Major != expected.Major || semver.Minor != expected.Minor || semver.Patch != expected.Patch {
		t.Fatalf("Result is not correct (expected %s, got %s)", expected.StringValue(), semver.StringValue())
	}
}

func assertBool(t *testing.T, eval jsonte.Result, expected bool) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := eval.Value.(types.JsonBool); !ok {
		t.Fatalf("Result is not a boolean (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if eval.Value.(types.JsonBool).BoolValue() != expected {
		t.Fatalf("Result is not correct (expected %t, got %t)", expected, eval.Value)
	}
}

func assertString(t *testing.T, eval jsonte.Result, expected string) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := eval.Value.(types.JsonString); !ok {
		t.Fatalf("Result is not a string (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if eval.Value.(types.JsonString).StringValue() != expected {
		t.Fatalf("Result is not correct (expected %s, got %s)", expected, eval.Value)
	}
}

func assertNull(t *testing.T, eval jsonte.Result) {
	t.Helper()
	assertAction(t, eval, types.Value)
	if eval.Value != types.Null {
		t.Fatalf("Result is not null (%s)", types.ToString(eval.Value))
	}
}

func assertError(t *testing.T, text string, error []string) {
	t.Helper()
	eval, err := jsonte.QuickEval(text, "#")
	assertAction(t, eval, types.Value)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	split := strings.Split(err.Error(), "\n")
	for i := 0; i < len(split); i++ {
		split[i] = strings.TrimSpace(split[i])
	}
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
}

func evaluate(t *testing.T, text string) jsonte.Result {
	t.Helper()
	eval, err := jsonte.QuickEval(text, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func evaluateWithScope(t *testing.T, text string, scope utils.NavigableMap[string, interface{}]) jsonte.Result {
	t.Helper()
	s := deque.Deque[types.JsonObject]{}
	s.PushBack(types.AsObject(scope))
	eval, err := jsonte.Eval(text, s, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func TestRangeOperator(t *testing.T) {
	eval := evaluate(t, "1..10")
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(types.JsonArray))
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

func TestScope(t *testing.T) {
	eval := evaluateWithScope(t, `b`, utils.ToNavigableMap("a", types.AsNumber(1), "b", types.AsNumber(2)))
	assertNumber(t, eval, 2)
}

func TestConcatenation(t *testing.T) {
	eval := evaluate(t, `'a' + 'b'`)
	assertString(t, eval, "ab")
}

func TestConcatenation2(t *testing.T) {
	eval := evaluate(t, `'a' + 1`)
	assertString(t, eval, "a1")
}

func TestConcatenation3(t *testing.T) {
	eval := evaluate(t, `1 + 'b'`)
	assertString(t, eval, "1b")
}

func TestConcatenation4(t *testing.T) {
	eval := evaluate(t, `1 + 2 + 'b'`)
	assertString(t, eval, "3b")
}

func TestConcatenation5(t *testing.T) {
	eval := evaluate(t, `'a' + 1 + 2`)
	assertString(t, eval, "a12")
}

func TestConcatenation6(t *testing.T) {
	eval := evaluate(t, `'a' + (1 + 2)`)
	assertString(t, eval, "a3")
}

func TestNullCoalescing(t *testing.T) {
	eval := evaluate(t, `1 ?? 2`)
	assertNumber(t, eval, 1)
}

func TestNullCoalescing2(t *testing.T) {
	eval := evaluate(t, `null ?? 2`)
	assertNumber(t, eval, 2)
}

func TestNullCoalescing3(t *testing.T) {
	eval := evaluate(t, `null ?? null`)
	assertNull(t, eval)
}

func TestNullCoalescing4(t *testing.T) {
	eval := evaluate(t, `null ?? null ?? 3`)
	assertNumber(t, eval, 3)
}

func TestTernaryOperator(t *testing.T) {
	eval := evaluate(t, `true ? 1 : 2`)
	assertNumber(t, eval, 1)
}

func TestTernaryOperator2(t *testing.T) {
	eval := evaluate(t, `false ? 1 : 2`)
	assertNumber(t, eval, 2)
}

func TestTernaryOperator3(t *testing.T) {
	eval := evaluate(t, `true ? 1 : 2 + 3`)
	assertNumber(t, eval, 1)
}

func TestTernaryOperator4(t *testing.T) {
	eval := evaluate(t, `false ? 1 : 2 + 3`)
	assertNumber(t, eval, 5)
}

func TestTernaryOperator5(t *testing.T) {
	eval := evaluate(t, `true ? 1 + 2 : 3`)
	assertNumber(t, eval, 3)
}

func TestTernaryOperator6(t *testing.T) {
	eval := evaluate(t, `false ? 1 + 2 : 3`)
	assertNumber(t, eval, 3)
}

func TestTernaryOperator7(t *testing.T) {
	eval := evaluate(t, `false ? 1 : true ? 2 : 3`)
	assertNumber(t, eval, 2)
}

func TestTernaryOperator8(t *testing.T) {
	eval := evaluate(t, `false ? 1 : false ? 2 : 3`)
	assertNumber(t, eval, 3)
}

func TestNullForgivingObjectOperator(t *testing.T) {
	eval := evaluate(t, `{}?.a`)
	assertNull(t, eval)
}

func TestNullForgivingObjectOperator2(t *testing.T) {
	eval := evaluate(t, `{"a": 1}?.a`)
	assertNumber(t, eval, 1)
}

func TestNullForgivingArrayOperator(t *testing.T) {
	eval := evaluate(t, `[]?[0]`)
	assertNull(t, eval)
}

func TestNullForgivingArrayOperator2(t *testing.T) {
	eval := evaluate(t, `[1]?[0]`)
	assertNumber(t, eval, 1)
}

func TestNegativeArrayIndex(t *testing.T) {
	eval := evaluate(t, `(1..10)[-1]`)
	assertNumber(t, eval, 10)
}

func TestNegativeStringIndex(t *testing.T) {
	eval := evaluate(t, `'hello'[-1]`)
	assertString(t, eval, "o")
}

func TestFunctionNotFound(t *testing.T) {
	expression := `spilt('hello world', ' ')`
	expectedError := []string{
		"Function 'spilt' not found, did you mean 'split'?",
	}
	assertError(t, expression, expectedError)
}

func TestOutOfBoundsAccess(t *testing.T) {
	expression := `'abc'[3]`
	expectedError := []string{
		"Cannot access 'abc'[3]",
		"[+]: Index out of bounds: 3",
	}
	assertError(t, expression, expectedError)
}
