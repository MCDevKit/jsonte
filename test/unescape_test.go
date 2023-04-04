package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"strings"
	"testing"
)

func TestNotEscapedString(t *testing.T) {
	expected := "Hello World!"
	actual := jsonte.UnescapeString(`Hello World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestNewLineString(t *testing.T) {
	expected := "Hello\nWorld!"
	actual := jsonte.UnescapeString(`Hello\nWorld!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestCarriageReturnString(t *testing.T) {
	expected := "Hello\rWorld!"
	actual := jsonte.UnescapeString(`Hello\rWorld!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestTabString(t *testing.T) {
	expected := "Hello\tWorld!"
	actual := jsonte.UnescapeString(`Hello\tWorld!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestBackslashString(t *testing.T) {
	expected := "Hello\\World!"
	actual := jsonte.UnescapeString(`Hello\\World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestDoubleQuoteString(t *testing.T) {
	expected := "Hello\"World!"
	actual := jsonte.UnescapeString(`Hello\"World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestSingleQuoteString(t *testing.T) {
	expected := "Hello'World!"
	actual := jsonte.UnescapeString(`Hello\'World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestUnicodeString(t *testing.T) {
	expected := "Hello\u0000World!"
	actual := jsonte.UnescapeString(`Hello\u0000World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestUnicodeString2(t *testing.T) {
	expected := "HelloAWorld!"
	actual := jsonte.UnescapeString(`Hello\u0041World!`)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}

func TestCustomEndString(t *testing.T) {
	expected := "Hello'World!"
	expectedEscaped := "Hello\\'World!"
	var sb strings.Builder
	i := 0
	ended, escaped := jsonte.UnescapeStringToBuffer(`Hello\'World!'asd`, &sb, &i, '\'')
	actual := sb.String()
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
	if !ended {
		t.Fatalf("Expected string to end")
	}
	if escaped != expectedEscaped {
		t.Fatalf("Expected escaped: %s, got: %s", expectedEscaped, escaped)
	}
}

func TestEmptyString(t *testing.T) {
	expected := ""
	actual := jsonte.UnescapeString(``)
	if actual != expected {
		t.Fatalf("Expected: %s, got: %s", expected, actual)
	}
}
