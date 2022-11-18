package test

import (
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
