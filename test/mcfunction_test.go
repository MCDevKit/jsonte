package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func assertMCFunction(t *testing.T, template, expected string) {
	t.Helper()
	process, err := jsonte.ProcessMCFunction(template, types.NewJsonObject())
	if err != nil {
		t.Fatal(err)
	}
	if process != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, process)
	}
}

func TestNoOperationMCFunction(t *testing.T) {
	expected := `/give @s diamond_block 1 0 {"minecraft:can_place_on":{"blocks":["dirt"]},"minecraft:can_destroy":{"blocks":["quartz_block"]}}
say "Hello World"`
	assertMCFunction(t, expected, expected)
}

func TestSimpleMCFunction(t *testing.T) {
	f := `/give @s diamond_block 1 0 {"minecraft:can_place_on":{"blocks":#{["dirt", "stone"]}},"minecraft:can_destroy":{"blocks":["quartz_block"]}}
say "Hello World"`
	expected := `/give @s diamond_block 1 0 {"minecraft:can_place_on":{"blocks":["dirt","stone"]},"minecraft:can_destroy":{"blocks":["quartz_block"]}}
say "Hello World"`
	assertMCFunction(t, f, expected)
}

func TestSimpleMCFunction2(t *testing.T) {
	f := `/give @s diamond_block 1 0 {"minecraft:can_place_on":{"blocks":#{["dirt", "stone"]}},"minecraft:can_destroy":{"blocks":["quartz_block"]}}
say "#{'Hello World'}"`
	expected := `/give @s diamond_block 1 0 {"minecraft:can_place_on":{"blocks":["dirt","stone"]},"minecraft:can_destroy":{"blocks":["quartz_block"]}}
say "Hello World"`
	assertMCFunction(t, f, expected)
}

func TestParsingStrings(t *testing.T) {
	f := `say "#{'}'}"`
	expected := `say "}"`
	assertMCFunction(t, f, expected)
}

func TestParsingStrings2(t *testing.T) {
	f := `say "#{"'"}"`
	expected := `say "'"`
	assertMCFunction(t, f, expected)
}

func TestParsingStrings3(t *testing.T) {
	f := `#{"say test\nsay test2"}`
	expected := "say test\nsay test2"
	assertMCFunction(t, f, expected)
}

func TestMultilineExpression(t *testing.T) {
	f := `#{
	(1..10)
		.map(x => x * 2)
		.filter(x => mod(x, 3) == 0)
		.join(", ")
}`
	expected := "6, 12, 18"
	assertMCFunction(t, f, expected)
}

func TestMCFunctionNormalComment(t *testing.T) {
	f := `# Just a comment\n#Another comment\nsay hello`
	assertMCFunction(t, f, f)
}
