package test

import (
	"errors"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"reflect"
	"testing"
)

func assertAction(t *testing.T, eval jsonte.Result, action types.JsonAction) {
	if eval.Action != action {
		t.Errorf("Action is not %d", action)
	}
}

func assertArray(t *testing.T, eval jsonte.Result, expected *types.JsonArray) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsArray(eval.Value) {
		t.Fatalf("Result is not an array (%s)", reflect.TypeOf(eval.Value).Name())
	}
	array, ok := eval.Value.(*types.JsonArray)
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
			compareJsonObject(t, expected.Value[i].(*types.JsonObject), array.Value[i].(*types.JsonObject), fmt.Sprintf("#[%d]", i), true)
		} else if types.IsArray(expected.Value[i]) {
			if array.Value[i] == nil {
				t.Fatalf("Array element %d is null", i)
			}
			if !types.IsArray(array.Value[i]) {
				t.Fatalf("Array element %d is not an array (%s)", i, reflect.TypeOf(array.Value[i]).Name())
			}
			compareJsonArray(t, expected.Value[i].(*types.JsonArray), array.Value[i].(*types.JsonArray), fmt.Sprintf("#[%d]", i))
		} else if !array.Value[i].Equals(expected.Value[i]) {
			t.Errorf("Array element %d is not correct (expected %s, got %s)", i, types.ToString(expected.Value[i]), types.ToString(array.Value[i]))
		}
	}
}

func assertObject(t *testing.T, eval jsonte.Result, expected *types.JsonObject) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsObject(eval.Value) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	obj, ok := eval.Value.(*types.JsonObject)
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	compareJsonObject(t, expected, obj, "#", true)
}

func assertObjectContains(t *testing.T, eval jsonte.Result, expected *types.JsonObject) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsObject(eval.Value) {
		t.Fatalf("Result is not an object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	obj, ok := eval.Value.(*types.JsonObject)
	if !ok {
		t.Fatalf("Result is not a JSON object (%s)", reflect.TypeOf(eval.Value).Name())
	}
	compareJsonObject(t, expected, obj, "#", false)
}

func assertNumber(t *testing.T, eval jsonte.Result, expected float64) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsNumber(eval.Value) {
		t.Fatalf("Result is not a number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	number, ok := eval.Value.(*types.JsonNumber)
	if !ok {
		t.Fatalf("Result is not a JSON number (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if number.FloatValue() != expected {
		t.Fatalf("Result is not correct (expected %f, got %f)", expected, number.FloatValue())
	}
}

func assertSemver(t *testing.T, eval jsonte.Result, expected types.Semver) {
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

func assertJsonPath(t *testing.T, eval jsonte.Result, expected types.JsonPath) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if !types.IsJsonPath(eval.Value) {
		t.Fatalf("Result is not a JSON Path (%s)", reflect.TypeOf(eval.Value).Name())
	}
	path := types.AsJsonPath(eval.Value)
	if path.StringValue() != expected.StringValue() {
		t.Fatalf("Result is not correct (expected %s, got %s)", expected.StringValue(), path.StringValue())
	}
}

func assertBool(t *testing.T, eval jsonte.Result, expected bool) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := eval.Value.(*types.JsonBool); !ok {
		t.Fatalf("Result is not a boolean (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if eval.Value.(*types.JsonBool).BoolValue() != expected {
		t.Fatalf("Result is not correct (expected %t, got %s)", expected, eval.Value.StringValue())
	}
}

func assertString(t *testing.T, eval jsonte.Result, expected string) {
	assertAction(t, eval, types.Value)
	if eval.Value == nil {
		t.Fatalf("Result is null")
	}
	if _, ok := eval.Value.(*types.JsonString); !ok {
		t.Fatalf("Result is not a string (%s)", reflect.TypeOf(eval.Value).Name())
	}
	if eval.Value.(*types.JsonString).StringValue() != expected {
		t.Fatalf("Result is not correct (expected %s, got %s)", expected, eval.Value.StringValue())
	}
}

func assertNull(t *testing.T, eval jsonte.Result) {
	assertAction(t, eval, types.Value)
	if !types.IsNull(eval.Value) {
		t.Fatalf("Result is not null (%s)", types.ToString(eval.Value))
	}
}

func assertError(t *testing.T, text string, error []string) {
	eval, err := jsonte.QuickEval(text, "#")
	assertAction(t, eval, types.Value)
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

func evaluate(t *testing.T, text string) jsonte.Result {
	eval, err := jsonte.QuickEval(text, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func evaluateWithScope(t *testing.T, text string, scope utils.NavigableMap[string, interface{}]) jsonte.Result {
	s := deque.Deque[*types.JsonObject]{}
	s.PushBack(types.AsObject(scope))
	eval, err := jsonte.Eval(text, s, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func evaluateWithObjectScope(t *testing.T, text string, scope *types.JsonObject) jsonte.Result {
	s := deque.Deque[*types.JsonObject]{}
	s.PushBack(scope)
	eval, err := jsonte.Eval(text, s, "#")
	if err != nil {
		t.Fatal(err)
	}
	return eval
}

func TestRangeOperator(t *testing.T) {
	eval := evaluate(t, "1..10")
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).(*types.JsonArray))
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
		"Index out of bounds: 3",
	}
	assertError(t, expression, expectedError)
}

func TestObjectPlusOperator(t *testing.T) {
	expression := evaluate(t, `{'a': 1} + {'b': 2}`)
	assertObject(t, expression, types.AsObject(utils.ToNavigableMap(
		"b", 2,
		"a", 1,
	)))
}

func TestObjectPlusOperator2(t *testing.T) {
	expression := evaluate(t, `null + {'b': 2}`)
	assertObject(t, expression, types.AsObject(map[string]interface{}{
		"b": 2,
	}))
}

func TestObjectPlusOperator3(t *testing.T) {
	expression := evaluate(t, `undefined + {'b': 2}`)
	assertObject(t, expression, types.AsObject(map[string]interface{}{
		"b": 2,
	}))
}

func TestObjectPlusOperator4(t *testing.T) {
	expression := evaluate(t, `{'a': 1} + {'a': 2}`)
	assertObject(t, expression, types.AsObject(map[string]interface{}{
		"a": 1,
	}))
}

func TestArrayPlusOperator(t *testing.T) {
	expression := evaluate(t, `[1] + [2]`)
	assertArray(t, expression, types.AsArray([]interface{}{1, 2}))
}

func TestArrayPlusOperator2(t *testing.T) {
	expression := evaluate(t, `null + [2]`)
	assertArray(t, expression, types.AsArray([]interface{}{2}))
}

func TestArrayPlusOperator3(t *testing.T) {
	expression := evaluate(t, `undefined + [2]`)
	assertArray(t, expression, types.AsArray([]interface{}{2}))
}

func TestOrShortCircuit(t *testing.T) {
	eval := evaluateWithScope(t, `true || error()`, utils.ToNavigableMap("error", types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		t.Fatalf("error() should not be called")
		return nil, nil
	}, "error", []string{}, []string{})))
	assertBool(t, eval, true)
}

func TestOrShortCircuit2(t *testing.T) {
	eval := evaluateWithScope(t, `!false || error()`, utils.ToNavigableMap("error", types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		t.Fatalf("error() should not be called")
		return nil, nil
	}, "error", []string{}, []string{})))
	assertBool(t, eval, true)
}

func TestAndShortCircuit(t *testing.T) {
	eval := evaluateWithScope(t, `false && error()`, utils.ToNavigableMap("error", types.NewLambda(func(this *types.JsonLambda, args []types.JsonType) (types.JsonType, error) {
		t.Fatalf("error() should not be called")
		return nil, nil
	}, "error", []string{}, []string{})))
	assertBool(t, eval, false)
}

func TestCaseSensitivity(t *testing.T) {
	eval := evaluate(t, `(1..10).Count()`)
	assertNumber(t, eval, 10)
}

func TestEncoding(t *testing.T) {
	eval := evaluateWithScope(t, `test()`, utils.ToNavigableMap("test", types.NewString("() => '§'")))
	assertString(t, eval, "§")
}

func TestSpreadArray(t *testing.T) {
	eval := evaluate(t, `[1, 2, ...[3, 4], 5]`)
	assertArray(t, eval, types.AsArray([]interface{}{1, 2, 3, 4, 5}))
}

func TestSpreadArray2(t *testing.T) {
	eval := evaluate(t, `[...[1, 2], 3, 4, 5]`)
	assertArray(t, eval, types.AsArray([]interface{}{1, 2, 3, 4, 5}))
}

func TestSpreadArray3(t *testing.T) {
	eval := evaluate(t, `[1, 2, ...[3, 4], ...[5, 6]]`)
	assertArray(t, eval, types.AsArray([]interface{}{1, 2, 3, 4, 5, 6}))
}

func TestSpreadArray4(t *testing.T) {
	eval := evaluate(t, `[...[1, 2], ...[3, 4], 5]`)
	assertArray(t, eval, types.AsArray([]interface{}{1, 2, 3, 4, 5}))
}

func TestSpreadArray5(t *testing.T) {
	eval := evaluate(t, `[...[1, 2], ...[3, 4], ...[5, 6]]`)
	assertArray(t, eval, types.AsArray([]interface{}{1, 2, 3, 4, 5, 6}))
}

func TestSpreadObject(t *testing.T) {
	eval := evaluate(t, `{'a': 1, ...{'b': 2}, 'c': 3}`)
	assertObject(t, eval, types.AsObject(utils.ToNavigableMap(
		"a", 1,
		"b", 2,
		"c", 3,
	)))
}

func TestSpreadObject2(t *testing.T) {
	eval := evaluate(t, `{'a': 1, ...{'a': 2}, 'c': 3}`)
	assertObject(t, eval, types.AsObject(utils.ToNavigableMap(
		"a", 2,
		"c", 3,
	)))
}

func TestSpreadObject3(t *testing.T) {
	eval := evaluate(t, `{'a': 1, ...{'b': 2}, ...{'c': 3}}`)
	assertObject(t, eval, types.AsObject(utils.ToNavigableMap(
		"a", 1,
		"b", 2,
		"c", 3,
	)))
}

func TestLambdaInfo(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedVars  []string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "single_variable",
			input:         `x => x`,
			expectedVars:  []string{"x"},
			expectedArgs:  []string{"x"},
			expectedError: nil,
		},
		{
			name:          "two_variables",
			input:         `(x, y) => [x, y]`,
			expectedVars:  []string{"x", "y"},
			expectedArgs:  []string{"x", "y"},
			expectedError: nil,
		},
		{
			name:          "no_variables",
			input:         `() => null`,
			expectedVars:  []string{},
			expectedArgs:  []string{},
			expectedError: nil,
		},
		{
			name:          "outside_variables",
			input:         `() => x`,
			expectedVars:  []string{"x"},
			expectedArgs:  []string{},
			expectedError: nil,
		},
		{
			name:          "embedded_lambda",
			input:         `x => !arr2.any(y => x.startsWith(y))`,
			expectedVars:  []string{"arr2", "x", "y"},
			expectedArgs:  []string{"x", "y"},
			expectedError: nil,
		},
	}

	for idx, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vars, args, err := jsonte.ParseLambda(tc.input)

			if !errors.Is(err, tc.expectedError) {
				t.Fatalf("Test case %d (%s) - Expected error: %v, got: %v", idx, tc.name, tc.expectedError, err)
			}

			if len(args) != len(tc.expectedArgs) {
				t.Fatalf("Test case %d (%s) - Expected %d argument(s), got: %d", idx, tc.name, len(tc.expectedArgs), len(args))
			}

			if len(vars) != len(tc.expectedVars) {
				t.Fatalf("Test case %d (%s) - Expected %d variable(s), got: %d", idx, tc.name, len(tc.expectedVars), len(vars))
			}

			for i, arg := range args {
				if arg != tc.expectedArgs[i] {
					t.Fatalf("Test case %d (%s) - Expected argument '%s', got: '%s'", idx, tc.name, tc.expectedArgs[i], arg)
				}
			}

			for i, v := range vars {
				if v != tc.expectedVars[i] {
					t.Fatalf("Test case %d (%s) - Expected variable '%s', got: '%s'", idx, tc.name, tc.expectedVars[i], v)
				}
			}
		})
	}
}

func TestEqualsObjects(t *testing.T) {
	eval := evaluate(t, `{"a": 1} == {"a": 1}`)
	assertBool(t, eval, true)
}

func TestEqualsObjects2(t *testing.T) {
	eval := evaluate(t, `{"a": 1} == {"a": 2}`)
	assertBool(t, eval, false)
}

func TestEqualsObjects3(t *testing.T) {
	eval := evaluate(t, `{"a": 1} == {"b": 1}`)
	assertBool(t, eval, false)
}

func TestEqualsNestedObjects(t *testing.T) {
	eval := evaluate(t, `{"a": {"b": 1}} == {"a": {"b": 1}}`)
	assertBool(t, eval, true)
}

func TestEqualsNestedObjects2(t *testing.T) {
	eval := evaluate(t, `{"a": {"b": 1}} == {"a": {"b": 2}}`)
	assertBool(t, eval, false)
}

func TestEqualsArrays(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3] == [1, 2, 3]`)
	assertBool(t, eval, true)
}

func TestEqualsArrays2(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3] == [1, 2, 4]`)
	assertBool(t, eval, false)
}

func TestEqualsNestedArrays(t *testing.T) {
	eval := evaluate(t, `[[1, 2], [3, 4]] == [[1, 2], [3, 4]]`)
	assertBool(t, eval, true)
}

func TestEqualsNestedArrays2(t *testing.T) {
	eval := evaluate(t, `[[1, 2], [3, 4]] == [[1, 2], [3, 5]]`)
	assertBool(t, eval, false)
}

func TestEqualsMixed(t *testing.T) {
	eval := evaluate(t, `[1, 2, {"a": 1}] == [1, 2, {"a": 1}]`)
	assertBool(t, eval, true)
}

func TestEqualsStrings(t *testing.T) {
	eval := evaluate(t, `"hello" == "hello"`)
	assertBool(t, eval, true)
}

func TestEqualsStrings2(t *testing.T) {
	eval := evaluate(t, `"hello" == "world"`)
	assertBool(t, eval, false)
}

func TestNullCoalesceOnIndexAccessor(t *testing.T) {
	eval := evaluateWithScope(t, `a?['asd'] ?? false`, utils.ToNavigableMap("a", types.NewJsonObject()))
	assertBool(t, eval, false)
}

func TestNullCoalesce(t *testing.T) {
	eval := evaluateWithScope(t, `a?.asd ?? false`, utils.ToNavigableMap("a", types.NewJsonObject()))
	assertBool(t, eval, false)
}

func TestExpressionCache(t *testing.T) {
	jsonte.ClearExpressionCache()
	if jsonte.ExpressionCacheSize() != 0 {
		t.Fatalf("cache should be empty")
	}
	evaluate(t, "1 + 2")
	if jsonte.ExpressionCacheSize() != 1 {
		t.Fatalf("expected cache size 1 after first eval, got %d", jsonte.ExpressionCacheSize())
	}
	evaluate(t, "1 + 2")
	if jsonte.ExpressionCacheSize() != 1 {
		t.Fatalf("cache size should remain 1 when expression is reused, got %d", jsonte.ExpressionCacheSize())
	}
	evaluate(t, "2 + 2")
	if jsonte.ExpressionCacheSize() != 2 {
		t.Fatalf("cache size should grow for new expressions, got %d", jsonte.ExpressionCacheSize())
	}
}
