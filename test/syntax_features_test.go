package test

import (
	"strings"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
)

func assertErrorWithScope(t *testing.T, text string, scope utils.NavigableMap[string, interface{}], expectedParts ...string) {
	t.Helper()
	s := deque.Deque[*types.JsonObject]{}
	s.PushBack(types.AsObject(scope))
	_, err := jsonte.Eval(text, s, "#")
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	for _, part := range expectedParts {
		if !strings.Contains(err.Error(), part) {
			t.Fatalf("Error %q does not contain %q", err.Error(), part)
		}
	}
}

// Tests for syntax features: object shorthand, colonless ternary, operator
// precedence, unary plus, extended number literals, array slicing, the `in`
// operator, computed object keys, compound assignments, increment/decrement
// and single-statement control-flow bodies.

func syntaxScope() utils.NavigableMap[string, interface{}] {
	return utils.ToNavigableMap(
		"a", 7,
		"b", 3,
		"obj", utils.ToNavigableMap("x", 1, "y", 2),
		"arr", []interface{}{10, 20, 30, 40},
	)
}

// 1. Object literal shorthand {a} == {a: a}

func TestObjectShorthand(t *testing.T) {
	eval := evaluateWithScope(t, "{a, b}", syntaxScope())
	expected := types.AsObject(utils.ToNavigableMap("a", 7, "b", 3))
	assertObject(t, eval, expected)
}

func TestObjectShorthandInvalid(t *testing.T) {
	s := syntaxScope()
	assertErrorWithScope(t, "{a + b}", s, "Object literal entry 'a+b' must be 'key: value', a spread, or a plain name")
}

// 2. Colonless ternary a ? b

func TestColonlessTernaryTruthy(t *testing.T) {
	eval := evaluate(t, "true ? 42")
	assertNumber(t, eval, 42)
}

func TestColonlessTernaryFalsy(t *testing.T) {
	eval := evaluate(t, "false ? 42")
	assertNull(t, eval)
}

// 3. Operator precedence: relational binds tighter than equality, ?? loosest

func TestRelationalEqualityPrecedence(t *testing.T) {
	eval := evaluate(t, "1 < 2 == true")
	assertBool(t, eval, true)
}

func TestRelationalEqualityPrecedence2(t *testing.T) {
	eval := evaluate(t, "3 > 2 == 2 > 1")
	assertBool(t, eval, true)
}

func TestNullCoalescingPrecedence(t *testing.T) {
	eval := evaluate(t, "null ?? 1 == 2")
	assertBool(t, eval, false)
}

// 6. Unary plus

func TestUnaryPlus(t *testing.T) {
	eval := evaluate(t, "+5")
	assertNumber(t, eval, 5)
}

func TestUnaryPlusExpression(t *testing.T) {
	eval := evaluate(t, "1 + +5")
	assertNumber(t, eval, 6)
}

// 7. Number literal formats

func TestScientificNotation(t *testing.T) {
	eval := evaluate(t, "1e5")
	assertNumber(t, eval, 100000)
}

func TestScientificNotationNegativeExponent(t *testing.T) {
	eval := evaluate(t, "25e-2")
	assertNumber(t, eval, 0.25)
}

func TestScientificNotationWithFraction(t *testing.T) {
	eval := evaluate(t, "1.5e2")
	assertNumber(t, eval, 150)
}

func TestHexLiteral(t *testing.T) {
	eval := evaluate(t, "0x1F")
	assertNumber(t, eval, 31)
}

func TestLeadingDotFloat(t *testing.T) {
	eval := evaluate(t, ".5 + .25")
	assertNumber(t, eval, 0.75)
}

func TestRangeStillWorksWithNewNumbers(t *testing.T) {
	eval := evaluate(t, "1..3")
	assertArray(t, eval, types.Box([]interface{}{1, 2, 3}).(*types.JsonArray))
}

// 8. Array slicing with a range index

func TestArraySliceRange(t *testing.T) {
	eval := evaluateWithScope(t, "arr[1..2]", syntaxScope())
	assertArray(t, eval, types.Box([]interface{}{20, 30}).(*types.JsonArray))
}

func TestArraySliceNegative(t *testing.T) {
	eval := evaluateWithScope(t, "arr[-2..-1]", syntaxScope())
	assertArray(t, eval, types.Box([]interface{}{30, 40}).(*types.JsonArray))
}

func TestArraySliceOutOfBounds(t *testing.T) {
	s := syntaxScope()
	assertErrorWithScope(t, "arr[2..10]", s, "Cannot access arr[2..10]", "Index out of bounds: 4")
}

// 9. `in` operator

func TestInOperatorObject(t *testing.T) {
	eval := evaluateWithScope(t, "'x' in obj", syntaxScope())
	assertBool(t, eval, true)
	eval = evaluateWithScope(t, "'z' in obj", syntaxScope())
	assertBool(t, eval, false)
}

func TestInOperatorArray(t *testing.T) {
	eval := evaluateWithScope(t, "20 in arr", syntaxScope())
	assertBool(t, eval, true)
	eval = evaluateWithScope(t, "25 in arr", syntaxScope())
	assertBool(t, eval, false)
}

func TestInOperatorString(t *testing.T) {
	eval := evaluate(t, "'ell' in 'hello'")
	assertBool(t, eval, true)
}

func TestInOperatorPrecedence(t *testing.T) {
	// `in` binds tighter than equality
	eval := evaluateWithScope(t, "'x' in obj == true", syntaxScope())
	assertBool(t, eval, true)
}

// 10. Computed object keys

func TestComputedObjectKey(t *testing.T) {
	eval := evaluateWithScope(t, "{['key_' + b]: a}", syntaxScope())
	expected := types.AsObject(utils.ToNavigableMap("key_3", 7))
	assertObject(t, eval, expected)
}

func TestComputedObjectKeyNumber(t *testing.T) {
	eval := evaluate(t, "{[1 + 2]: 'v'}")
	expected := types.AsObject(utils.ToNavigableMap("3", "v"))
	assertObject(t, eval, expected)
}

// 13. Compound assignments

func TestSubtractAssign(t *testing.T) {
	eval := evaluateScript(t, "x = 10; x -= 3; return x;")
	assertReturn(t, eval, types.AsNumber(7))
}

func TestMultiplyAssign(t *testing.T) {
	eval := evaluateScript(t, "x = 10; x *= 3; return x;")
	assertReturn(t, eval, types.AsNumber(30))
}

func TestDivideAssign(t *testing.T) {
	eval := evaluateScript(t, "x = 10; x /= 3; return x;")
	assertReturn(t, eval, types.AsNumber(3))
}

func TestDivideAssignFloat(t *testing.T) {
	eval := evaluateScript(t, "x = 10.0; x /= 4; return x;")
	assertReturn(t, eval, types.AsNumber(2.5))
}

func TestModuloAssign(t *testing.T) {
	eval := evaluateScript(t, "x = 10; x %= 3; return x;")
	assertReturn(t, eval, types.AsNumber(1))
}

func TestCompoundAssignArrayElement(t *testing.T) {
	eval := evaluateScript(t, "x = [1, 2, 3]; x[1] *= 10; return x[1];")
	assertReturn(t, eval, types.AsNumber(20))
}

// 14. Increment / decrement

func TestIncrement(t *testing.T) {
	eval := evaluateScript(t, "x = 1; x++; return x;")
	assertReturn(t, eval, types.AsNumber(2))
}

func TestDecrement(t *testing.T) {
	eval := evaluateScript(t, "x = 1; x--; return x;")
	assertReturn(t, eval, types.AsNumber(0))
}

func TestIncrementInLoop(t *testing.T) {
	eval := evaluateScript(t, "x = 0; while x < 5 { x++; } return x;")
	assertReturn(t, eval, types.AsNumber(5))
}

// 15. Single-statement control-flow bodies

func TestSingleStatementIf(t *testing.T) {
	eval := evaluateScript(t, "x = 0; if true x = 5; return x;")
	assertReturn(t, eval, types.AsNumber(5))
}

func TestSingleStatementIfElse(t *testing.T) {
	eval := evaluateScript(t, "if false return 1; else return 2;")
	assertReturn(t, eval, types.AsNumber(2))
}

func TestSingleStatementFor(t *testing.T) {
	eval := evaluateScript(t, "for i in 1..3 if i == 2 return i * 10; return 0;")
	assertReturn(t, eval, types.AsNumber(20))
}

func TestSingleStatementWhile(t *testing.T) {
	eval := evaluateScript(t, "x = 0; while x < 3 x++; return x;")
	assertReturn(t, eval, types.AsNumber(3))
}

func TestBlockBodiesStillWork(t *testing.T) {
	eval := evaluateScript(t, "x = 0; while x < 3 { x += 1; } if x == 3 { return x; } else { return 0; }")
	assertReturn(t, eval, types.AsNumber(3))
}
