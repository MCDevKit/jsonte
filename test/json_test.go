package test

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
	"reflect"
	"strings"
	"testing"
)

func safeTypeName(v interface{}) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).Name()
}

func compareJsonObject(t *testing.T, expected types.JsonObject, actual types.JsonObject, path string, checkExtra bool) {
	for _, key := range expected.Keys() {
		value1 := expected.Get(key)
		if actual.ContainsKey(key) {
			value2 := actual.Get(key)
			newPath := fmt.Sprintf("%s/%s", path, key)
			if v1, ok := value1.(types.JsonObject); ok {
				if v2, ok := value2.(types.JsonObject); ok {
					compareJsonObject(t, v1, v2, newPath, true)
				} else {
					t.Errorf("Field %s is not an object (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(value2))
				}
			} else if v1, ok := value1.(types.JsonArray); ok {
				if v2, ok := value2.(types.JsonArray); ok {
					compareJsonArray(t, v1, v2, newPath)
				} else {
					t.Errorf("Field %s is not an array (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(value2))
				}
			} else if v1, ok := value1.(types.JsonNumber); ok {
				if v2, ok := value2.(types.JsonNumber); ok {
					if v1.FloatValue() != v2.FloatValue() {
						t.Errorf("Field %s is not equal (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(v2))
					}
				} else {
					t.Errorf("Field %s is not a number (expected %s (%s), got %s (%s))", newPath, types.ToString(v1), safeTypeName(v1), types.ToString(value2), safeTypeName(value2))
				}
			} else if types.IsNumber(value1) && types.IsNumber(value2) {
				if types.AsNumber(value1).FloatValue() != types.AsNumber(value2).FloatValue() {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, types.ToString(value1), safeTypeName(value1), types.ToString(value2), safeTypeName(value2))
				}
			} else {
				if value1 != value2 {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, types.ToString(value1), safeTypeName(value1), types.ToString(value2), safeTypeName(value2))
				}
			}
		} else {
			t.Errorf("Object does not contain key %s", key)
		}
	}
	if checkExtra {
		for _, key := range actual.Keys() {
			if !expected.ContainsKey(key) {
				t.Errorf("Object contains unexpected key %s/%s", path, key)
			}
		}
		if actual.Size() == expected.Size() {
			for i := 0; i < actual.Size(); i++ {
				if actual.Keys()[i] != expected.Keys()[i] {
					t.Errorf("Keys are not in the right order at %d (expected %s, got %s)", i, types.ToString(types.AsArray(expected.Keys())), types.ToString(types.AsArray(actual.Keys())))
					break
				}
			}
		}
	}
	if t.Failed() {
		t.Errorf("\nExpected: %s\n  Actual: %s", types.ToString(expected), types.ToString(actual))
	}
}

func compareJsonArray(t *testing.T, expected types.JsonArray, actual types.JsonArray, path string) {
	for i := 0; i < int(math.Min(float64(len(expected.Value)), float64(len(actual.Value)))); i++ {
		newPath := fmt.Sprintf("%s[%d]", path, i)
		value1 := expected.Value[i]
		value2 := actual.Value[i]
		if v1, ok := value1.(types.JsonObject); ok {
			if v2, ok := value2.(types.JsonObject); ok {
				compareJsonObject(t, v1, v2, newPath, true)
			} else {
				t.Errorf("Element %s is not an object (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(value2))
			}
		} else if v1, ok := value1.(types.JsonArray); ok {
			if v2, ok := value2.(types.JsonArray); ok {
				compareJsonArray(t, v1, v2, newPath)
			} else {
				t.Errorf("Element %s is not an array (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(value2))
			}
		} else if v1, ok := value1.(types.JsonNumber); ok {
			if v2, ok := value2.(types.JsonNumber); ok {
				if v1.FloatValue() != v2.FloatValue() {
					t.Errorf("Element %s is not equal (expected %s, got %s)", newPath, types.ToString(v1), types.ToString(v2))
				}
			} else {
				t.Errorf("Element %s is not a number (expected %s (%s), got %s (%s))", newPath, types.ToString(v1), safeTypeName(v1), types.ToString(value2), safeTypeName(value2))
			}
		} else if types.IsNumber(value1) && types.IsNumber(value2) {
			if types.AsNumber(value1).FloatValue() != types.AsNumber(value2).FloatValue() {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, types.ToString(value1), safeTypeName(value1), types.ToString(value2), safeTypeName(value2))
			}
		} else {
			if value1 != value2 {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, types.ToString(value1), safeTypeName(value1), types.ToString(value2), safeTypeName(value2))
			}
		}
	}
	for i := 0; i < len(actual.Value); i++ {
		if i >= len(expected.Value) {
			t.Errorf("Array contains unexpected element %s[%d]", path, i)
		}
	}
}

func assertTemplateWithModule(t *testing.T, template, module, expected string, globalScope types.JsonObject) {
	mod, err := jsonte.LoadModule(module, globalScope, -1)
	if err != nil {
		t.Fatal(err)
	}
	process, err := jsonte.Process("test", template, globalScope, map[string]jsonte.JsonModule{
		mod.Name.StringValue(): mod,
	}, -1)
	if err != nil {
		t.Fatal(err)
	}
	exp, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	compareJsonObject(t, exp, process.Get("test"), "#", true)
}

func assertTemplate(t *testing.T, template, expected string) {
	process, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err != nil {
		t.Fatal(err)
	}
	exp, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	compareJsonObject(t, exp, process.Get("test"), "#", true)
}

func assertTemplateMultiple(t *testing.T, template, expected string) {
	process, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err != nil {
		t.Fatal(err)
	}
	exp, err := types.ParseJsonObject([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range exp.Keys() {
		value := exp.Get(key)
		if !process.ContainsKey(key) {
			t.Errorf("Missing file %s", key)
			continue
		}
		compareJsonObject(t, value.(types.JsonObject), process.Get(key), fmt.Sprintf("%s#", key), true)
	}
	for _, key := range process.Keys() {
		if !exp.ContainsKey(key) {
			t.Errorf("Unexpected file %s", key)
		}
	}
}

func TestSimpleTemplate(t *testing.T) {
	template := `{
		"$template": {
			"test": "{{=1..3}}"
		}
	}`
	expected := `{
		"test": [1, 2, 3]
	}`
	assertTemplate(t, template, expected)
}

func TestSimpleTemplateDifferentCase(t *testing.T) {
	template := `{
		"$Template": {
			"test": "{{=1..3}}"
		}
	}`
	expected := `{
		"test": [1, 2, 3]
	}`
	assertTemplate(t, template, expected)
}

func TestSimpleIterationInObject(t *testing.T) {
	template := `{
		"$template": {
			"{{#1..3}}": {
				"test{{index}}": "{{=value}}"
			}
		}
	}`
	expected := `{
		"test0": 1,
		"test1": 2,
		"test2": 3
	}`
	assertTemplate(t, template, expected)
}

func TestNamedValue(t *testing.T) {
	template := `{
		"$template": {
			"{{#1..3 as i}}": {
				"test{{index}}": "{{=i}}"
			}
		}
	}`
	expected := `{
		"test0": 1,
		"test1": 2,
		"test2": 3
	}`
	assertTemplate(t, template, expected)
}

func TestNamedValueAndIndex(t *testing.T) {
	template := `{
		"$template": {
			"{{#1..3 as v, i}}": {
				"test{{i}}": "{{=v}}"
			}
		}
	}`
	expected := `{
		"test0": 1,
		"test1": 2,
		"test2": 3
	}`
	assertTemplate(t, template, expected)
}

func TestSimpleIterationInArray(t *testing.T) {
	template := `{
		"$template": {
			"test": [
				{
					"{{#1..3}}": {
						"test{{index}}": "{{=value}}"
					}
				}
			]
		}
	}`
	expected := `{
		"test": [
			{
				"test0": 1
			},
			{
				"test1": 2
			},
			{
				"test2": 3
			}
		]
	}`
	assertTemplate(t, template, expected)
}

func TestSimplePredicateInArray(t *testing.T) {
	template := `{
		"$template": {
			"test": [
				{
					"{{?true}}": 1
				},
				{
					"{{?false}}": 2
				}
			]
		}
	}`
	expected := `{
		"test": [
			1
		]
	}`
	assertTemplate(t, template, expected)
}

func TestNestedIterationInArray(t *testing.T) {
	template := `{
		"$template": {
			"test": [
				{
					"{{#1..3 as outer}}": {
						"{{#1..3 as inner}}": {
							"test{{outer}}-{{inner}}": "{{outer}}-{{inner}}"
						}
					}
				}
			]
		}
	}`
	expected := `{
		"test": [
			{
				"test1-1": "1-1"
			},
			{
				"test1-2": "1-2"
			},
			{
				"test1-3": "1-3"
			},
			{
				"test2-1": "2-1"
			},
			{
				"test2-2": "2-2"
			},
			{
				"test2-3": "2-3"
			},
			{
				"test3-1": "3-1"
			},
			{
				"test3-2": "3-2"
			},
			{
				"test3-3": "3-3"
			}
		]
	}`
	assertTemplate(t, template, expected)
}

func TestSimplePredicateInObject(t *testing.T) {
	template := `{
		"$template": {
			"{{?true}}": {
				"test": 1
			}
		}
	}`
	expected := `{
		"test": 1
	}`
	assertTemplate(t, template, expected)
}

func TestSimplePredicateInObject2(t *testing.T) {
	template := `{
		"$template": {
			"{{?false}}": {
				"test": 1
			}
		}
	}`
	expected := "{}"
	assertTemplate(t, template, expected)
}

func TestSimpleIterationInObjectWithName(t *testing.T) {
	template := `{
		"$template": {
			"{{#1..3 as num}}": {
				"test{{index}}": "{{=num}}"
			}
		}
	}`
	expected := `{
		"test0": 1,
		"test1": 2,
		"test2": 3
	}`
	assertTemplate(t, template, expected)
}

func TestSimpleModule(t *testing.T) {
	module := `{
		"$module": "simple",
		"$scope": {
			"asd": 123
		},
		"$template": {
			"asd": "{{=asd}}",
			"overrideMe": -1,
			"objectAndObject": {
				"asd": 123
			},
			"arrayAndArray": [
				1,
				2,
				3
			]
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"overrideMe": 1,
			"objectAndObject": {
				"qwe": 456
			},
			"arrayAndArray": [
				4,
				5,
				6
			]
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1,
		"objectAndObject": {
			"asd": 123,
			"qwe": 456
		},
		"arrayAndArray": [
			1,
			2,
			3,
			4,
			5,
			6
		]
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
}

func TestSimpleModuleDifferentCase(t *testing.T) {
	module := `{
		"$Module": "simple",
		"$Scope": {
			"asd": 123
		},
		"$Template": {
			"asd": "{{=asd}}",
			"overrideMe": -1
		}
	}`
	template := `{
		"$Extend": "simple",
		"$Template": {
			"overrideMe": 1
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
}

func TestSimpleCopy(t *testing.T) {
	file := `{
		"asd": 123,
		"overrideMe": -1
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$copy": "file.json",
		"$template": {
			"overrideMe": 1,
			"oldOverrideMe": "{{=$copy.overrideMe}}"
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1,
		"oldOverrideMe": -1
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestSimpleCopyDifferentCase(t *testing.T) {
	file := `{
		"asd": 123,
		"overrideMe": -1
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$Copy": "file.json",
		"$Template": {
			"overrideMe": 1,
			"oldOverrideMe": "{{=$copy.overrideMe}}"
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1,
		"oldOverrideMe": -1
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestTemplateCopy(t *testing.T) {
	file := `{
		"$scope": {
			"asd": 123
		},
		"$template": {
			"asd": "{{=asd}}",
			"overrideMe": -1
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.templ": file,
	}, false)
	template := `{
		"$copy": "file.templ",
		"$template": {
			"overrideMe": 1
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestDeleteNullsInCopy(t *testing.T) {
	file := `{
		"asd": 123,
		"overrideMe": -1,
		"removeMe": {
			"removeMeToo": "asd"
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$copy": "file.json",
		"$template": {
			"overrideMe": null,
			"removeMe": null
		}
	}`
	expected := `{
		"asd": 123
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

// The following test is for a bug, that would sometimes skip an element to remove, if it was right after another removed element.
func TestCleanEntity(t *testing.T) {
	file := `{
  "format_version": "1.8.0",
  "minecraft:client_entity": {
    "description": {
      "identifier": "minecraft:parrot",
      "materials": { "default": "parrot"},
      "textures": {
        "blue": "textures/entity/parrot/parrot_blue",
        "green": "textures/entity/parrot/parrot_green",
        "red_blue": "textures/entity/parrot/parrot_red_blue",
        "yellow_blue": "textures/entity/parrot/parrot_yellow_blue",
        "grey": "textures/entity/parrot/parrot_grey"
      },
      "geometry": {
        "default": "geometry.parrot"
      },
      "scripts": {
        "pre_animation": [
          "variable.state = query.is_dancing ? 3 : (query.is_sitting ? 2 : (!query.is_on_ground && !query.is_jumping && !query.is_riding ? 0 : 1));",
          "variable.dance.x = Math.cos(query.life_time * 57.3 * 20.0);",
          "variable.dance.y = -Math.sin(query.life_time * 57.3 * 20.0);",
          "variable.wing_flap = ((math.sin(query.wing_flap_position * 57.3) + 1) * query.wing_flap_speed);"
        ]
      },
      "animations": {
        "moving": "animation.parrot.moving",
        "base": "animation.parrot.base",
        "dance": "animation.parrot.dance",
        "sitting": "animation.parrot.sitting",
        "flying": "animation.parrot.flying",
        "standing": "animation.parrot.standing",
        "look_at_target": "animation.common.look_at_target"
      },
      "animation_controllers": [
        { "setup": "controller.animation.parrot.setup" },
        { "move": "controller.animation.parrot.move" }
      ],
      "render_controllers": [ "controller.render.parrot" ],
      "spawn_egg": {
        "texture": "spawn_egg",
        "texture_index": 43
      }
    }
  }
}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$copy": "file.json",
		"$template": {
        "minecraft:client_entity": {
            "description": {
                "identifier": "example:parrot",
                "scripts": null,
                "animation_controllers": null,
                "animations": null,
                "particle_effects": null,
                "spawn_egg": null
            }
        }
		}
	}`
	expected := `{
  "format_version": "1.8.0",
  "minecraft:client_entity": {
    "description": {
      "identifier": "example:parrot",
      "materials": {
        "default": "parrot"
      },
      "textures": {
        "blue": "textures/entity/parrot/parrot_blue",
        "green": "textures/entity/parrot/parrot_green",
        "red_blue": "textures/entity/parrot/parrot_red_blue",
        "yellow_blue": "textures/entity/parrot/parrot_yellow_blue",
        "grey": "textures/entity/parrot/parrot_grey"
      },
      "geometry": {
        "default": "geometry.parrot"
      },
      "render_controllers": [
        "controller.render.parrot"
      ]
    }
  }
}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestDeleteNulls(t *testing.T) {
	file := `{
		"$scope": {
			"asd": 123
		},
		"$template": {
			"asd": "{{=asd}}",
			"overrideMe": -1,
			"removeMe": {
				"removeMeToo": "asd"
			}
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.templ": file,
	}, false)
	template := `{
		"$copy": "file.templ",
		"$template": {
			"overrideMe": null,
			"removeMe": null
		}
	}`
	expected := `{
		"asd": 123
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestDeleteNulls2(t *testing.T) {
	file := `{
		"asd": 1,
		"asd2": 2,
		"asd3": 3
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$scope": {
			"del": ["asd", "asd2"]
		},
		"$copy": "file.json",
		"$template": {
			"{{#del}}": {
				"{{value}}": null
			}
		}
	}`
	expected := `{
		"asd3": 3
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestCopyAndExtend(t *testing.T) {
	file := `{
		"asd": 123,
		"overrideMe": -1,
		"removeMe": 1,
		"removeMe1": {
			"removeMe2": 1
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	module := `{
		"$module": "simple",
		"$template": {
			"removeMe": null
		}
	}`
	template := `{
		"$copy": "file.json",
		"$extend": "simple",
		"$template": {
			"overrideMe": 1,
			"{{?true}}": {
				"$removeMe1": null
			}
		}
	}`
	expected := `{
		"asd": 123,
		"overrideMe": 1
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestReplaceInTemplate(t *testing.T) {
	file := `{
		"example": ["one"]
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	module := `{
		"$module": "simple",
		"$template": {
			"$example": ["four","five"]
		}
	}`
	template := `{
	"$copy": "file.json",
	"$extend": "simple",
	"$template": {
		"{{?true}}": {
			"$example": ["two", "three"]
		}
	}
}`
	expected := `{
		"example": ["two", "three"]
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestMultipleFiles(t *testing.T) {
	template := `{
		"$files": {
			"fileName": "file{{index}}",
			"array": "{{1..3}}"
		},	
		"$template": {
			"test": "{{=index}}"
		}
	}`
	expected := `{
		"file0": {
			"test": 0
		},
		"file1": {
			"test": 1
		},
		"file2": {
			"test": 2
		}
	}`
	assertTemplateMultiple(t, template, expected)
}

func TestMultipleFilesDifferentCase(t *testing.T) {
	template := `{
		"$Files": {
			"file_name": "file{{index}}",
			"Array": "{{1..3}}"
		},	
		"$Template": {
			"test": "{{=index}}"
		}
	}`
	expected := `{
		"file0": {
			"test": 0
		},
		"file1": {
			"test": 1
		},
		"file2": {
			"test": 2
		}
	}`
	assertTemplateMultiple(t, template, expected)
}

func TestEmptyArray(t *testing.T) {
	template := `{
		"$template": {
			"test": ["{{[]}}"]
		}
	}`
	expected := `{
		"test": []
	}`
	assertTemplate(t, template, expected)
}

func TestKeepTypes(t *testing.T) {
	template := `{
		"$template": {
			"decimal": 0.0,
			"integer": 0,
			"string": "",
			"array": [],
			"object": {},
			"null": null,
			"true": true,
			"false": false,
			"templatedDecimal": "{{0.0}}",
			"templatedInteger": "{{0}}",
			"templatedString": "{{\"\"}}",
			"templatedArray": "{{[]}}"
		}
	}`
	expected := `{
		"decimal": 0.0,
		"integer": 0,
		"string": "",
		"array": [],
		"object": {},
		"true": true,
		"false": false,
		"templatedDecimal": 0.0,
		"templatedInteger": 0,
		"templatedString": "",
		"templatedArray": []
	}`
	assertTemplate(t, template, expected)
}

func TestModuleOverride(t *testing.T) {
	module := `{
		"$module": "simple",
		"$template": {
			"groups": {
				"$object_1": {
					"value": 0
				}
			}
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"groups": {
				"{{#1..3}}": {
					"object_{{value}}": {
						"value": "{{=value}}"
					}
				}
			}
		}
	}`
	expected := `{
		"groups": {
			"object_1": {
				"value": 0
			},
			"object_2": {
				"value": 2
			},
			"object_3": {
				"value": 3
			}
		}
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
}

func TestJsonParser(t *testing.T) {
	template := `/*comment*/{
		"obj": {
			// This is a comment!
			"decimal": 0.0,
			"integer": 0,
/* block */	"string": "escape chars \n\t\r\b\f \\ \" \u1234",//
			"array": [],
			"object": {},
			"null": null,
			"true": true,
			"false": false
		}
	}`
	expected := utils.NewNavigableMap[string, interface{}]()
	obj := utils.NewNavigableMap[string, interface{}]()
	obj.Put("decimal", 0.0)
	obj.Put("integer", 0)
	obj.Put("string", "escape chars \n\t\r\b\f \\ \" \u1234")
	obj.Put("array", []interface{}{})
	obj.Put("object", utils.NewNavigableMap[string, interface{}]())
	obj.Put("null", nil)
	obj.Put("true", true)
	obj.Put("false", false)
	expected.Put("obj", obj)

	object, err := types.ParseJsonObject([]byte(template))
	if err != nil {
		t.Error(err)
	}
	compareJsonObject(t, object, types.Box(expected).(types.JsonObject), "#", true)

	expMini := "{\"obj\":{\"decimal\":0.0,\"integer\":0,\"string\":\"escape chars \\n\\t\\r\\b\\f \\\\ \\\" ሴ\",\"array\":[],\"object\":{},\"null\":null,\"true\":true,\"false\":false}}"
	expPretty := `{
  "obj": {
    "decimal": 0.0,
    "integer": 0,
    "string": "escape chars \n\t\r\b\f \\ \" ሴ",
    "array": [
    ],
    "object": {
    },
    "null": null,
    "true": true,
    "false": false
  }
}`
	if types.ToString(object) != expMini {
		t.Error("Unexpected string representation of object")
		t.Errorf("Expected: %s", expMini)
		t.Errorf("Actual: %s", types.ToString(object))
	}
	if types.ToPrettyString(object) != expPretty {
		t.Error("Unexpected string representation of object")
		t.Errorf("Expected: %s", expPretty)
		t.Errorf("Actual: %s", types.ToPrettyString(object))
	}
}

func TestCorrectAssert(t *testing.T) {
	template := `{
		"$template": {
			"$assert": "true"
		}
	}`
	expected := `{}`
	assertTemplate(t, template, expected)
}

func TestIncorrectAssert(t *testing.T) {
	template := `{
		"$template": {
			"$assert": "false"
		}
	}`
	_, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if strings.Trim(strings.Split(err.Error(), "\n")[0], "\r\n \t") != "Assertion failed for 'false' at $template/$assert" {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestIncorrectAssertDifferentCase(t *testing.T) {
	template := `{
		"$template": {
			"$Assert": "false"
		}
	}`
	_, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if strings.Trim(strings.Split(err.Error(), "\n")[0], "\r\n \t") != "Assertion failed for 'false' at $template/$assert" {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestIncorrectAssertWithMessage(t *testing.T) {
	template := `{
		"$template": {
			"$assert": {
				"condition": "false",
				"message": "This is a test"
			}
		}
	}`
	_, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if strings.Trim(strings.Split(err.Error(), "\n")[0], "\r\n \t") != "This is a test at $template/$assert" {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestIncorrectAssertWithMessageDifferentCase(t *testing.T) {
	template := `{
		"$template": {
			"$Assert": {
				"Condition": "false",
				"Message": "This is a test"
			}
		}
	}`
	_, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if strings.Trim(strings.Split(err.Error(), "\n")[0], "\r\n \t") != "This is a test at $template/$assert" {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestMultipleAsserts(t *testing.T) {
	template := `{
		"$template": {
			"$assert": ["true", "true", "false", "true"]
		}
	}`
	_, err := jsonte.Process("test", template, types.NewJsonObject(), map[string]jsonte.JsonModule{}, -1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if strings.Trim(strings.Split(err.Error(), "\n")[0], "\r\n \t") != "Assertion failed for 'false' at $template/$assert[2]" {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestMultiplePredicatesModifyingSameField(t *testing.T) {
	template := `{
		"$template": {
			"{{?true}}": {
				"field": {
					"set1": "value1"
				}
			},
			"{{?true == true}}": {
				"field": {
					"set2": "value2"
				}
			}
		}
	}`
	expected := `{
		"field": {
			"set1": "value1",
			"set2": "value2"
		}
	}`
	assertTemplate(t, template, expected)
}

func TestIteration(t *testing.T) {
	template := `{
	"$scope": {
		"arr": [
			{
				"name": "value1",
				"dummy": "dummy1"
			},
			{
				"name": "value2",
				"dummy": false
			},
			{
				"name": "value3"
			}
		]
	},
	"$template": {
		"animation_controllers": {
			"{{#arr.filter(x => !x?.dummy)}}": {
				"controller.animation.test.{{name}}": {
					"asd": "asd"
				}
			}
		}
	}
}`
	expected := `{
  "animation_controllers": {
    "controller.animation.test.value2": {
      "asd": "asd"
    },
    "controller.animation.test.value3": {
      "asd": "asd"
    }
  }
}`
	assertTemplate(t, template, expected)
}

func TestTemplateScope(t *testing.T) {
	template := `{
		"$scope": {
			"asd": "{{(1..10).map(x => x * 2)}}"
		},
		"$template": {
			"data": "{{asd}}"
		}
	}`
	expected := `{
		"data": [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
	}`
	assertTemplate(t, template, expected)
}

func TestTemplateModuleScope(t *testing.T) {
	module := `{
		"$module": "simple",
		"$scope": {
			"asd": "{{globalData}}"
		},
		"$template": {
			"asd": "{{=asd}}"
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"qwe": "{{=asd}}"
		}
	}`
	expected := `{
		"asd": 5
	}`
	assertTemplateWithModule(t, template, module, expected, types.AsObject(map[string]interface{}{
		"globalData": 5,
	}))
}

func TestArrayMergeAndOverride(t *testing.T) {
	module := `{
		"$module": "simple",
		"$template": {
			"arr1": [1, 2, 3],
			"arr2": [1, 2, 3],
			"arr3": [4, 5, 6],
			"arr4": [3, 4],
			"arr5": [3, 4]
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"arr1": [4, 5, 6],
			"$arr2": [1, 2, 3, 4, 5, 6],
			"^arr3": [1, 2, 3],
			"^arr4": [1, 2],
			"arr4": [5, 6],
			"arr5": [5, 6],
			"^arr5": [1, 2]
		}
	}`
	expected := `{
		"arr1": [1, 2, 3, 4, 5, 6],
		"arr2": [1, 2, 3, 4, 5, 6],
		"arr3": [1, 2, 3, 4, 5, 6],
		"arr4": [1, 2, 3, 4, 5, 6],
		"arr5": [1, 2, 3, 4, 5, 6]
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
}

func TestOverrideNestedFields(t *testing.T) {
	file := `{
		"minecraft:client_entity": {
			"description": {
				"identifier": "minecraft:cow",
				"scripts": {
					"animate": ["test"]
				}
			}
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	template := `{
		"$copy": "file.json",
		"$template": {
			"minecraft:client_entity": {
				"description": {
					"{{?true}}": {
						"$scripts": {
							"$animate": ["test2"]
						}
					}
				}
			}
		}
	}`
	expected := `{
		"minecraft:client_entity": {
			"description": {
				"identifier": "minecraft:cow",
				"scripts": {
					"animate": ["test2"]
				}
			}
		}
	}`
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestConditionalReplaceArray(t *testing.T) {
	file := `{
		"example": ["one"]
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.json": file,
	}, false)
	module := `{
		"$module": "simple",
		"$template": {
			"$example": ["four", "five"]
		}
	}`
	template := `{
		"$copy": "file.json",
		"$extend": "simple",
		"$template": {
        	"{{?true}}": {
				"$example": ["two", "three"]
			}
		}
	}`
	expected := `{
		"example": ["two", "three"]
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestModuleReplacingTemplate(t *testing.T) {
	module := `{
		"$module": "simple",
		"$template": {
			"$example": ["four", "five"]
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"example": ["two", "three"]
		}
	}`
	expected := `{
		"example": ["four", "five"]
	}`
	assertTemplateWithModule(t, template, module, expected, types.NewJsonObject())
	safeio.Resolver = safeio.DefaultIOResolver
}
