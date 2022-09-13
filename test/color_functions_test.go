package test

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func TestHexToArray(t *testing.T) {
	eval := evaluate(t, `hexToArray('#ff0000')`)
	assertArray(t, eval, utils.JsonArray{utils.ToNumber(1), utils.ToNumber(0), utils.ToNumber(0)})
}
