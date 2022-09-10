package test

import (
	"testing"
)

func TestMod(t *testing.T) {
	eval := evaluate(t, `mod(5, 2)`)
	assertNumber(t, eval, 1)
}

func TestMod2(t *testing.T) {
	eval := evaluate(t, `mod(5, 3)`)
	assertNumber(t, eval, 2)
}

func TestSin(t *testing.T) {
	eval := evaluate(t, `sin(0)`)
	assertNumber(t, eval, 0)
}

func TestSin2(t *testing.T) {
	eval := evaluate(t, `sin(90)`)
	assertNumber(t, eval, 1)
}

func TestCos(t *testing.T) {
	eval := evaluate(t, `cos(0)`)
	assertNumber(t, eval, 1)
}

func TestCos2(t *testing.T) {
	eval := evaluate(t, `cos(90)`)
	assertNumber(t, eval, 0)
}

func TestTan(t *testing.T) {
	eval := evaluate(t, `tan(0)`)
	assertNumber(t, eval, 0)
}

func TestTan2(t *testing.T) {
	eval := evaluate(t, `tan(45)`)
	assertNumber(t, eval, 1)
}

func TestAsin(t *testing.T) {
	eval := evaluate(t, `asin(0)`)
	assertNumber(t, eval, 0)
}

func TestAsin2(t *testing.T) {
	eval := evaluate(t, `asin(1)`)
	assertNumber(t, eval, 90)
}

func TestAcos(t *testing.T) {
	eval := evaluate(t, `acos(0)`)
	assertNumber(t, eval, 90)
}

func TestAcos2(t *testing.T) {
	eval := evaluate(t, `acos(1)`)
	assertNumber(t, eval, 0)
}

func TestAtan(t *testing.T) {
	eval := evaluate(t, `atan(0)`)
	assertNumber(t, eval, 0)
}

func TestAtan2(t *testing.T) {
	eval := evaluate(t, `atan(1)`)
	assertNumber(t, eval, 45)
}

func TestAtan3(t *testing.T) {
	eval := evaluate(t, `atan2(1, 1)`)
	assertNumber(t, eval, 45)
}
