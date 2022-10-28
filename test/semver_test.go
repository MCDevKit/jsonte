package test

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func TestSemverParsing(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0')`)
	assertSemver(t, eval, utils.Semver{Major: 1, Minor: 8, Patch: 0})
}

func TestSemverParsing1(t *testing.T) {
	eval := evaluate(t, `semver([1, 8, 0])`)
	assertSemver(t, eval, utils.Semver{Major: 1, Minor: 8, Patch: 0})
}

func TestSemverParsing2(t *testing.T) {
	eval := evaluate(t, `semver(1, 8, 0)`)
	assertSemver(t, eval, utils.Semver{Major: 1, Minor: 8, Patch: 0})
}

func TestSemverComparison(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') > semver('1.7.0')`)
	assertBool(t, eval, true)
}

func TestSemverComparison2(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') > semver('1.8.0')`)
	assertBool(t, eval, false)
}

func TestSemverComparison3(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') >= semver('1.8.0')`)
	assertBool(t, eval, true)
}

func TestSemverComparison4(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') <= semver('1.8.0')`)
	assertBool(t, eval, true)
}

func TestSemverComparison5(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') < semver('1.8.0')`)
	assertBool(t, eval, false)
}

func TestSemverComparison6(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') == semver('1.8.0')`)
	assertBool(t, eval, true)
}

func TestSemverComparison7(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0') != semver('1.8.0')`)
	assertBool(t, eval, false)
}

func TestSemverAccessors(t *testing.T) {
	eval := evaluate(t, `semver('1.8.0').major`)
	assertNumber(t, eval, 1)
}
