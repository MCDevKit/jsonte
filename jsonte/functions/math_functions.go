package functions

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
	"math/rand"
	"strconv"
)

func RegisterMathFunctions() {
	const group = "math"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Math functions",
		Summary: "Math functions allow for executing more complicated arithmetic.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "floor",
		Body:  floor,
		Docs: Docs{
			Summary: "Returns the largest integer less than or equal to the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to floor.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{floor(3.6)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "ceil",
		Body:  ceil,
		Docs: Docs{
			Summary: "Returns the smallest integer greater than or equal to the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to ceil.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 4",
    "test": "{{ceil(3.6)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "round",
		Body:  round,
		Docs: Docs{
			Summary: "Returns the nearest integer to the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to round.",
				},
				{
					Name:     "precision",
					Summary:  "The number of decimal places to round to.",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3.142",
    "test": "{{round(3.1415, 3)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "round",
		Body:  roundPrecision,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "sin",
		Body:  sin,
		Docs: Docs{
			Summary: "Returns the sine of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the sine of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 0.7071067811865475",
    "test": "{{sin(45)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "cos",
		Body:  cos,
		Docs: Docs{
			Summary: "Returns the cosine of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the cosine of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 0.7071067811865476",
    "test": "{{cos(45)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "tan",
		Body:  tan,
		Docs: Docs{
			Summary: "Returns the tangent of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the tangent of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{tan(45)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "asin",
		Body:  asin,
		Docs: Docs{
			Summary: "Returns the arcsine of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the arcsine of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 44.991348337162016",
    "test": "{{asin(0.707)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "acos",
		Body:  acos,
		Docs: Docs{
			Summary: "Returns the arccosine of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the arccosine of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 45.008651662837984",
    "test": "{{acos(0.707)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "atan",
		Body:  atan,
		Docs: Docs{
			Summary: "Returns the arctangent of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the arctangent of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 45",
    "test": "{{atan(1)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "atan2",
		Body:  atan2,
		Docs: Docs{
			Summary: "Returns the arctangent of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "y",
					Summary: "The first number to get the arctangent of.",
				},
				{
					Name:    "x",
					Summary: "The second number to get the arctangent of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 45",
    "test": "{{atan2(1, 1)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "sqrt",
		Body:  sqrt,
		Docs: Docs{
			Summary: "Returns the square root of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the square root of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{sqrt(9)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "abs",
		Body:  abs,
		Docs: Docs{
			Summary: "Returns the absolute value of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to get the absolute value of.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{abs(-3)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "min",
		Body:  min,
		Docs: Docs{
			Summary: "Returns the smallest of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number1",
					Summary: "The first number to compare.",
				},
				{
					Name:    "number2",
					Summary: "The second number to compare.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{min(3, 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "max",
		Body:  max,
		Docs: Docs{
			Summary: "Returns the largest of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number1",
					Summary: "The first number to compare.",
				},
				{
					Name:    "number2",
					Summary: "The second number to compare.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 5",
    "test": "{{max(3, 5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "clamp",
		Body:  clamp,
		Docs: Docs{
			Summary: "Returns the given number clamped between the given minimum and maximum.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to clamp.",
				},
				{
					Name:    "min",
					Summary: "The minimum value.",
				},
				{
					Name:    "max",
					Summary: "The maximum value.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 5",
    "test": "{{clamp(3, 5, 10)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "pi",
		Body:  pi,
		Docs: Docs{
			Summary: "Returns the value of pi.",
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3.141593",
    "test": "{{pi()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "mod",
		Body:  mod,
		Docs: Docs{
			Summary: "Returns the remainder of the first number divided by the second number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to divide.",
				},
				{
					Name:    "divisor",
					Summary: "The number to divide by.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{mod(5, 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "rotationToNormal",
		Body:  rotationToNormal,
		Docs: Docs{
			Summary: "Returns normal (direction) vector based on pitch and yaw rotation.",
			Arguments: []Argument{
				{
					Name:    "pitch",
					Summary: "A pitch rotation (rotation in x-axis)",
				},
				{
					Name:    "yaw",
					Summary: "A yaw rotation (rotation in y-axis)",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [0,-0,1]",
    "test": "{{rotationToNormal(0, 0)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitwiseAnd",
		Body:  bitwiseAnd,
		Docs: Docs{
			Summary: "Returns the bitwise AND of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number1",
					Summary: "The first number to bitwise AND.",
				},
				{
					Name:    "number2",
					Summary: "The second number to bitwise AND.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 2",
    "test": "{{bitwiseAnd(3, 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitwiseOr",
		Body:  bitwiseOr,
		Docs: Docs{
			Summary: "Returns the bitwise OR of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number1",
					Summary: "The first number to bitwise OR.",
				},
				{
					Name:    "number2",
					Summary: "The second number to bitwise OR.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{bitwiseOr(3, 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitwiseXor",
		Body:  bitwiseXor,
		Docs: Docs{
			Summary: "Returns the bitwise XOR of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number1",
					Summary: "The first number to bitwise XOR.",
				},
				{
					Name:    "number2",
					Summary: "The second number to bitwise XOR.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{bitwiseXor(3, 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitwiseNot",
		Body:  bitwiseNot,
		Docs: Docs{
			Summary: "Returns the bitwise NOT of the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to bitwise NOT.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be -4",
    "test": "{{bitwiseNot(3)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitshiftLeft",
		Body:  bitshiftLeft,
		Docs: Docs{
			Summary: "Returns the bitwise left shift of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to bitwise left shift.",
				},
				{
					Name:    "shift",
					Summary: "The number of bits to shift.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 12",
    "test": "{{bitshiftLeft(3, 2)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "bitshiftRight",
		Body:  bitshiftRight,
		Docs: Docs{
			Summary: "Returns the bitwise right shift of the given numbers.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to bitwise right shift.",
				},
				{
					Name:    "shift",
					Summary: "The number of bits to shift.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{bitshiftRight(3, 1)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "numberOfDigits",
		Body:  numberOfDigits,
		Docs: Docs{
			Summary: "Returns the number of digits in the given number.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to count the number of digits.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 3",
    "test": "{{numberOfDigits(123)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "pow",
		Body:  pow,
		Docs: Docs{
			Summary: "Returns the given number raised to the given power.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to raise to the given power.",
				},
				{
					Name:    "power",
					Summary: "The power to raise the given number to.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 27",
    "test": "{{pow(3, 3)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "random",
		Body:  randomNumber,
		Docs: Docs{
			Summary: "Returns a random number between 0 and 1.",
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be a random number between 0 and 1",
    "test": "{{random()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "randomInt",
		Body:  randomInt,
		Docs: Docs{
			Summary: "Returns a random integer number between the first and second argument.",
			Arguments: []Argument{
				{
					Name:    "min",
					Summary: "The minimum value of the random number.",
				},
				{
					Name:    "max",
					Summary: "The maximum value of the random number.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be a random number between 0 and 1",
    "test": "{{random()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "wrap",
		Body:  wrapFull,
		Docs: Docs{
			Summary: "Returns the given number wrapped between the first and second argument.",
			Arguments: []Argument{
				{
					Name:    "value",
					Summary: "The value to wrap.",
				},
				{
					Name:     "start",
					Summary:  "The start of the range to wrap the value between.",
					Optional: true,
				},
				{
					Name:    "end",
					Summary: "The end of the range to wrap the value between.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{wrap(10, 0, 4)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "wrap",
		Body:  wrapShort,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "toHex",
		Body:  toHexFixed,
		Docs: Docs{
			Summary: "Returns the given number as a hexadecimal string.",
			Arguments: []Argument{
				{
					Name:    "number",
					Summary: "The number to convert to a hexadecimal string.",
				},
				{
					Name:     "digits",
					Summary:  "The number of digits to pad the hexadecimal string with.",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 007b",
    "test": "{{toHex(123, 4)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "toHex",
		Body:  toHex,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "fromHex",
		Body:  fromHex,
		Docs: Docs{
			Summary: "Returns the given hexadecimal string as a number.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The hexadecimal string to convert to a number.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 123",
    "test": "{{fromHex('7b')}}"
  }
}
</code>`,
		},
	})
}

func wrapFull(value, start, end *types.JsonNumber) *types.JsonNumber {
	return types.AsNumber(utils.WrapRange(int(value.IntValue()), int(start.IntValue()), int(end.IntValue())))
}

func wrapShort(value, end *types.JsonNumber) *types.JsonNumber {
	return types.AsNumber(utils.WrapRange(int(value.IntValue()), 0, int(end.IntValue())))
}

func randomNumber() *types.JsonNumber {
	return types.AsNumber(rand.Float64())
}

func randomInt(a, b *types.JsonNumber) (*types.JsonNumber, error) {
	if a.IntValue() > b.IntValue() {
		a, b = b, a
	}
	return types.AsNumber(rand.Intn(int(b.IntValue()-a.IntValue())) + int(a.IntValue())), nil
}

func floor(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Floor(a.FloatValue()),
		Decimal: false,
	}
}

func ceil(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Ceil(a.FloatValue()),
		Decimal: false,
	}
}

func round(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Round(a.FloatValue()),
		Decimal: false,
	}
}

func roundPrecision(a *types.JsonNumber, precision *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
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

func sin(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Sin(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func cos(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Cos(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func tan(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Tan(toRadians(a.FloatValue())),
		Decimal: true,
	}
}

func asin(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   fromRadians(math.Asin(a.FloatValue())),
		Decimal: true,
	}
}

func acos(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   fromRadians(math.Acos(a.FloatValue())),
		Decimal: true,
	}
}

func atan(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   fromRadians(math.Atan(a.FloatValue())),
		Decimal: true,
	}
}

func atan2(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   fromRadians(math.Atan2(a.FloatValue(), b.FloatValue())),
		Decimal: true,
	}
}

func sqrt(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Sqrt(a.FloatValue()),
		Decimal: a.Decimal,
	}
}

func abs(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Abs(a.FloatValue()),
		Decimal: a.Decimal,
	}
}

func min(a, b *types.JsonNumber) *types.JsonNumber {
	if a.FloatValue() < b.FloatValue() {
		return a
	}
	return b
}

func max(a, b *types.JsonNumber) *types.JsonNumber {
	if a.FloatValue() > b.FloatValue() {
		return a
	}
	return b
}

func clamp(a, min, max *types.JsonNumber) *types.JsonNumber {
	if a.FloatValue() < min.FloatValue() {
		return min
	}
	if a.FloatValue() > max.FloatValue() {
		return max
	}
	return a
}

func mod(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Mod(a.FloatValue(), b.FloatValue()),
		Decimal: a.Decimal || b.Decimal,
	}
}

func pi() *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Pi,
		Decimal: true,
	}
}

func rotationToNormal(xRot, yRot *types.JsonNumber) *types.JsonArray {
	x := roundPrecision(types.AsNumber(cos(xRot).FloatValue()*sin(yRot).FloatValue()), types.AsNumber(5))
	y := roundPrecision(types.AsNumber(-sin(xRot).FloatValue()), types.AsNumber(5))
	z := roundPrecision(types.AsNumber(cos(yRot).FloatValue()*cos(xRot).FloatValue()), types.AsNumber(5))
	return &types.JsonArray{Value: []types.JsonType{x, y, z}}
}

func bitwiseAnd(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(a.IntValue() & b.IntValue()),
		Decimal: false,
	}
}

func bitwiseOr(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(a.IntValue() | b.IntValue()),
		Decimal: false,
	}
}

func bitwiseXor(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(a.IntValue() ^ b.IntValue()),
		Decimal: false,
	}
}

func bitwiseNot(a *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(^a.IntValue()),
		Decimal: false,
	}
}

func bitshiftLeft(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(a.IntValue() << b.IntValue()),
		Decimal: false,
	}
}

func bitshiftRight(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   float64(a.IntValue() >> b.IntValue()),
		Decimal: false,
	}
}

func numberOfDigits(a *types.JsonNumber) *types.JsonNumber {
	if a.IntValue() == 0 {
		return types.AsNumber(1)
	}
	return &types.JsonNumber{
		Value:   math.Log10(a.FloatValue()) + 1,
		Decimal: false,
	}
}

func pow(a, b *types.JsonNumber) *types.JsonNumber {
	return &types.JsonNumber{
		Value:   math.Pow(a.FloatValue(), b.FloatValue()),
		Decimal: a.Decimal || b.Decimal,
	}
}

func toHex(a *types.JsonNumber) *types.JsonString {
	return &types.JsonString{
		Value: fmt.Sprintf("%x", int64(a.IntValue())),
	}
}

func toHexFixed(a *types.JsonNumber, digits *types.JsonNumber) *types.JsonString {
	return &types.JsonString{
		Value: fmt.Sprintf("%0"+strconv.Itoa(int(digits.IntValue()))+"x", int64(a.IntValue())),
	}
}

func fromHex(a *types.JsonString) (*types.JsonNumber, error) {
	i, err := strconv.ParseInt(a.Value, 16, 64)
	if err != nil {
		return types.AsNumber(0), burrito.WrapErrorf(err, "Failed to parse hex string %s", a.Value)
	}
	return types.AsNumber(i), nil
}
