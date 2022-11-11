package functions

import (
	"errors"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"image/color"
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
			Summary: "Converts a hex color to an array of RGB values.",
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
}

func hexToArray(str types.JsonString) (types.JsonArray, error) {
	fast, err := parseHexColorFast(str.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to parse hex color %s", str)
	}
	return types.JsonArray{Value: []types.JsonType{
		types.AsNumber(float64(fast.R) / 255),
		types.AsNumber(float64(fast.G) / 255),
		types.AsNumber(float64(fast.B) / 255),
		types.AsNumber(float64(fast.A) / 255),
	}}, nil
}

// from https://stackoverflow.com/a/54200713
var errInvalidFormat = errors.New("invalid format")

func parseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
