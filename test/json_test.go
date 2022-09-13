package test

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"math"
	"reflect"
	"testing"
)

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
					t.Errorf("Field %s is not a number (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
				}
			} else if value1 != value2 {
				t.Errorf("Field %s is not equal (expected %s (%s), got %s (%s)", newPath, utils.ToString(value1), reflect.TypeOf(value1).Name(), utils.ToString(value2), reflect.TypeOf(value2).Name())
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
				t.Errorf("Element %s is not a number (expected %s, got %s)", newPath, utils.ToString(v1), utils.ToString(value2))
			}
		} else {
			if value1 != value2 {
				t.Errorf("Element %s is not equal (expected %s (%s), got %s (%s))", newPath, utils.ToString(value1), reflect.TypeOf(value1).Name(), utils.ToString(value2), reflect.TypeOf(value2).Name())
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"file.json": []byte(file),
	})
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"file.templ": []byte(file),
	})
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"file.templ": []byte(file),
	})
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"file.json": []byte(file),
	})
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
