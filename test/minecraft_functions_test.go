//go:build !quick

package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"io"
	"strings"
	"testing"
)

func TestGetOldRPFile(t *testing.T) {
	safeio.Resolver = CacheFS
	expected := `{
  "format_version": "1.8.0",
  "minecraft:client_entity": {
    "description": {
      "identifier": "minecraft:wither",
      "min_engine_version": "1.8.0",
      "materials": {
        "default": "wither_boss",
        "armor": "wither_boss_armor"
      },
      "textures": {
        "default": "textures/entity/wither_boss/wither",
        "armor_white": "textures/entity/wither_boss/wither_armor_white",
        "armor_blue": "textures/entity/wither_boss/wither_armor_blue",
        "invulnerable": "textures/entity/wither_boss/wither_invulnerable"
      },
      "geometry": {
        "default": "geometry.witherBoss",
        "armor": "geometry.witherBoss.armor.v1.8"
      },
      "scripts": {
        "pre_animation": [
          "variable.base_scale = 2;",
          "variable.swell_clamped = Math.clamp(query.swell_amount, 0.0, 1.0);",
          "variable.wobble = 1.0 + Math.sin(query.swell_amount * 5730) * query.swell_amount * 0.01;",
          "variable.swell_adjustment = Math.pow(variable.swell_clamped, 4);",
          "variable.scale_xz = (1.0 + variable.swell_adjustment * 0.4) * variable.wobble;",
          "variable.scale_y = (1.0 + variable.swell_adjustment * 0.1) / variable.wobble;",
          "variable.body_base_rotation = Math.cos(query.life_time * 114.6);",
          "variable.upper_body_rotation = (0.065 + 0.05 * variable.body_base_rotation) * 180 + query.target_x_rotation;",
          "variable.is_invulnerable = query.invulnerable_ticks > 0.0;",
          "variable.display_normal_skin = (query.invulnerable_ticks <= 0) || ((query.invulnerable_ticks <= 80) && (Math.mod(query.invulnerable_ticks / 5, 2) == 1));"
        ],
        "scalex": "variable.scale_xz * variable.base_scale",
        "scaley": "variable.scale_y * variable.base_scale",
        "scalez": "variable.scale_xz * variable.base_scale"
      },
      "animations": {
        "scale": "animation.wither_boss.scale",
        "move": "animation.wither_boss.move",
        "look_at_target": "animation.wither_boss.look_at_target"
      },
      "animation_controllers": [
        { "move": "controller.animation.wither_boss.move" }
      ],
      "render_controllers": [ "controller.render.wither_boss", "controller.render.wither_boss_armor_white", "controller.render.wither_boss_armor_blue" ]
    }
  }
}`
	eval := evaluate(t, `getRPFile('entity/wither.entity.json', semver('1.19.50'))`)
	assertString(t, eval, "packs/1.19.50.2/RP/entity/wither.entity.json")
	open, err := safeio.Resolver.Open(eval.Value.StringValue())
	if err != nil {
		t.Fatal(err)
	}
	all, err := io.ReadAll(open)
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
