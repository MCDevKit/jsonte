package types

import (
	"strconv"
	"unicode/utf8"
)

// indentSpaces is sliced to produce indentation without allocating.
const indentSpaces = "                                                                                                                                "

func appendIndent(dst []byte, indent int) []byte {
	n := indent * 2
	for n > len(indentSpaces) {
		dst = append(dst, indentSpaces...)
		n -= len(indentSpaces)
	}
	return append(dst, indentSpaces[:n]...)
}

// appendEscapedString appends s escaped exactly like json.MarshalJSONC's writeString
// (without surrounding quotes). Invalid UTF-8 sequences become U+FFFD, matching the
// []rune conversion used by the original serializer.
func appendEscapedString(dst []byte, s string) []byte {
	if !utf8.ValidString(s) {
		s = string([]rune(s))
	}
	start := 0
	for i := 0; i < len(s); i++ {
		var esc byte
		switch s[i] {
		case '\\':
			esc = '\\'
		case '"':
			esc = '"'
		case '\n':
			esc = 'n'
		case '\r':
			esc = 'r'
		case '\t':
			esc = 't'
		case '\f':
			esc = 'f'
		case '\b':
			esc = 'b'
		default:
			continue
		}
		dst = append(dst, s[start:i]...)
		dst = append(dst, '\\', esc)
		start = i + 1
	}
	return append(dst, s[start:]...)
}

func appendQuotedString(dst []byte, s string) []byte {
	dst = append(dst, '"')
	dst = appendEscapedString(dst, s)
	return append(dst, '"')
}

// appendJSON serializes v into dst, producing output byte-identical to
// json.MarshalJSONC(v.Unbox(), pretty) for nested values. It returns ok=false
// when it encounters a value it cannot serialize, in which case the caller must
// fall back to the Unbox-based path.
func appendJSON(dst []byte, v JsonType, pretty bool, indent int) ([]byte, bool) {
	switch t := v.(type) {
	case *JsonObject:
		dst = append(dst, '{')
		if pretty {
			dst = append(dst, '\n')
		}
		indent++
		ok := true
		first := true
		t.RangeEntries(func(key string, value JsonType) bool {
			if !first {
				dst = append(dst, ',')
				if pretty {
					dst = append(dst, '\n')
				}
			}
			first = false
			if pretty {
				dst = appendIndent(dst, indent)
			}
			dst = appendQuotedString(dst, key)
			dst = append(dst, ':')
			if pretty {
				dst = append(dst, ' ')
			}
			dst, ok = appendJSON(dst, value, pretty, indent)
			return !ok
		})
		if !ok {
			return dst, false
		}
		if pretty && !first {
			dst = append(dst, '\n')
		}
		indent--
		if pretty {
			dst = appendIndent(dst, indent)
		}
		return append(dst, '}'), true
	case *JsonArray:
		dst = append(dst, '[')
		if pretty {
			dst = append(dst, '\n')
		}
		indent++
		var ok bool
		for i, value := range t.Value {
			if i > 0 {
				dst = append(dst, ',')
				if pretty {
					dst = append(dst, '\n')
				}
			}
			if pretty {
				dst = appendIndent(dst, indent)
			}
			dst, ok = appendJSON(dst, value, pretty, indent)
			if !ok {
				return dst, false
			}
		}
		if pretty && len(t.Value) > 0 {
			dst = append(dst, '\n')
		}
		indent--
		if pretty {
			dst = appendIndent(dst, indent)
		}
		return append(dst, ']'), true
	case *JsonString:
		return appendQuotedString(dst, t.Value), true
	case *JsonNumber:
		if t.Decimal {
			f := t.FloatValue()
			// If the float doesn't have a decimal point, force `.0` to be appended
			if f == float64(int64(f)) {
				return strconv.AppendFloat(dst, f, 'f', 1, 64), true
			}
			return strconv.AppendFloat(dst, f, 'f', -1, 64), true
		}
		return strconv.AppendInt(dst, int64(t.IntValue()), 10), true
	case *JsonBool:
		return strconv.AppendBool(dst, t.Value), true
	case *JsonNull:
		return append(dst, "null"...), true
	case *JsonSignal:
		return appendJSON(dst, t.Value, pretty, indent)
	case *JsonLambda, *JsonPath, *Semver:
		if s, ok := v.Unbox().(string); ok {
			return appendQuotedString(dst, s), true
		}
		return dst, false
	case nil:
		return append(dst, "null"...), true
	}
	return dst, false
}
