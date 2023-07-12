package functions

import (
	"errors"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"image/color"
	"math"
)

func RegisterColorFunctions() {
	const group = "color"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Color functions",
		Summary: "Color functions are related to converting and manipulating colors.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "hexToArray",
		Body:  hexToArray,
		Docs: Docs{
			Summary: "Converts a hex color to an array of RGBA values.",
			Arguments: []Argument{
				{
					Name:    "hex",
					Summary: "The hex color to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [0.2, 0.4, 0.6, 1]",
    "test": "{{hexToArray('#336699')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "arrayToHex",
		Body:  arrayToHex,
		Docs: Docs{
			Summary: "Converts an array of RGB(A) values to a hex color. If the array has transparency, the hex color will be in the format #AARRGGBB.",
			Arguments: []Argument{
				{
					Name:    "array",
					Summary: "The RGB(A) array to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '#336699'",
    "test": "{{arrayToHex([0.2, 0.4, 0.6])}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "hsl",
		Body:  hsl,
		Docs: Docs{
			Summary: "Creates a new color from hue, saturation and lightness values.",
			Arguments: []Argument{
				{
					Name:    "hsl",
					Summary: "The array of hue, saturation and lightness values.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [0, 255, 0]",
    "test": "{{hsl([120, 1, 0.5])}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "greyscale",
		Body:  greyscale,
		Docs: Docs{
			Summary: "Sets a gray color with the same lightness as the input color.",
			Arguments: []Argument{
				{
					Name:    "color",
					Summary: "The color to convert to greyscale as RGB array.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [0.6, 0.6, 0.6]",
    "test": "{{greyscale([0.3, 0.6, 0.9])}}"
  }
}
</code>`,
		},
	})
}

func hexToArray(str types.JsonString) (types.JsonArray, error) {
	fast, err := parseHexColorFast(str.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to parse hex color %s", str.StringValue())
	}
	return types.JsonArray{Value: []types.JsonType{
		types.AsNumber(float64(fast.R) / 255),
		types.AsNumber(float64(fast.G) / 255),
		types.AsNumber(float64(fast.B) / 255),
		types.AsNumber(float64(fast.A) / 255),
	}}, nil
}

func arrayToHex(array types.JsonArray) (types.JsonString, error) {
	if len(array.Value) < 3 || len(array.Value) > 4 {
		return types.EmptyString, burrito.WrappedErrorf("Expected array of length 3 or 4, got %d", len(array.Value))
	}
	var r, g, b byte
	if !types.IsNumber(array.Value[0]) {
		return types.EmptyString, burrito.WrappedErrorf("Expected number, got %s", array.Value[0].StringValue())
	}
	if !types.IsNumber(array.Value[1]) {
		return types.EmptyString, burrito.WrappedErrorf("Expected number, got %s", array.Value[1].StringValue())
	}
	if !types.IsNumber(array.Value[2]) {
		return types.EmptyString, burrito.WrappedErrorf("Expected number, got %s", array.Value[2].StringValue())
	}
	r = byte(types.AsNumber(array.Value[0]).FloatValue() * 255)
	g = byte(types.AsNumber(array.Value[1]).FloatValue() * 255)
	b = byte(types.AsNumber(array.Value[2]).FloatValue() * 255)
	if len(array.Value) == 4 {
		if !types.IsNumber(array.Value[3]) {
			return types.EmptyString, burrito.WrappedErrorf("Expected number, got %s", array.Value[3].StringValue())
		}
		a := byte(types.AsNumber(array.Value[3]).FloatValue() * 255)
		return types.AsString(fmt.Sprintf("#%02x%02x%02x%02x", a, r, g, b)), nil
	}
	return types.AsString(fmt.Sprintf("#%02x%02x%02x", r, g, b)), nil
}

func hsl(array types.JsonArray) (types.JsonArray, error) {
	if len(array.Value) != 3 {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected array of length 3, got %d", len(array.Value))
	}
	if !types.IsNumber(array.Value[0]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[0].StringValue())
	}
	if !types.IsNumber(array.Value[1]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[1].StringValue())
	}
	if !types.IsNumber(array.Value[2]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[2].StringValue())
	}
	h := types.AsNumber(array.Value[0]).FloatValue()
	s := types.AsNumber(array.Value[1]).FloatValue()
	l := types.AsNumber(array.Value[2]).FloatValue()
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2
	var r, g, b float64
	switch {
	case h < 60:
		r = c
		g = x
		b = 0
	case h < 120:
		r = x
		g = c
		b = 0
	case h < 180:
		r = 0
		g = c
		b = x
	case h < 240:
		r = 0
		g = x
		b = c
	case h < 300:
		r = x
		g = 0
		b = c
	case h < 360:
		r = c
		g = 0
		b = x
	}
	result := types.NewJsonArray()
	result.Value = make([]types.JsonType, 3)
	result.Value[0] = types.AsNumber((r + m) * 255)
	result.Value[1] = types.AsNumber((g + m) * 255)
	result.Value[2] = types.AsNumber((b + m) * 255)
	return result, nil
}

func greyscale(array types.JsonArray) (types.JsonArray, error) {
	if len(array.Value) != 3 {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected array of length 3, got %d", len(array.Value))
	}
	if !types.IsNumber(array.Value[0]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[0].StringValue())
	}
	if !types.IsNumber(array.Value[1]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[1].StringValue())
	}
	if !types.IsNumber(array.Value[2]) {
		return types.NewJsonArray(), burrito.WrappedErrorf("Expected number, got %s", array.Value[2].StringValue())
	}
	r := types.AsNumber(array.Value[0]).FloatValue()
	g := types.AsNumber(array.Value[1]).FloatValue()
	b := types.AsNumber(array.Value[2]).FloatValue()
	result := types.NewJsonArray()
	avg := (r + g + b) / 3
	result.Value = make([]types.JsonType, 3)
	result.Value[0] = types.AsNumber(avg)
	result.Value[1] = types.AsNumber(avg)
	result.Value[2] = types.AsNumber(avg)
	return result, nil
}

// from https://stackoverflow.com/a/54200713
var errInvalidFormat = errors.New("invalid format")

func parseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b rune) byte {
		switch {
		case b >= '0' && b <= '9':
			return byte(b - '0')
		case b >= 'a' && b <= 'f':
			return byte(b - 'a' + 10)
		case b >= 'A' && b <= 'F':
			return byte(b - 'A' + 10)
		}
		err = errInvalidFormat
		return 0
	}

	runes := []rune(s)
	switch len(runes) {
	case 7:
		c.R = hexToByte(runes[1])<<4 + hexToByte(runes[2])
		c.G = hexToByte(runes[3])<<4 + hexToByte(runes[4])
		c.B = hexToByte(runes[5])<<4 + hexToByte(runes[6])
	case 4:
		c.R = hexToByte(runes[1]) * 17
		c.G = hexToByte(runes[2]) * 17
		c.B = hexToByte(runes[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
