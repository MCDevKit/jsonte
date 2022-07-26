package functions

import (
	"jsonte/jsonte/utils"
	"math"
)

func RegisterMathFunctions() {
	RegisterFunction(JsonFunction{
		Name: "pow",
		Body: pow,
	})
}

func pow(a, b utils.JsonNumber) (utils.JsonNumber, error) {
	return utils.JsonNumber{
		Value:   math.Pow(a.FloatValue(), b.FloatValue()),
		Decimal: a.Decimal || b.Decimal,
	}, nil
}
