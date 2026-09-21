package test

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
)

func TestMolangFieldMinified(t *testing.T) {
	template := `{
		"$template": {
			"$molang:script": "query.life_time > 1.0 ? variable.x : 0"
		}
	}`
	expected := `{
		"script": "q.life_time>1?v.x:0"
	}`
	assertTemplate(t, template, expected)
}

func TestMolangFieldWithExpression(t *testing.T) {
	template := `{
		"$template": {
			"$molang:script": "variable.x = {{1 + 1}};"
		}
	}`
	expected := `{
		"script": "v.x=2;"
	}`
	assertTemplate(t, template, expected)
}

// An array of strings is one program split over entries, so it collapses into a single string
func TestMolangFieldArrayJoined(t *testing.T) {
	template := `{
		"$template": {
			"$molang:pre_animation": [
				"variable.x = 1.0;",
				"variable.y = query.life_time;"
			]
		}
	}`
	expected := `{
		"pre_animation": "v.x=1;v.y=q.life_time;"
	}`
	assertTemplate(t, template, expected)
}

// Entries may split a single expression, so they are validated together, not one by one
func TestMolangFieldArraySplitExpression(t *testing.T) {
	template := `{
		"$template": {
			"$molang:script": [
				"variable.x = ",
				"1.0;"
			]
		}
	}`
	expected := `{
		"script": "v.x=1;"
	}`
	assertTemplate(t, template, expected)
}

// Entries are joined with a newline, so that neighbouring tokens do not merge
func TestMolangFieldArrayNotGlued(t *testing.T) {
	template := `{
		"$template": {
			"$molang:script": [
				"variable.x = 1",
				"0;"
			]
		}
	}`
	assertTemplateError(t, template, []string{
		"Invalid Molang: expected ';' or end of block, got NUMBER (line 2, column 1: 0;) at $template/$molang:script",
	})
}

// An array that is not all strings is not a program, so its values are handled separately
func TestMolangFieldArrayOfObjects(t *testing.T) {
	template := `{
		"$template": {
			"$molang:part_visibility": [
				{"head": "query.is_baby"},
				{"body": "variable.x > 1.0"}
			]
		}
	}`
	expected := `{
		"part_visibility": [
			{"head": "q.is_baby"},
			{"body": "v.x>1"}
		]
	}`
	assertTemplate(t, template, expected)
}

func TestMolangFieldInObject(t *testing.T) {
	template := `{
		"$template": {
			"$molang:scripts": {
				"scale": "variable.size * 1.0",
				"nested": {
					"value": "query.is_baby ? 0.5 : 1.0"
				}
			}
		}
	}`
	expected := `{
		"scripts": {
			"scale": "v.size*1",
			"nested": {
				"value": "q.is_baby?0.5:1"
			}
		}
	}`
	assertTemplate(t, template, expected)
}

// Molang fields in JSON may hold numbers and empty strings, which are left alone
func TestMolangFieldNonStringValues(t *testing.T) {
	template := `{
		"$template": {
			"$molang:rotation": 0,
			"$molang:computed": "{{=1 + 1}}",
			"$molang:empty": ""
		}
	}`
	expected := `{
		"rotation": 0,
		"computed": 2,
		"empty": ""
	}`
	assertTemplate(t, template, expected)
}

func TestMolangFieldInArrayElement(t *testing.T) {
	template := `{
		"$template": {
			"parts": [
				{
					"$molang:condition": "query.is_baby"
				}
			]
		}
	}`
	expected := `{
		"parts": [{"condition": "q.is_baby"}]
	}`
	assertTemplate(t, template, expected)
}

func TestMolangFieldTemplatedKey(t *testing.T) {
	template := `{
		"$template": {
			"$molang:{{'scr' + 'ipt'}}": "variable.x = 1.0;"
		}
	}`
	expected := `{
		"script": "v.x=1;"
	}`
	assertTemplate(t, template, expected)
}

// The prefix must survive module merging, which strips $ from override keys
func TestMolangFieldFromModule(t *testing.T) {
	module := `{
		"$module": "molangModule",
		"$template": {
			"$molang:script": "variable.x = 1.0;"
		}
	}`
	template := `{
		"$extend": "molangModule",
		"$template": {
			"$molang:other": "variable.y = 2.0;"
		}
	}`
	expected := `{
		"script": "v.x=1;",
		"other": "v.y=2;"
	}`
	assertTemplateWithModule(t, template, expected, nil, module)
}

func TestMolangFieldInvalid(t *testing.T) {
	template := `{
		"$template": {
			"$molang:script": "variable.x = "
		}
	}`
	assertTemplateError(t, template, []string{
		"Invalid Molang: unexpected token EOF (line 1, column 14: variable.x =) at $template/$molang:script",
	})
}

func TestMolangFieldInvalidWithoutValidation(t *testing.T) {
	jsonte.SetMolangValidation(false)
	t.Cleanup(func() { jsonte.SetMolangValidation(true) })
	template := `{
		"$template": {
			"$molang:script": "variable.x = "
		}
	}`
	expected := `{
		"script": "v.x="
	}`
	assertTemplate(t, template, expected)
}
