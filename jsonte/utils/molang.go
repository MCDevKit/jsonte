package utils

import (
	"strings"
	"unicode"
)

type tokenKind int

const (
	tokenUnknown tokenKind = iota
	tokenName
	tokenNumber
	tokenString
	tokenPunct
)

type token struct {
	kind tokenKind
	text string
}

func Minify(src string) string {
	tokens := tokenize(src)
	if len(tokens) == 0 {
		return ""
	}

	var b strings.Builder
	prev := tokens[0]
	b.WriteString(prev.text)

	for _, cur := range tokens[1:] {
		if needsSpace(prev, cur) {
			b.WriteByte(' ')
		}
		b.WriteString(cur.text)
		prev = cur
	}

	return b.String()
}

func MinifyWithShortAccessors(src string) string {
	tokens := tokenize(src)

	for i := range tokens {
		if tokens[i].kind != tokenName {
			continue
		}

		switch {
		case strings.HasPrefix(tokens[i].text, "query."):
			tokens[i].text = "q." + strings.TrimPrefix(tokens[i].text, "query.")
		case strings.HasPrefix(tokens[i].text, "variable."):
			tokens[i].text = "v." + strings.TrimPrefix(tokens[i].text, "variable.")
		case strings.HasPrefix(tokens[i].text, "context."):
			tokens[i].text = "c." + strings.TrimPrefix(tokens[i].text, "context.")
		case strings.HasPrefix(tokens[i].text, "temp."):
			tokens[i].text = "t." + strings.TrimPrefix(tokens[i].text, "temp.")
		}
	}

	if len(tokens) == 0 {
		return ""
	}

	var b strings.Builder
	prev := tokens[0]
	b.WriteString(prev.text)

	for _, cur := range tokens[1:] {
		if needsSpace(prev, cur) {
			b.WriteByte(' ')
		}
		b.WriteString(cur.text)
		prev = cur
	}

	return b.String()
}

func tokenize(src string) []token {
	out := make([]token, 0, len(src)/2)

	for i := 0; i < len(src); {
		c := src[i]

		if isWhitespace(c) {
			i++
			continue
		}

		// Molang comments in bridge tokenizer are "# until newline".
		if c == '#' {
			i++
			for i < len(src) && src[i] != '\n' {
				i++
			}
			continue
		}

		if isNameStart(c) {
			start := i
			i++
			for i < len(src) && isNamePart(src[i]) {
				i++
			}
			out = append(out, token{kind: tokenName, text: src[start:i]})
			continue
		}

		if isDigit(c) || c == '.' {
			start := i
			hasDecimal := c == '.'
			i++

			for i < len(src) {
				if isDigit(src[i]) {
					i++
					continue
				}

				if src[i] == '.' && !hasDecimal {
					hasDecimal = true
					i++
					continue
				}

				break
			}

			// Match bridge behavior: "0.5f" is accepted as float notation,
			// but only when there is a decimal part.
			if hasDecimal && i < len(src) && src[i] == 'f' {
				i++
			}

			out = append(out, token{kind: tokenNumber, text: src[start:i]})
			continue
		}

		if c == '\'' {
			start := i
			i++

			for i < len(src) && src[i] != '\'' {
				i++
			}

			if i < len(src) && src[i] == '\'' {
				i++
			}

			out = append(out, token{kind: tokenString, text: src[start:i]})
			continue
		}

		out = append(out, token{kind: tokenPunct, text: src[i : i+1]})
		i++
	}

	return out
}

func needsSpace(a, b token) bool {
	// Identifiers/keywords must not merge:
	// "return query.x" cannot become "returnquery.x".
	if isWordLike(a) && isWordLike(b) {
		return true
	}

	// Number followed by name can become a different/invalid token:
	// "1 query.x" -> "1query.x".
	if a.kind == tokenNumber && b.kind == tokenName {
		return true
	}

	// Name followed by number can also merge:
	// "v.x 1" -> "v.x1".
	if a.kind == tokenName && b.kind == tokenNumber {
		return true
	}

	// Dot is ambiguous because bridge tokenizer treats "." as part of both
	// names and numbers. Preserve separation for malformed/edge input.
	if a.text == "." || b.text == "." {
		return true
	}

	return false
}

func isWordLike(t token) bool {
	return t.kind == tokenName || t.kind == tokenNumber
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isNameStart(c byte) bool {
	return c == '_' || isASCIIAlpha(c)
}

func isNamePart(c byte) bool {
	return c == '_' || c == '.' || isASCIIAlpha(c) || isDigit(c)
}

func isASCIIAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Optional helper if you want case-insensitive comparison elsewhere.
// Do not use it inside Minify unless you intentionally want to normalize identifiers.
func lowerASCII(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if r <= unicode.MaxASCII {
			b.WriteByte(byte(unicode.ToLower(r)))
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
