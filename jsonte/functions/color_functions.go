package functions

import (
	"errors"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"image/color"
)

func RegisterColorFunctions() {
	RegisterFunction(JsonFunction{
		Name: "hexToArray",
		Body: hexToArray,
	})
}

func hexToArray(str string) (utils.JsonArray, error) {
	fast, err := parseHexColorFast(str)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to parse hex color %s", str)
	}
	return utils.JsonArray{
		utils.ToNumber(float64(fast.R) / 255),
		utils.ToNumber(float64(fast.G) / 255),
		utils.ToNumber(float64(fast.B) / 255),
	}, nil
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
