package test

import (
	"errors"
	"strings"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func assertMolang(t *testing.T, template, expected string) {
	t.Helper()
	process, err := jsonte.ProcessMolangFile(template, types.NewJsonObject())
	if err != nil {
		t.Fatal(err)
	}
	if process != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, process)
	}
}

func assertMolangWithScope(t *testing.T, template, expected string, scope map[string]interface{}) {
	t.Helper()
	obj := types.NewJsonObject()
	for k, v := range scope {
		switch val := v.(type) {
		case string:
			obj.Put(k, types.AsString(val))
		case int:
			obj.Put(k, types.AsNumber(float64(val)))
		case float64:
			obj.Put(k, types.AsNumber(val))
		case bool:
			obj.Put(k, types.AsBool(val))
		}
	}
	process, err := jsonte.ProcessMolangFile(template, obj)
	if err != nil {
		t.Fatal(err)
	}
	if process != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, process)
	}
}

// No expressions - file should only be minified
func TestNoOperationMolang(t *testing.T) {
	input := `variable.x = 1.0; return variable.x;`
	expected := `v.x=1;return v.x;`
	assertMolang(t, input, expected)
}

// Simple #{ } expression substitution
func TestSimpleMolang(t *testing.T) {
	input := `variable.speed = #{1 + 1};`
	expected := `v.speed=2;`
	assertMolang(t, input, expected)
}

// Expression substitution with a string value
func TestMolangStringExpression(t *testing.T) {
	input := `variable.name = '#{'hello'}';`
	expected := `v.name='hello';`
	assertMolang(t, input, expected)
}

// Minification: long accessors are replaced with short forms
func TestMolangShortAccessors(t *testing.T) {
	input := `query.life_time < 0.01 ? variable.x : temp.y`
	expected := `q.life_time<0.01?v.x:t.y`
	assertMolang(t, input, expected)
}

// All four long accessors are shortened
func TestMolangAllShortAccessors(t *testing.T) {
	input := `query.foo + variable.bar + context.baz + temp.qux`
	expected := `q.foo+v.bar+c.baz+t.qux`
	assertMolang(t, input, expected)
}

// # comments are stripped by the minifier
func TestMolangCommentsStripped(t *testing.T) {
	input := `# this is a comment
variable.x = 1.0;
return variable.x;`
	expected := `v.x=1;return v.x;`
	assertMolang(t, input, expected)
}

// Expression resolved from scope
func TestMolangWithScope(t *testing.T) {
	input := `variable.speed = #{speed};`
	expected := `v.speed=42;`
	assertMolangWithScope(t, input, expected, map[string]interface{}{"speed": 42})
}

// Multiple expressions in a single file
func TestMolangMultipleExpressions(t *testing.T) {
	input := `variable.x = #{1 + 1}; variable.y = #{2 * 3};`
	expected := `v.x=2;v.y=6;`
	assertMolang(t, input, expected)
}

// Multi-line expression
func TestMolangMultilineExpression(t *testing.T) {
	input := `variable.result = #{
	(1..5).reduce((a, b) => a + b, 0)
};`
	expected := `v.result=15;`
	assertMolang(t, input, expected)
}

// Nested braces inside an expression do not terminate it prematurely
func TestMolangNestedBraces(t *testing.T) {
	input := `variable.x = #{{'a': 1}['a']};`
	expected := `v.x=1;`
	assertMolang(t, input, expected)
}

// String contents inside #{ } are not parsed as expression tokens
func TestMolangStringWithBrace(t *testing.T) {
	input := `#{'hello world'}`
	expected := `hello world`
	// "hello world" is not valid Molang, so this checks the expression handling with validation disabled
	process, err := jsonte.ProcessMolangFileWithValidation(input, types.NewJsonObject(), false)
	if err != nil {
		t.Fatal(err)
	}
	if process != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, process)
	}
}

// Empty file
func TestMolangEmptyFile(t *testing.T) {
	assertMolang(t, "", "")
}

// File with no jsonte expressions - minification only
func TestMolangOnlyMinify(t *testing.T) {
	input := `
		query.life_time < 0.01 ? 0.0 : variable.hand_bob + 1;
		return variable.hand_bob;
	`
	expected := `q.life_time<0.01?0:v.hand_bob+1;return v.hand_bob;`
	assertMolang(t, input, expected)
}

// jsonte expression result is fed through the minifier afterward
func TestMolangExpressionThenMinify(t *testing.T) {
	input := `query.life_time < #{0.01};`
	expected := `q.life_time<0.01;`
	assertMolang(t, input, expected)
}

// Invalid expression - function returns an error
func TestMolangInvalidExpression(t *testing.T) {
	input := `variable.x = #{1 +};`
	_, err := jsonte.ProcessMolangFile(input, types.NewJsonObject())
	if err == nil {
		t.Fatal("Expected error for invalid expression, got nil")
	}
}

// Scope value substitution in multi-line file
func TestMolangWithScopeMultiline(t *testing.T) {
	input := "variable.speed = #{speed};\nreturn variable.speed;"
	expected := "v.speed=1.5;return v.speed;"
	assertMolangWithScope(t, input, expected, map[string]interface{}{"speed": 1.5})
}

func TestMolangLineComment(t *testing.T) {
	input := "variable.x = 1.0; # this is a comment\nreturn variable.x;"
	expected := "v.x=1;return v.x;"
	assertMolang(t, input, expected)
}

func TestMolangLineCommentOnly(t *testing.T) {
	input := "# full line comment\nvariable.x = 1.0;\nreturn variable.x;"
	expected := "v.x=1;return v.x;"
	assertMolang(t, input, expected)
}

func TestMolangHashBraceNotStripped(t *testing.T) {
	// #{expr} must not be treated as a comment
	input := `query.life_time < #{0.5};`
	expected := `q.life_time<0.5;`
	assertMolang(t, input, expected)
}

func TestMolangCommentWithExpression(t *testing.T) {
	// comment after a real #{} expression
	input := "variable.x = #{42}; # set x\nreturn variable.x;"
	expected := "v.x=42;return v.x;"
	assertMolang(t, input, expected)
}

func TestMolangCommentInsideStringNotStripped(t *testing.T) {
	// '# ' inside a single-quoted string literal must not be removed
	input := `variable.s = 'hello # world'; return variable.s;`
	expected := `v.s='hello # world';return v.s;`
	assertMolang(t, input, expected)
}

// Valid Molang is minified by molang-go, which also drops redundant parentheses
func TestMolangMinifyRedundantParentheses(t *testing.T) {
	input := `variable.x = (query.life_time > 0.5) ? (1 + 2) * 3 : 0;`
	expected := `v.x=q.life_time>0.5?(1+2)*3:0;`
	assertMolang(t, input, expected)
}

// A fragment that is not valid Molang (e.g. with a placeholder replaced later) fails by default
func TestMolangInvalidFailsByDefault(t *testing.T) {
	_, err := jsonte.ProcessMolangFile(`variable.i = __RECT_INDEX__;`, types.NewJsonObject())
	var invalid *jsonte.InvalidMolangError
	if !errors.As(err, &invalid) {
		t.Fatalf("Expected InvalidMolangError, got %v", err)
	}
}

// With validation disabled, invalid Molang falls back to token minification without an error
func TestMolangPlaceholderFallback(t *testing.T) {
	input := `variable.i = __RECT_INDEX__; variable.x = (1.0);`
	expected := `v.i=__RECT_INDEX__;v.x=(1.0);`
	output, err := jsonte.ProcessMolangFileWithValidation(input, types.NewJsonObject(), false)
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, output)
	}
}

// Valid Molang is still minified by molang-go when validation is disabled
func TestMolangValidWithValidationDisabled(t *testing.T) {
	output, err := jsonte.ProcessMolangFileWithValidation(`variable.x = (1.0);`, types.NewJsonObject(), false)
	if err != nil {
		t.Fatal(err)
	}
	if output != `v.x=1;` {
		t.Fatalf("Expected v.x=1; got %s", output)
	}
}

// Molang the game rejects at load time is an error
func TestMolangErrorForEngineRule(t *testing.T) {
	_, err := jsonte.ProcessMolangFile(`variable.x = query.y ?? 1;`, types.NewJsonObject())
	var invalid *jsonte.InvalidMolangError
	if !errors.As(err, &invalid) {
		t.Fatalf("Expected InvalidMolangError for ?? with a query on the left-hand side, got %v", err)
	}
}

// The error points at the line and column instead of embedding the whole source
func TestMolangErrorPosition(t *testing.T) {
	input := "variable.a = 1;\nvariable.i = __RECT_INDEX__;"
	_, err := jsonte.ProcessMolangFile(input, types.NewJsonObject())
	var invalid *jsonte.InvalidMolangError
	if !errors.As(err, &invalid) {
		t.Fatalf("Expected InvalidMolangError, got %v", err)
	}
	if invalid.Line != 2 || invalid.Column != 14 || strings.Contains(invalid.Error(), "variable.a") {
		t.Fatalf("Unexpected error: %s", invalid)
	}
}
