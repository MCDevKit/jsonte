package test

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
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

func TestRound(t *testing.T) {
	eval := evaluate(t, `round(1.5)`)
	assertNumber(t, eval, 2)
}

func TestRound2(t *testing.T) {
	eval := evaluate(t, `round(1.4)`)
	assertNumber(t, eval, 1)
}

func TestCeil(t *testing.T) {
	eval := evaluate(t, `ceil(1.4)`)
	assertNumber(t, eval, 2)
}

func TestCeil2(t *testing.T) {
	eval := evaluate(t, `ceil(1.5)`)
	assertNumber(t, eval, 2)
}

func TestFloor(t *testing.T) {
	eval := evaluate(t, `floor(1.5)`)
	assertNumber(t, eval, 1)
}

func TestFloor2(t *testing.T) {
	eval := evaluate(t, `floor(1.4)`)
	assertNumber(t, eval, 1)
}

func TestSqrt(t *testing.T) {
	eval := evaluate(t, `sqrt(4)`)
	assertNumber(t, eval, 2)
}

func TestSqrt2(t *testing.T) {
	eval := evaluate(t, `sqrt(9)`)
	assertNumber(t, eval, 3)
}

func TestAbs(t *testing.T) {
	eval := evaluate(t, `abs(-1)`)
	assertNumber(t, eval, 1)
}

func TestAbs2(t *testing.T) {
	eval := evaluate(t, `abs(1)`)
	assertNumber(t, eval, 1)
}

func TestPow(t *testing.T) {
	eval := evaluate(t, `pow(2, 3)`)
	assertNumber(t, eval, 8)
}

func TestPow2(t *testing.T) {
	eval := evaluate(t, `pow(3, 2)`)
	assertNumber(t, eval, 9)
}

func TestRoundPrecision(t *testing.T) {
	eval := evaluate(t, `round(1.23456, 3)`)
	assertNumber(t, eval, 1.235)
}

func TestRoundPrecision2(t *testing.T) {
	eval := evaluate(t, `round(1.23456, 2)`)
	assertNumber(t, eval, 1.23)
}

func TestMin(t *testing.T) {
	eval := evaluate(t, `min(1, 2)`)
	assertNumber(t, eval, 1)
}

func TestMin2(t *testing.T) {
	eval := evaluate(t, `min(2, 1)`)
	assertNumber(t, eval, 1)
}

func TestMax(t *testing.T) {
	eval := evaluate(t, `max(1, 2)`)
	assertNumber(t, eval, 2)
}

func TestMax2(t *testing.T) {
	eval := evaluate(t, `max(2, 1)`)
	assertNumber(t, eval, 2)
}

func TestClamp(t *testing.T) {
	eval := evaluate(t, `clamp(1, 0, 2)`)
	assertNumber(t, eval, 1)
}

func TestClamp2(t *testing.T) {
	eval := evaluate(t, `clamp(3, 0, 2)`)
	assertNumber(t, eval, 2)
}

func TestClamp3(t *testing.T) {
	eval := evaluate(t, `clamp(-1, 0, 2)`)
	assertNumber(t, eval, 0)
}

func TestPi(t *testing.T) {
	eval := evaluate(t, `pi()`)
	// Special case prepared for automatic decimal truncation
	assertNumber(t, eval, utils.ToNumber(math.Pi).FloatValue())
}

func TestBitwiseAnd(t *testing.T) {
	eval := evaluate(t, `bitwiseAnd(1, 1)`)
	assertNumber(t, eval, 1)
}

func TestBitwiseAnd2(t *testing.T) {
	eval := evaluate(t, `bitwiseAnd(1, 0)`)
	assertNumber(t, eval, 0)
}

func TestBitwiseOr(t *testing.T) {
	eval := evaluate(t, `bitwiseOr(1, 1)`)
	assertNumber(t, eval, 1)
}

func TestBitwiseOr2(t *testing.T) {
	eval := evaluate(t, `bitwiseOr(1, 0)`)
	assertNumber(t, eval, 1)
}

func TestBitwiseXor(t *testing.T) {
	eval := evaluate(t, `bitwiseXor(1, 1)`)
	assertNumber(t, eval, 0)
}

func TestBitwiseXor2(t *testing.T) {
	eval := evaluate(t, `bitwiseXor(1, 0)`)
	assertNumber(t, eval, 1)
}

func TestBitwiseNot(t *testing.T) {
	eval := evaluate(t, `bitwiseNot(1)`)
	assertNumber(t, eval, -2)
}

func TestBitwiseNot2(t *testing.T) {
	eval := evaluate(t, `bitwiseNot(0)`)
	assertNumber(t, eval, -1)
}

func TestBitshiftLeft(t *testing.T) {
	eval := evaluate(t, `bitshiftLeft(1, 1)`)
	assertNumber(t, eval, 2)
}

func TestBitshiftLeft2(t *testing.T) {
	eval := evaluate(t, `bitshiftLeft(1, 2)`)
	assertNumber(t, eval, 4)
}

func TestBitshiftRight(t *testing.T) {
	eval := evaluate(t, `bitshiftRight(4, 1)`)
	assertNumber(t, eval, 2)
}

func TestBitshiftRight2(t *testing.T) {
	eval := evaluate(t, `bitshiftRight(4, 2)`)
	assertNumber(t, eval, 1)
}

func TestNumberOfDigits(t *testing.T) {
	eval := evaluate(t, `numberOfDigits(123)`)
	assertNumber(t, eval, 3)
}

func TestNumberOfDigits2(t *testing.T) {
	eval := evaluate(t, `numberOfDigits(12345)`)
	assertNumber(t, eval, 5)
}
