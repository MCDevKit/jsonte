package test

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
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

func compareJsonObject(t *testing.T, expected utils.JsonObject, actual utils.JsonObject, path string) {
	t.Helper()
	for key, value1 := range expected {
		if value2, ok := actual[key]; ok {
			newPath := fmt.Sprintf("%s/%s", path, key)
			if v1, ok := value1.(utils.JsonObject); ok {
				if v2, ok := value2.(utils.JsonObject); ok {
					compareJsonObject(t, v1, v2, newPath)
				} else {
					t.Errorf("Field %s is not an object (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
				}
			} else if v1, ok := value1.(utils.JsonArray); ok {
				if v2, ok := value2.(utils.JsonArray); ok {
					compareJsonArray(t, v1, v2, newPath)
				} else {
					t.Errorf("Field %s is not an array (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
				}
			} else if v1, ok := value1.(utils.JsonNumber); ok {
				if v2, ok := value2.(utils.JsonNumber); ok {
					if v1.FloatValue() != v2.FloatValue() {
						t.Errorf("Field %s is not equal (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(v2))
					}
				} else {
					t.Errorf("Field %s is not a number (expected %s (%s), got %s (%s))", newPath, utils.ToString(v1), safeTypeName(v1), utils.ToString(value2), safeTypeName(value2))
				}
			} else if utils.IsNumber(value1) && utils.IsNumber(value2) {
				if utils.ToNumber(value1).FloatValue() != utils.ToNumber(value2).FloatValue() {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
				}
			} else {
				if value1 != value2 {
					t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
				}
			}
		} else {
			t.Errorf("Object does not contain key %s", key)
		}
	}
	for key := range actual {
		if _, ok := expected[key]; !ok {
			t.Errorf("Object contains unexpected key %s/%s", path, key)
		}
	}
}

func compareJsonArray(t *testing.T, expected utils.JsonArray, actual utils.JsonArray, path string) {
	t.Helper()
	for i := 0; i < int(math.Min(float64(len(expected)), float64(len(actual)))); i++ {
		newPath := fmt.Sprintf("%s[%d]", path, i)
		value1 := expected[i]
		value2 := actual[i]
		if v1, ok := value1.(utils.JsonObject); ok {
			if v2, ok := value2.(utils.JsonObject); ok {
				compareJsonObject(t, v1, v2, newPath)
			} else {
				t.Errorf("Element %s is not an object (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
			}
		} else if v1, ok := value1.(utils.JsonArray); ok {
			if v2, ok := value2.(utils.JsonArray); ok {
				compareJsonArray(t, v1, v2, newPath)
			} else {
				t.Errorf("Element %s is not an array (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
			}
		} else if v1, ok := value1.(utils.JsonNumber); ok {
			if v2, ok := value2.(utils.JsonNumber); ok {
				if v1.FloatValue() != v2.FloatValue() {
					t.Errorf("Element %s is not equal (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(v2))
				}
			} else {
				t.Errorf("Element %s is not a number (expected %s (%s), got %s (%s))", newPath, utils.ToString(v1), safeTypeName(v1), utils.ToString(value2), safeTypeName(value2))
			}
		} else if utils.IsNumber(value1) && utils.IsNumber(value2) {
			if utils.ToNumber(value1).FloatValue() != utils.ToNumber(value2).FloatValue() {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
			}
		} else {
			if value1 != value2 {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), safeTypeName(value1), utils.ToString(value2), safeTypeName(value2))
			}
		}
	}
	for i := 0; i < len(actual); i++ {
		if i >= len(expected) {
			t.Errorf("Array contains unexpected element %s[%d]", path, i)
		}
	}
}

func assertTemplateWithModule(t *testing.T, template, module string, expected utils.JsonObject) {
	t.Helper()
	mod, err := jsonte.LoadModule(module)
	if err != nil {
		t.Fatal(err)
	}
	process, err := jsonte.Process("test", template, utils.JsonObject{}, map[string]jsonte.JsonModule{
		mod.Name: mod,
	}, -1)
	if err != nil {
		t.Fatal(err)
	}
	compareJsonObject(t, expected, process["test"].(utils.JsonObject), "#")
}

func assertTemplate(t *testing.T, template string, expected utils.JsonObject) {
	t.Helper()
	process, err := jsonte.Process("test", template, utils.JsonObject{}, map[string]jsonte.JsonModule{}, -1)
	if err != nil {
		t.Fatal(err)
	}
	compareJsonObject(t, expected, process["test"].(utils.JsonObject), "#")
}

func assertTemplateMultiple(t *testing.T, template string, expected map[string]utils.JsonObject) {
	t.Helper()
	process, err := jsonte.Process("test", template, utils.JsonObject{}, map[string]jsonte.JsonModule{}, -1)
	if err != nil {
		t.Fatal(err)
	}
	for key, value := range expected {
		if _, ok := process[key]; !ok {
			t.Errorf("Missing file %s", key)
			continue
		}
		compareJsonObject(t, value, process[key].(utils.JsonObject), fmt.Sprintf("%s#", key))
	}
	for key := range process {
		if _, ok := expected[key]; !ok {
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
	expected := utils.JsonObject{
		"test": utils.JsonArray{1, 2, 3},
	}
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
	expected := utils.JsonObject{
		"test0": 1,
		"test1": 2,
		"test2": 3,
	}
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
	expected := utils.JsonObject{
		"test": utils.JsonArray{
			utils.JsonObject{
				"test0": 1,
			},
			utils.JsonObject{
				"test1": 2,
			},
			utils.JsonObject{
				"test2": 3,
			},
		},
	}
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
	expected := utils.JsonObject{
		"test": utils.JsonArray{
			1,
		},
	}
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
	expected := utils.JsonObject{
		"test": utils.JsonArray{
			utils.JsonObject{
				"test1-1": "1-1",
			},
			utils.JsonObject{
				"test1-2": "1-2",
			},
			utils.JsonObject{
				"test1-3": "1-3",
			},
			utils.JsonObject{
				"test2-1": "2-1",
			},
			utils.JsonObject{
				"test2-2": "2-2",
			},
			utils.JsonObject{
				"test2-3": "2-3",
			},
			utils.JsonObject{
				"test3-1": "3-1",
			},
			utils.JsonObject{
				"test3-2": "3-2",
			},
			utils.JsonObject{
				"test3-3": "3-3",
			},
		},
	}
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
	expected := utils.JsonObject{
		"test": 1,
	}
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
	expected := utils.JsonObject{}
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
	expected := utils.JsonObject{
		"test0": 1,
		"test1": 2,
		"test2": 3,
	}
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
			"overrideMe": -1
		}
	}`
	template := `{
		"$extend": "simple",
		"$template": {
			"overrideMe": 1
		}
	}`
	expected := utils.JsonObject{
		"overrideMe": 1,
		"asd":        123,
	}
	assertTemplateWithModule(t, template, module, expected)
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
			"overrideMe": 1
		}
	}`
	expected := utils.JsonObject{
		"overrideMe": 1,
		"asd":        123,
	}
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
	expected := utils.JsonObject{
		"overrideMe": 1,
		"asd":        123,
	}
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
			"overrideMe": -1
		}
	}`
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"file.templ": file,
	}, false)
	template := `{
		"$copy": "file.templ",
		"$template": {
			"overrideMe": null
		}
	}`
	expected := utils.JsonObject{
		"asd": 123,
	}
	assertTemplate(t, template, expected)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestCopyAndExtend(t *testing.T) {
	file := `{
		"asd": 123,
		"overrideMe": -1,
		"removeMe": 1
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
			"overrideMe": 1
		}
	}`
	expected := utils.JsonObject{
		"asd":        123,
		"overrideMe": 1,
	}
	assertTemplateWithModule(t, template, module, expected)
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
	expected := map[string]utils.JsonObject{
		"file0": {
			"test": 0,
		},
		"file1": {
			"test": 1,
		},
		"file2": {
			"test": 2,
		},
	}
	assertTemplateMultiple(t, template, expected)
}

func TestEmptyArray(t *testing.T) {
	template := `{
		"$template": {
			"test": ["{{[]}}"]
		}
	}`
	expected := utils.JsonObject{
		"test": utils.JsonArray{},
	}
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
	expected := utils.JsonObject{
		"decimal":          0.0,
		"integer":          0,
		"string":           "",
		"array":            utils.JsonArray{},
		"object":           utils.JsonObject{},
		"true":             true,
		"false":            false,
		"templatedDecimal": 0.0,
		"templatedInteger": 0,
		"templatedString":  "",
		"templatedArray":   utils.JsonArray{},
	}
	assertTemplate(t, template, expected)
}

func TestTrailingCommaError(t *testing.T) {
	template := `{
		"$template": {
			"test": "asd",
		},
	}`
	_, err := utils.ParseJson([]byte(template))
	if err == nil {
		t.Error("Expected error")
	}
	s := err.Error()
	if !strings.HasPrefix(s, "Most likely trailing comma at line 3\n") {
		t.Error("Expected trailing comma error")
	}
}

func TestNoQuotesObjectKey(t *testing.T) {
	template := `{
		"$template": {
			test: "asd",
		},
	}`
	_, err := utils.ParseJson([]byte(template))
	if err == nil {
		t.Error("Expected error")
	}
	s := err.Error()
	if !strings.HasPrefix(s, "Most likely missing quote at line 3, column 5\n") {
		t.Error("Expected quote error")
	}
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
	expected := utils.JsonObject{
		"groups": utils.JsonObject{
			"object_1": utils.JsonObject{
				"value": 0,
			},
			"object_2": utils.JsonObject{
				"value": 2,
			},
			"object_3": utils.JsonObject{
				"value": 3,
			},
		},
	}
	assertTemplateWithModule(t, template, module, expected)
}
