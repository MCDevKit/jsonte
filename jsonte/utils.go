package jsonte

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
)

const WrongTypeErrTag = "wrong_type"

// FindAnyCase returns the first key that matches the given key, ignoring case.
func FindAnyCase[T types.JsonType](o *types.JsonObject, key ...string) (*T, error) {
	if len(key) == 0 {
		return nil, burrito.WrappedError("The key is nil or empty")
	}
	var candidates [4]string
	candidateCount := 0
	addCandidate := func(candidate string) {
		if candidate == "" {
			return
		}
		for i := 0; i < candidateCount; i++ {
			if candidates[i] == candidate {
				return
			}
		}
		if candidateCount < len(candidates) {
			candidates[candidateCount] = candidate
			candidateCount++
		}
	}

	if len(key) == 1 {
		addCandidate(key[0])
	} else {
		addCandidate(strings.Join(key, "_"))
		addCandidate(strings.Join(key, ""))
	}
	addCandidate(toCamelCase(key...))

	for i := 0; i < candidateCount; i++ {
		candidate := candidates[i]
		if o.ContainsKey(candidate) {
			return typedResult[T](candidate, o.Get(candidate))
		}
	}

	var lowerCandidates [4]string
	lowerCount := 0
	for i := 0; i < candidateCount; i++ {
		lower := strings.ToLower(candidates[i])
		duplicate := false
		for j := 0; j < lowerCount; j++ {
			if lowerCandidates[j] == lower {
				duplicate = true
				break
			}
		}
		if !duplicate && lowerCount < len(lowerCandidates) {
			lowerCandidates[lowerCount] = lower
			lowerCount++
		}
	}

	for _, existing := range o.Keys() {
		lowerExisting := strings.ToLower(existing)
		for i := 0; i < lowerCount; i++ {
			if lowerExisting == lowerCandidates[i] {
				return typedResult[T](existing, o.Get(existing))
			}
		}
	}

	return nil, burrito.WrappedErrorf("The key %s is not found", key)
}

func typedResult[T types.JsonType](key string, value types.JsonType) (*T, error) {
	if typed, ok := value.(T); ok {
		return &typed, nil
	}
	expected := *new(T)
	err := burrito.WrappedErrorf("The value of %s is not a %s", key, types.TypeName(expected))
	burrito.AsBurritoError(err).AddTag(WrongTypeErrTag)
	return nil, err
}

func toCamelCase(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	var sb strings.Builder
	for _, part := range s {
		if part == "" {
			continue
		}
		r, size := utf8.DecodeRuneInString(part)
		if r == utf8.RuneError && size == 0 {
			continue
		}
		sb.WriteRune(unicode.ToUpper(r))
		for _, rr := range part[size:] {
			sb.WriteRune(unicode.ToLower(rr))
		}
	}
	return sb.String()
}

func UnescapeString(text string) string {
	runes := []rune(text)
	var sb strings.Builder
	i := 0
	UnescapeStringToBuffer(runes, &sb, &i, 0)
	return sb.String()
}

// UnescapeStringToBuffer unescapes a string to a buffer. If end is not 0, the unescaping will stop when the end rune is found.
// Returns true if the end rune is found and false otherwise.
// Also returns the escaped string for debugging purposes.
func UnescapeStringToBuffer(text []rune, sb *strings.Builder, i *int, end rune) (bool, string) {
	var debugBuilder strings.Builder
	escape := false
	for ; *i < len(text); *i++ {
		c := text[*i]
		if escape {
			debugBuilder.WriteRune(c)
			escape = false
			switch c {
			case 'r':
				sb.WriteRune('\r')
			case 'n':
				sb.WriteRune('\n')
			case 't':
				sb.WriteRune('\t')
			case 'b':
				sb.WriteRune('\b')
			case 'f':
				sb.WriteRune('\f')
			case 'v':
				sb.WriteRune('\v')
			case '\\':
				sb.WriteRune('\\')
			case '\'':
				sb.WriteRune('\'')
			case '"':
				sb.WriteRune('"')
			case 'u':
				if *i+4 >= len(text) {
					sb.WriteRune(c)
					continue
				}
				r, err := strconv.ParseInt(string(text[*i+1:*i+5]), 16, 32)
				if err != nil {
					utils.Logger.Warnf("Failed to parse unicode escape sequence: %s", err)
					sb.WriteRune(c)
					continue
				}
				*i += 4
				sb.WriteRune(rune(r))
			case end:
				sb.WriteRune(c)
			default:
				utils.Logger.Warnf("Unknown escape sequence: \\%c", c)
				sb.WriteRune(c)
			}
		} else if c == '\\' {
			debugBuilder.WriteRune(c)
			escape = true
			continue
		} else if c == end {
			return true, debugBuilder.String()
		} else {
			debugBuilder.WriteRune(c)
			sb.WriteRune(c)
		}
	}
	return false, debugBuilder.String()
}

var reservedNames = []string{
	"null",
	"true",
	"false",
	"undefined",
	"NaN",
	"if",
	"else",
	"for",
	"while",
	"do",
}

//const variableNamePattern = "^[a-zA-Z_$][a-zA-Z0-9_$]*$"

func VerifyReservedNames(o *types.JsonObject, path string) error {
	for _, key := range o.Value.Keys() {
		err := verifyReservedName(key, path+"."+key)
		if err != nil {
			return err
		}
		if v, ok := o.Value.Get(key).(*types.JsonObject); ok {
			err := VerifyReservedNames(v, path+"."+key)
			if err != nil {
				return err
			}
		} else if v, ok := o.Value.Get(key).(*types.JsonArray); ok {
			err := verifyReservedNamesArray(v, path+"."+key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func verifyReservedNamesArray(o *types.JsonArray, path string) error {
	for i, v := range o.Value {
		p := path + "[" + strconv.Itoa(i) + "]"
		if v, ok := v.(*types.JsonObject); ok {
			err := VerifyReservedNames(v, p)
			if err != nil {
				return err
			}
			continue
		}
		if a, ok := v.(*types.JsonArray); ok {
			err := verifyReservedNamesArray(a, p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func verifyReservedName(key string, path string) error {
	for _, name := range reservedNames {
		if key == name {
			// For now only warn about reserved names
			utils.Logger.Warnf("The key %s at %s is a reserved name. In the future versions, this will fail the compilation.", key, path)
			//return utils.WrappedJsonErrorf(path, "The key %s is a reserved name", key)
		}
	}
	return nil
}
