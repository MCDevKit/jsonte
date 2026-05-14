package utils

import "testing"

func TestMinify(t *testing.T) {
	input := `
		# comment
		variable.hand_bob = query.life_time < 0.01 ? 0.0 : variable.hand_bob + 1;
		return variable.hand_bob;
	`

	got := Minify(input)
	want := "variable.hand_bob=query.life_time<0.01?0.0:variable.hand_bob+1;return variable.hand_bob;"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestKeepsRequiredSpace(t *testing.T) {
	got := Minify("return query.life_time")
	want := "return query.life_time"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestStringAndHash(t *testing.T) {
	got := Minify("variable.x = 'a # not comment'; # comment\nreturn variable.x;")
	want := "variable.x='a # not comment';return variable.x;"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestMinifyWithShortAccessors(t *testing.T) {
	got := MinifyWithShortAccessors("variable.x = 'variable.x and query.x';\nreturn variable.x;")
	want := "v.x='variable.x and query.x';return v.x;"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
