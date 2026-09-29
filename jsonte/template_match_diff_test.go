package jsonte

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/utils"
	"go.uber.org/zap/zapcore"
)

// TestFindTemplateMatchesBytesMatchesRunes checks that the byte fast path returns exactly what the
// rune implementation returns, including errors, on random inputs full of delimiters.
func TestFindTemplateMatchesBytesMatchesRunes(t *testing.T) {
	utils.InitLogging(zapcore.ErrorLevel)
	pieces := []string{"{", "}", "#", "##", "\"", "'", "\\", "é", "日", "\n", "a", "b", " ", "{{", "}}", "#{", "##{", "\\n", "\xff"}
	tokens := [][2]string{{"#", ""}, {"##", ""}, {"{", "}"}}
	rng := rand.New(rand.NewSource(1))
	fast := 0
	defer func() { t.Logf("fast path used for %d inputs", fast) }()
	for n := 0; n < 300000; n++ {
		var sb strings.Builder
		length := rng.Intn(24)
		for k := 0; k < length; k++ {
			p := pieces[rng.Intn(len(pieces))]
			if p == "\xff" && rng.Intn(4) != 0 {
				continue
			}
			sb.WriteString(p)
		}
		input := sb.String()
		for _, tok := range tokens {
			want, wantErr := findTemplateMatchesRunes(input, tok[0], tok[1], false)
			got, ok := findTemplateMatchesBytes(input, tok[0], tok[1])
			if !strings.Contains(input, tok[0]+"{") {
				continue
			}
			if !ok {
				continue
			}
			fast++
			if wantErr != nil {
				t.Fatalf("input %q tokens %q: fast path succeeded but rune implementation failed: %v", input, tok, wantErr)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("input %q tokens %q:\n got  %#v\n want %#v", input, tok, got, want)
			}
		}
	}
}

// TestFindTemplateMatchesFallback checks FindTemplateMatches end to end, errors included.
func TestFindTemplateMatchesFallback(t *testing.T) {
	utils.InitLogging(zapcore.ErrorLevel)
	cases := []struct{ input, start, end string }{
		{`#{'a\'b'} x`, "#", ""},
		{`#{'abc`, "#", ""},
		{`#{a`, "#", ""},
		{`{{a}}`, "{", "}"},
		{`{{a}`, "{", "}"},
		{`{{a}x b`, "{", "}"},
		{`é#{"日"}é #{1}`, "#", ""},
	}
	for _, c := range cases {
		want, wantErr := findTemplateMatchesRunes(c.input, c.start, c.end, false)
		got, gotErr := FindTemplateMatches(c.input, c.start, c.end)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("input %q: got %#v, want %#v", c.input, got, want)
		}
		if (gotErr == nil) != (wantErr == nil) || (gotErr != nil && gotErr.Error() != wantErr.Error()) {
			t.Errorf("input %q: got error %v, want %v", c.input, gotErr, wantErr)
		}
	}
}
