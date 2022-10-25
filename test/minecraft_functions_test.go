//go:build !quick

package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
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
	open, err := safeio.Resolver.Open(eval.Value.(string))
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

func TestFindItemInfoByName(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:stone","legacyId":1,"metadata":5,"maxDurability":0,"damage":0,"armor":0,"maxStackSize":64,"tags":[],"category":"nature","nutrition":0,"fuelDuration":0,"aliases":[],"nameKey":"tile.stone.andesite.name","langName":"Andesite"}`
	eval := evaluate(t, `findItemInfoByName('andesite')`)
	object, err := utils.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObject(t, eval, object)
}

func TestFindItemInfoById(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:blue_glazed_terracotta","legacyId":231,"metadata":0,"maxDurability":0,"damage":0,"armor":0,"maxStackSize":64,"tags":[],"category":"construction","nutrition":0,"fuelDuration":0,"aliases":["minecraft:glazedTerracotta.blue"],"nameKey":"tile.glazedTerracotta.blue.name","langName":"Blue Glazed Terracotta"}`
	eval := evaluate(t, `findItemInfoById('blue_terracotta')`)
	object, err := utils.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObject(t, eval, object)
}

func TestGetItemInfo(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{"id":"minecraft:stone","legacyId":1,"metadata":0,"maxDurability":0,"damage":0,"armor":0,"maxStackSize":64,"tags":[],"category":"nature","nutrition":0,"fuelDuration":0,"aliases":[],"nameKey":"tile.stone.stone.name","langName":"Stone"}`
	eval := evaluate(t, `getItemInfo('stone', 0)`)
	object, err := utils.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	assertObject(t, eval, object)
}
