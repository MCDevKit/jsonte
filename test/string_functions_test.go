package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func TestReplace(t *testing.T) {
	eval := evaluate(t, `'Hello'.replace('l', 'L')`)
	assertString(t, eval, "HeLLo")
}

func TestJoin(t *testing.T) {
	eval := evaluate(t, `['Hello', 'World'].join(' ')`)
	assertString(t, eval, "Hello World")
}

func TestStringContains(t *testing.T) {
	eval := evaluate(t, `'Hello'.contains('l')`)
	assertBool(t, eval, true)
}

func TestStringSplit(t *testing.T) {
	eval := evaluate(t, `'Hello World'.split(' ')`)
	assertArray(t, eval, types.Box([]interface{}{"Hello", "World"}).(types.JsonArray))
}

func TestStringIndexOf(t *testing.T) {
	eval := evaluate(t, `'Hello'.indexOf('l')`)
	assertNumber(t, eval, 2)
}

func TestStringLastIndexOf(t *testing.T) {
	eval := evaluate(t, `'Hello'.lastIndexOf('l')`)
	assertNumber(t, eval, 3)
}

func TestHash(t *testing.T) {
	eval := evaluate(t, `'Hello'.hash()`)
	assertNumber(t, eval, 4116459851)
}

func TestStringSubstring(t *testing.T) {
	eval := evaluate(t, `'Hello'.substring(1, 3)`)
	assertString(t, eval, "el")
}

func TestStringSubstring2(t *testing.T) {
	eval := evaluate(t, `'Hello'.substring(1)`)
	assertString(t, eval, "ello")
}

func TestStringSubstring3(t *testing.T) {
	eval := evaluate(t, `'Hello'.substring(1, 1)`)
	assertString(t, eval, "")
}

func TestToLowerCase(t *testing.T) {
	eval := evaluate(t, `'Hello'.toLowerCase()`)
	assertString(t, eval, "hello")
}

func TestToUpperCase(t *testing.T) {
	eval := evaluate(t, `'Hello'.toUpperCase()`)
	assertString(t, eval, "HELLO")
}

func TestCapitalize(t *testing.T) {
	eval := evaluate(t, `'hello world'.capitalize()`)
	assertString(t, eval, "Hello world")
}

func TestTitle(t *testing.T) {
	eval := evaluate(t, `'hello world'.title()`)
	assertString(t, eval, "Hello World")
}

func TestSwapCase(t *testing.T) {
	eval := evaluate(t, `'Hello'.swapCase()`)
	assertString(t, eval, "hELLO")
}

func TestStartsWith(t *testing.T) {
	eval := evaluate(t, `'hello'.startsWith('he')`)
	assertBool(t, eval, true)
}

func TestEndsWith(t *testing.T) {
	eval := evaluate(t, `'hello'.endsWith('lo')`)
	assertBool(t, eval, true)
}

func TestTrim(t *testing.T) {
	eval := evaluate(t, `' hello '.trim()`)
	assertString(t, eval, "hello")
}

func TestRegexReplace(t *testing.T) {
	eval := evaluate(t, `'hello'.regexReplace('[eo]', 'L')`)
	assertString(t, eval, "hLllL")
}

func TestChars(t *testing.T) {
	eval := evaluate(t, `'hello'.chars()`)
	assertArray(t, eval, types.Box([]interface{}{"h", "e", "l", "l", "o"}).(types.JsonArray))
}

func TestLength(t *testing.T) {
	eval := evaluate(t, `'hello'.length()`)
	assertNumber(t, eval, 5)
}

func TestNumber(t *testing.T) {
	eval := evaluate(t, `'1.5'.number()`)
	assertNumber(t, eval, 1.5)
}

func TestWrongNumber(t *testing.T) {
	expression := `'hello'.number()`
	expectedError := []string{
		"Error calling function 'number' on hello",
		"String 'hello' is not a valid number",
	}
	assertError(t, expression, expectedError)
}

func TestFormatString(t *testing.T) {
	expression := `'%s %s'.format('Hello', 'World')`
	assertString(t, evaluate(t, expression), "Hello World")
}

func TestRegexMatch(t *testing.T) {
	expression := `'hello world'.regexMatch('^hello ').count() == 1`
	assertBool(t, evaluate(t, expression), true)
}

func TestRegexMatch2(t *testing.T) {
	expression := `'not hello world'.regexMatch('^hello ').count() == 0`
	assertBool(t, evaluate(t, expression), true)
}
