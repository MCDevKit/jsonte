package test

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func TestJsonPathParsing(t *testing.T) {
	eval := evaluate(t, `jsonpath('#/test/path[1]')`)
	assertJsonPath(t, eval, types.JsonPath{
		Path: []types.JsonType{
			types.JsonString{Value: "test"},
			types.JsonString{Value: "path"},
			types.AsNumber(1),
		},
	})
}

func TestJsonPathParsingWithoutHash(t *testing.T) {
	eval := evaluate(t, `jsonpath('/test/path[1]')`)
	assertJsonPath(t, eval, types.JsonPath{
		Path: []types.JsonType{
			types.JsonString{Value: "test"},
			types.JsonString{Value: "path"},
			types.AsNumber(1),
		},
	})
}

func TestJsonPathParent(t *testing.T) {
	eval := evaluate(t, `jsonpath('test/path[1]').parent()`)
	assertJsonPath(t, eval, types.JsonPath{
		Path: []types.JsonType{
			types.JsonString{Value: "test"},
			types.JsonString{Value: "path"},
		},
	})
}

func TestJsonPathParsingNumber(t *testing.T) {
	eval := evaluate(t, `jsonpath('#[1]')`)
	assertJsonPath(t, eval, types.JsonPath{
		Path: []types.JsonType{
			types.AsNumber(1),
		},
	})
}

func TestObjectAccessByJsonPath(t *testing.T) {
	eval := evaluate(t, `{"test": {"path": [1, 2, 3]}}[jsonpath('#/test/path[1]')]`)
	assertNumber(t, eval, 2)
}

func TestArrayAccessByJsonPath(t *testing.T) {
	eval := evaluate(t, `[1, 2, 3][jsonpath('#[1]')]`)
	assertNumber(t, eval, 2)
}
