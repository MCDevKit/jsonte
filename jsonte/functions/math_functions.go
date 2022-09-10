package functions

import (
	"jsonte/jsonte/utils"
	"math"
)

func RegisterMathFunctions() {
	RegisterFunction(JsonFunction{
		Name: "floor",
		Body: floor,
	})
	RegisterFunction(JsonFunction{
		Name: "ceil",
		Body: ceil,
	})
	RegisterFunction(JsonFunction{
		Name: "round",
		Body: round,
	})
	RegisterFunction(JsonFunction{
		Name: "round",
		Body: roundPrecision,
	})
	RegisterFunction(JsonFunction{
		Name: "sin",
		Body: sin,
	})
	RegisterFunction(JsonFunction{
		Name: "cos",
		Body: cos,
	})
	RegisterFunction(JsonFunction{
		Name: "tan",
		Body: tan,
	})
	RegisterFunction(JsonFunction{
		Name: "asin",
		Body: asin,
	})
	RegisterFunction(JsonFunction{
		Name: "acos",
		Body: acos,
	})
	RegisterFunction(JsonFunction{
		Name: "atan",
		Body: atan,
	})
	RegisterFunction(JsonFunction{
		Name: "atan2",
		Body: atan2,
	})
	RegisterFunction(JsonFunction{
		Name: "sqrt",
		Body: sqrt,
	})
	RegisterFunction(JsonFunction{
		Name: "abs",
		Body: abs,
	})
	RegisterFunction(JsonFunction{
		Name: "min",
		Body: min,
	})
	RegisterFunction(JsonFunction{
		Name: "max",
		Body: max,
	})
	RegisterFunction(JsonFunction{
		Name: "clamp",
		Body: clamp,
	})
	RegisterFunction(JsonFunction{
		Name: "pi",
		Body: pi,
	})
	RegisterFunction(JsonFunction{
		Name: "mod",
		Body: mod,
	})
	RegisterFunction(JsonFunction{
		Name: "rotationToNormal",
		Body: rotationToNormal,
	})
	RegisterFunction(JsonFunction{
		Name: "bitwiseAnd",
		Body: bitwiseAnd,
	})
	RegisterFunction(JsonFunction{
		Name: "bitwiseOr",
		Body: bitwiseOr,
	})
	RegisterFunction(JsonFunction{
		Name: "bitwiseXor",
		Body: bitwiseXor,
	})
	RegisterFunction(JsonFunction{
		Name: "bitwiseNot",
		Body: bitwiseNot,
	})
	RegisterFunction(JsonFunction{
		Name: "bitshiftLeft",
		Body: bitshiftLeft,
	})
	RegisterFunction(JsonFunction{
		Name: "bitshiftRight",
		Body: bitshiftRight,
	})
	RegisterFunction(JsonFunction{
		Name: "numberOfDigits",
		Body: numberOfDigits,
	})
	RegisterFunction(JsonFunction{
		Name: "pow",
		Body: pow,
	})
}

func floor(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Floor(a.FloatValue()),
		Decimal: false,
	}
}

func ceil(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Ceil(a.FloatValue()),
		Decimal: false,
	}
}

func round(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Round(a.FloatValue()),
		Decimal: false,
	}
}

func roundPrecision(a utils.JsonNumber, precision utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Round(a.FloatValue()*math.Pow(10, float64(precision.IntValue()))) / math.Pow(10, float64(precision.IntValue())),
		Decimal: true,
	}
}

func toRadians(a float64) float64 {
	return a * math.Pi / 180.0
}

func fromRadians(a float64) float64 {
	return a * 180.0 / math.Pi
}

func sin(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Sin(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func cos(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Cos(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func tan(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Tan(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func asin(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   fromRadians(math.Asin(a.FloatValue())),
		Decimal: true,
	}
}

func acos(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   fromRadians(math.Acos(a.FloatValue())),
		Decimal: true,
	}
}

func atan(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   fromRadians(math.Atan(a.FloatValue())),
		Decimal: true,
	}
}

func atan2(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   fromRadians(math.Atan2(a.FloatValue(), b.FloatValue())),
		Decimal: true,
	}
}

func sqrt(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Sqrt(a.FloatValue()),
		Decimal: a.Decimal,
	}
}

func abs(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Abs(a.FloatValue()),
		Decimal: a.Decimal,
	}
}

func min(a, b utils.JsonNumber) utils.JsonNumber {
	if a.FloatValue() < b.FloatValue() {
		return a
	}
	return b
}

func max(a, b utils.JsonNumber) utils.JsonNumber {
	if a.FloatValue() > b.FloatValue() {
		return a
	}
	return b
}

func clamp(a, min, max utils.JsonNumber) utils.JsonNumber {
	if a.FloatValue() < min.FloatValue() {
		return min
	}
	if a.FloatValue() > max.FloatValue() {
		return max
	}
	return a
}

func mod(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Mod(a.FloatValue(), b.FloatValue()),
		Decimal: a.Decimal || b.Decimal,
	}
}

func pi() utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Pi,
		Decimal: true,
	}
}

func rotationToNormal(xRot, yRot utils.JsonNumber) []utils.JsonNumber {
	x := roundPrecision(utils.ToNumber(cos(xRot).FloatValue()*sin(yRot).FloatValue()), utils.ToNumber(5))
	y := roundPrecision(utils.ToNumber(-sin(xRot).FloatValue()), utils.ToNumber(5))
	z := roundPrecision(utils.ToNumber(cos(yRot).FloatValue()*cos(xRot).FloatValue()), utils.ToNumber(5))
	return []utils.JsonNumber{x, y, z}
}

func bitwiseAnd(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(a.IntValue() & b.IntValue()),
		Decimal: false,
	}
}

func bitwiseOr(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(a.IntValue() | b.IntValue()),
		Decimal: false,
	}
}

func bitwiseXor(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(a.IntValue() ^ b.IntValue()),
		Decimal: false,
	}
}

func bitwiseNot(a utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(^a.IntValue()),
		Decimal: false,
	}
}

func bitshiftLeft(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(a.IntValue() << b.IntValue()),
		Decimal: false,
	}
}

func bitshiftRight(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   float64(a.IntValue() >> b.IntValue()),
		Decimal: false,
	}
}

func numberOfDigits(a utils.JsonNumber) utils.JsonNumber {
	if a.IntValue() == 0 {
		return utils.ToNumber(1)
	}
	return utils.JsonNumber{
		Value:   math.Log10(a.FloatValue()) + 1,
		Decimal: false,
	}
}

func pow(a, b utils.JsonNumber) utils.JsonNumber {
	return utils.JsonNumber{
		Value:   math.Pow(a.FloatValue(), b.FloatValue()),
		Decimal: a.Decimal || b.Decimal,
	}
}
