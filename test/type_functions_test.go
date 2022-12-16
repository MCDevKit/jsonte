package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func TestIsArray(t *testing.T) {
	eval := evaluate(t, `isArray([])`)
	assertBool(t, eval, true)
}

func TestIsArray2(t *testing.T) {
	eval := evaluate(t, `isArray(1..10)`)
	assertBool(t, eval, true)
}

func TestIsArray3(t *testing.T) {
	eval := evaluate(t, `isArray({})`)
	assertBool(t, eval, false)
}

func TestIsObject(t *testing.T) {
	eval := evaluate(t, `isObject({})`)
	assertBool(t, eval, true)
}

func TestIsObject2(t *testing.T) {
	eval := evaluate(t, `isObject([])`)
	assertBool(t, eval, false)
}

func TestIsString(t *testing.T) {
	eval := evaluate(t, `isString("test")`)
	assertBool(t, eval, true)
}

func TestIsString2(t *testing.T) {
	eval := evaluate(t, `isString(1)`)
	assertBool(t, eval, false)
}

func TestIsNumber(t *testing.T) {
	eval := evaluate(t, `isNumber(1)`)
	assertBool(t, eval, true)
}

func TestIsNumber2(t *testing.T) {
	eval := evaluate(t, `isNumber("1")`)
	assertBool(t, eval, false)
}

func TestIsNumber3(t *testing.T) {
	eval := evaluate(t, `isNumber(1.1)`)
	assertBool(t, eval, true)
}

func TestIsBoolean(t *testing.T) {
	eval := evaluate(t, `isBoolean(true)`)
	assertBool(t, eval, true)
}

func TestIsBoolean2(t *testing.T) {
	eval := evaluate(t, `isBoolean(1)`)
	assertBool(t, eval, false)
}

func TestIsSemver(t *testing.T) {
	eval := evaluate(t, `isSemver(semver("1.10.0"))`)
	assertBool(t, eval, true)
}

func TestIsSemver2(t *testing.T) {
	eval := evaluate(t, `isSemver(1)`)
	assertBool(t, eval, false)
}

func TestAsString(t *testing.T) {
	eval := evaluate(t, `asString(1)`)
	assertString(t, eval, "1")
}

func TestAsString2(t *testing.T) {
	eval := evaluate(t, `asString("1")`)
	assertString(t, eval, "1")
}

func TestAsString3(t *testing.T) {
	eval := evaluate(t, `asString(true)`)
	assertString(t, eval, "true")
}

func TestAsString4(t *testing.T) {
	eval := evaluate(t, `asString({})`)
	assertString(t, eval, "{}")
}

func TestAsString5(t *testing.T) {
	eval := evaluate(t, `asString([])`)
	assertString(t, eval, "[]")
}

func TestAsString6(t *testing.T) {
	eval := evaluate(t, `asString(null)`)
	assertString(t, eval, "null")
}

func TestAsNumber(t *testing.T) {
	eval := evaluate(t, `asNumber(1)`)
	assertNumber(t, eval, 1)
}

func TestAsNumber2(t *testing.T) {
	eval := evaluate(t, `asNumber("1")`)
	assertNumber(t, eval, 1)
}

func TestAsNumber3(t *testing.T) {
	eval := evaluate(t, `asNumber(true)`)
	assertNumber(t, eval, 1)
}

func TestAsNumber4(t *testing.T) {
	eval := evaluate(t, `asNumber({})`)
	assertNumber(t, eval, 0)
}

func TestAsNumber5(t *testing.T) {
	eval := evaluate(t, `asNumber([])`)
	assertNumber(t, eval, 0)
}

func TestAsNumber6(t *testing.T) {
	eval := evaluate(t, `asNumber(null)`)
	assertNumber(t, eval, 0)
}

func TestAsBoolean(t *testing.T) {
	eval := evaluate(t, `asBoolean(1)`)
	assertBool(t, eval, true)
}

func TestAsBoolean2(t *testing.T) {
	eval := evaluate(t, `asBoolean("1")`)
	assertBool(t, eval, true)
}

func TestAsBoolean3(t *testing.T) {
	eval := evaluate(t, `asBoolean(true)`)
	assertBool(t, eval, true)
}

func TestAsBoolean4(t *testing.T) {
	eval := evaluate(t, `asBoolean({})`)
	assertBool(t, eval, false)
}

func TestAsBoolean5(t *testing.T) {
	eval := evaluate(t, `asBoolean([])`)
	assertBool(t, eval, false)
}

func TestAsBoolean6(t *testing.T) {
	eval := evaluate(t, `asBoolean(null)`)
	assertBool(t, eval, false)
}

func TestAsBoolean7(t *testing.T) {
	eval := evaluate(t, `asBoolean("")`)
	assertBool(t, eval, false)
}

func TestAsBoolean8(t *testing.T) {
	eval := evaluate(t, `asBoolean(0)`)
	assertBool(t, eval, false)
}

func TestAsBoolean9(t *testing.T) {
	eval := evaluate(t, `asBoolean(false)`)
	assertBool(t, eval, false)
}

func TestAsBoolean10(t *testing.T) {
	eval := evaluate(t, `asBoolean([1])`)
	assertBool(t, eval, true)
}

func TestAsBoolean11(t *testing.T) {
	eval := evaluate(t, `asBoolean({a: 1})`)
	assertBool(t, eval, true)
}

func TestParseArray(t *testing.T) {
	eval := evaluate(t, `parseArray('["a", "b", "c"]')`)
	assertArray(t, eval, types.AsArray([]interface{}{"a", "b", "c"}))
}

func TestParseArray2(t *testing.T) {
	assertError(t, `parseArray('{"a": "b"}')`, []string{"Error calling function 'parseArray'", "String is not a JSON array"})
}

func TestArrayObject3(t *testing.T) {
	assertError(t, `parseArray('asd')`, []string{"Error calling function 'parseArray'", "Failed to parse string as a valid JSON", "Unexpected token 'a' at line 1, column 0 at #"})
}

func TestParseObject(t *testing.T) {
	eval := evaluate(t, `parseObject('{"a": "b"}')`)
	assertObject(t, eval, types.AsObject(map[string]interface{}{"a": "b"}))
}

func TestParseObject2(t *testing.T) {
	assertError(t, `parseObject('["a", "b", "c"]')`, []string{"Error calling function 'parseObject'", "String is not a JSON object"})
}

func TestParseObject3(t *testing.T) {
	assertError(t, `parseObject('asd')`, []string{"Error calling function 'parseObject'", "Failed to parse string as a valid JSON", "Unexpected token 'a' at line 1, column 0 at #"})
}
