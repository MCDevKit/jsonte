package json

import (
	"encoding/json"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/utils"
)

func TestUnmarshallJSONCErrors(t *testing.T) {
	cases := map[string]string{
		`{"a": {"b": [1, {"c": [0, 1, tre]}]}}`: "Unexpected token 'e' (U+0065), expected 'u' (U+0075) at line 1, column 32 at #/a/b[1]/c[2]",
		`{"a": {"b": [1, {"c": [0, 1, "x`:       "Unexpected end of file at line 1, column 32 at #/a/b[1]/c[2]",
		`{"é": ["ø", {"x": nulx}]}`:             "Unexpected token 'x' (U+0078), expected 'l' (U+006C) at line 1, column 22 at #/é[1]/x",
		"{\"a\":\n  \"l1\nl2 ü\" x}":            "Unexpected token 'x' (U+0078) at line 3, column 6 at #",
		`[1,]`:                                  "Unexpected token ']' (U+005D) at line 1, column 3 at #[1]",
		`{"a": 1} x`:                            "Expected end of file at line 1, column 9 at #",
	}
	for in, expected := range cases {
		_, err := UnmarshallJSONC([]byte(in))
		if err == nil {
			t.Fatalf("expected an error for %q", in)
		}
		if err.Error() != expected {
			t.Errorf("unexpected error for %q:\n got: %s\nwant: %s", in, err.Error(), expected)
		}
	}
}

func TestUnmarshallJSONCValues(t *testing.T) {
	v, err := UnmarshallJSONC([]byte(`{"s": ["x\ny\u00e9\u0041\q\/\"\\ żółć 😀", "plain"], "n": [-1.5e3, 7], "b": [true, false, null], "e": {}}`))
	if err != nil {
		t.Fatal(err)
	}
	m := v.(utils.NavigableMap[string, interface{}])
	s := m.Get("s").([]interface{})
	if s[0] != "x\nyéA\\q/\"\\ żółć 😀" || s[1] != "plain" {
		t.Errorf("unexpected strings: %q", s)
	}
	n := m.Get("n").([]interface{})
	if n[0] != interface{}(json.Number("-1.5e3")) || n[1] != interface{}(json.Number("7")) {
		t.Errorf("unexpected numbers: %v", n)
	}
	b := m.Get("b").([]interface{})
	if b[0] != true || b[1] != false || b[2] != nil {
		t.Errorf("unexpected constants: %v", b)
	}
	if e := m.Get("e").(utils.NavigableMap[string, interface{}]); e.Size() != 0 {
		t.Errorf("expected an empty object, got %v", e.Keys())
	}
}
