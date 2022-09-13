package test

import (
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
	assertArray(t, eval, []interface{}{"Hello", "World"})
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
	eval := evaluate(t, `'hello'.capitalize()`)
	assertString(t, eval, "Hello")
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
	assertArray(t, eval, []interface{}{"h", "e", "l", "l", "o"})
}

func TestLength(t *testing.T) {
	eval := evaluate(t, `'hello'.length()`)
	assertNumber(t, eval, 5)
}
