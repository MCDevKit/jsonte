//go:build !quick

package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetLatestBPFile(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{}, true)
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
