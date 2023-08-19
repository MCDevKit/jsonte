package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func TestHexToArray(t *testing.T) {
	eval := evaluate(t, `hexToArray('#ff0000')`)
	assertArray(t, eval, types.Box([]interface{}{types.AsNumber(1), types.AsNumber(0), types.AsNumber(0), types.AsNumber(1)}).(*types.JsonArray))
}
