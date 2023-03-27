package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"testing"
)

func assertLang(t *testing.T, template, expected string) {
	process, err := jsonte.ProcessLangFile(template, types.NewJsonObject())
	if err != nil {
		t.Fatal(err)
	}
	if process != expected {
		t.Fatalf("Expected: \n%s\ngot: \n%s", expected, process)
	}
}

func TestNoOperationLang(t *testing.T) {
	expected := `test.value=Hello World`
	assertLang(t, expected, expected)
}

func TestSimpleLang(t *testing.T) {
	f := `test.value=##{'Hello World'}`
	expected := `test.value=Hello World`
	assertLang(t, f, expected)
}

func TestSimpleLang2(t *testing.T) {
	f := `test.value=##{"Hello World"}`
	expected := `test.value=Hello World`
	assertLang(t, f, expected)
}

func TestLangParsingStrings(t *testing.T) {
	f := `test.value=##{'}'}`
	expected := `test.value=}`
	assertLang(t, f, expected)
}

func TestLangParsingStrings2(t *testing.T) {
	f := `test.value=##{"'"}`
	expected := `test.value='`
	assertLang(t, f, expected)
}

func TestLangParsingStrings3(t *testing.T) {
	f := `##{"test.value=1\ntest.value2=2"}`
	expected := "test.value=1\ntest.value2=2"
	assertLang(t, f, expected)
}

func TestLangNormalComment(t *testing.T) {
	f := `## Just a comment\n##Another comment\ntest.value=1\ntest.value2=2`
	assertLang(t, f, f)
}
