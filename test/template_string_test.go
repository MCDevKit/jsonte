package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func TestTemplateStringEmpty(t *testing.T) {
	eval := evaluate(t, "``")
	assertString(t, eval, "")
}

func TestTemplateStringLiteral(t *testing.T) {
	eval := evaluate(t, "`hello world`")
	assertString(t, eval, "hello world")
}

func TestTemplateStringSingleInterp(t *testing.T) {
	eval := evaluateWithScope(t, "`hello ${name}`", utils.ToNavigableMap("name", types.NewString("world")))
	assertString(t, eval, "hello world")
}

func TestTemplateStringMultipleInterps(t *testing.T) {
	eval := evaluateWithScope(t, "`${a} + ${b} = ${a + b}`",
		utils.ToNavigableMap("a", types.AsNumber(1), "b", types.AsNumber(2)))
	assertString(t, eval, "1 + 2 = 3")
}

func TestTemplateStringArithmetic(t *testing.T) {
	eval := evaluateWithScope(t, "`result: ${x * 2 + 1}`", utils.ToNavigableMap("x", types.AsNumber(5)))
	assertString(t, eval, "result: 11")
}

func TestTemplateStringObjectInInterp(t *testing.T) {
	eval := evaluate(t, "`val: ${{key: 'v'}.key}`")
	assertString(t, eval, "val: v")
}

func TestTemplateStringDollarLiteral(t *testing.T) {
	eval := evaluate(t, "`price: $10`")
	assertString(t, eval, "price: $10")
}

func TestTemplateStringEscapeNewline(t *testing.T) {
	eval := evaluate(t, "`line1\\nline2`")
	assertString(t, eval, "line1\nline2")
}

func TestTemplateStringEscapeTab(t *testing.T) {
	eval := evaluate(t, "`col1\\tcol2`")
	assertString(t, eval, "col1\tcol2")
}

func TestTemplateStringEscapeBacktick(t *testing.T) {
	eval := evaluate(t, "`back\\`tick`")
	assertString(t, eval, "back`tick")
}

func TestTemplateStringInterpOnly(t *testing.T) {
	eval := evaluateWithScope(t, "`${name}`", utils.ToNavigableMap("name", types.NewString("Alice")))
	assertString(t, eval, "Alice")
}

func TestTemplateStringConditionInInterp(t *testing.T) {
	eval := evaluateWithScope(t, "`${x > 0 ? 'pos' : 'neg'}`", utils.ToNavigableMap("x", types.AsNumber(5)))
	assertString(t, eval, "pos")
}

func TestTemplateStringFunctionCallInInterp(t *testing.T) {
	eval := evaluate(t, "`${length('hello')}`")
	assertString(t, eval, "5")
}

func TestTemplateStringNestedBraces(t *testing.T) {
	eval := evaluate(t, "`${{a: 1, b: 2}.a}`")
	assertString(t, eval, "1")
}

func TestTemplateStringIsString(t *testing.T) {
	eval := evaluate(t, "`hello` + ' world'")
	assertString(t, eval, "hello world")
}

func TestTemplateStringInArray(t *testing.T) {
	eval := evaluateWithScope(t, "[`a${x}b`, `c${x}d`]", utils.ToNavigableMap("x", types.AsNumber(1)))
	assertArray(t, eval, &types.JsonArray{Value: []types.JsonType{types.NewString("a1b"), types.NewString("c1d")}})
}
