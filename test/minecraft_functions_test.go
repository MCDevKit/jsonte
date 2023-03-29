//go:build !quick

package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetLatestBPFile(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{
  "format_version": "1.10",
  "minecraft:item": {
    "description": {
      "identifier": "minecraft:apple"
    },

    "components": {
      "minecraft:use_duration": 32,
      "minecraft:food": {
        "nutrition": 4,
        "saturation_modifier": "low"
      }
    }
  }
}`
	eval := evaluate(t, `getLatestBPFile('items/apple.json')`)
	assertString(t, eval, "packs/BP/items/apple.json")
	open, err := safeio.Resolver.Open(eval.Value.StringValue())
	if err != nil {
		t.Fatal(err)
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ReplaceAll(string(all), "\r\n", "\n") != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, string(all))
	}
	err = open.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindFileAndFail(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := []string{
		"Error calling function 'getLatestBPFile'",
		"File 'items/please_dont_make_such_item.json' does not exist",
		"file does not exist",
	}
	assertError(t, "getLatestBPFile('items/please_dont_make_such_item.json')", expected)
}

func TestFindFileWithoutFail(t *testing.T) {
	safeio.Resolver = CacheFS
	eval := evaluate(t, "getLatestBPFile('items/please_dont_make_such_item.json', false)")
	assertNull(t, eval)
}

func TestFindItemInfoByName(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:stone","legacyId":1,"metadata":5}`
	eval := evaluate(t, `findItemInfoByName('andesite')`)
	object, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObjectContains(t, eval, object)
}

func TestFindItemInfoById(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:blue_glazed_terracotta","legacyId":231,"metadata":0}`
	eval := evaluate(t, `findItemInfoById('blue_terracotta')`)
	object, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObjectContains(t, eval, object)
}

func TestGetItemInfo(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:stone","legacyId":1,"metadata":0}`
	eval := evaluate(t, `getItemInfo('stone', 0)`)
	object, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObjectContains(t, eval, object)
}
