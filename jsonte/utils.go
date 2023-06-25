package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
)

const WrongTypeErrTag = "wrong_type"

// FindAnyCase returns the first key that matches the given key, ignoring case.
func FindAnyCase[T types.JsonType](o types.JsonObject, key ...string) (*T, error) {
	if key == nil || len(key) == 0 {
		return nil, burrito.WrappedError("The key is nil or empty")
	}
	if len(key) > 1 {
		snakeCase := strings.Join(key, "_")
		for _, k := range o.Keys() {
			if strings.EqualFold(k, snakeCase) {
				t, ok := o.Get(k).(T)
				if !ok {
					err := burrito.WrappedErrorf("The value of %s is not a %s", k, types.TypeName(*new(T)))
					burrito.AsBurritoError(err).AddTag(WrongTypeErrTag)
					return nil, err
				}
				return &t, nil
			}
		}
	}
	camelCase := toCamelCase(key...)
	for _, k := range o.Keys() {
		if strings.EqualFold(k, camelCase) {
			t, ok := o.Get(k).(T)
			if !ok {
				err := burrito.WrappedErrorf("The value of %s is not a %s", k, types.TypeName(*new(T)))
				burrito.AsBurritoError(err).AddTag(WrongTypeErrTag)
				return nil, err
			}
			return &t, nil
		}
	}
	return nil, burrito.WrappedErrorf("The key %s is not found", key)
}

func toCamelCase(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	var sb strings.Builder
	for _, v := range s {
		sb.WriteString(cases.Title(language.Und).String(v))
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
