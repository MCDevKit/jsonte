package json

import (
	"encoding/json"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"strconv"
	"unicode/utf8"
)

const UnexpectedTokenExpectedError = "Unexpected token '%c' (%[1]U), expected '%c' (%[2]U) at line %d, column %d"
const UnexpectedTokenError = "Unexpected token '%c' (%[1]U) at line %d, column %d"
const UnexpectedEofError = "Unexpected end of file at line %d, column %d"
const ExpectedTokenError = "Expected '%c' at line %d, column %d"
const ExpectedEofError = "Expected end of file at line %d, column %d"

const (
	TokenOpenObject     = '{'
	TokenCloseObject    = '}'
	TokenOpenArray      = '['
	TokenCloseArray     = ']'
	TokenComma          = ','
	TokenColon          = ':'
	TokenDoubleQuote    = '"'
	TokenBackslash      = '\\'
	TokenSlash          = '/'
	TokenAsterisk       = '*'
	TokenNewline        = '\n'
	TokenSpace          = ' '
	TokenCarriageReturn = '\r'
	TokenTab            = '\t'
	TokenFormFeed       = '\f'
	TokenBackspace      = '\b'
	TokenEof            = 4
)

// StringReader reads runes from a UTF-8 encoded byte slice. The marker is a
// byte offset, while the column is counted in runes.
type StringReader struct {
	str    []byte
	marker int
	column int
	line   int
}

func NewStringReader(str []byte) *StringReader {
	if !utf8.Valid(str) {
		// Replace every invalid byte with U+FFFD, exactly like a []rune
		// conversion does, so that the reader only ever sees valid UTF-8.
		str = []byte(string([]rune(string(str))))
	}
	return &StringReader{
		str:    str,
		marker: 0,
		column: 0,
		line:   1,
	}
}

func (sr *StringReader) Peek() rune {
	if sr.marker >= len(sr.str) {
		return TokenEof
	}
	if c := sr.str[sr.marker]; c < utf8.RuneSelf {
		return rune(c)
	}
	r, _ := utf8.DecodeRune(sr.str[sr.marker:])
	return r
}

func (sr *StringReader) Read() rune {
	var c rune
	size := 1
	if sr.marker >= len(sr.str) {
		c = TokenEof
	} else if b := sr.str[sr.marker]; b < utf8.RuneSelf {
		c = rune(b)
	} else {
		c, size = utf8.DecodeRune(sr.str[sr.marker:])
	}
	if c == TokenNewline {
		sr.line++
		sr.column = 0
	} else {
		sr.column++
	}
	sr.marker += size
	return c
}

// ValueBuilder creates the values produced by the JSONC parser.
// The slices passed to Array and Object are scratch buffers owned by the
// parser and are reused after the call returns, so implementations must copy
// their contents.
type ValueBuilder interface {
	String(s string) interface{}
	Number(n json.Number) interface{}
	Bool(b bool) interface{}
	Null() interface{}
	Array(values []interface{}) interface{}
	Object(keys []string, values []interface{}) interface{}
}

// defaultBuilder produces NavigableMap[string, interface{}], []interface{},
// string, json.Number, bool and nil values.
type defaultBuilder struct{}

func (defaultBuilder) String(s string) interface{}      { return s }
func (defaultBuilder) Number(n json.Number) interface{} { return n }
func (defaultBuilder) Bool(b bool) interface{}          { return b }
func (defaultBuilder) Null() interface{}                { return nil }

func (defaultBuilder) Array(values []interface{}) interface{} {
	result := make([]interface{}, len(values))
	copy(result, values)
	return result
}

func (defaultBuilder) Object(keys []string, values []interface{}) interface{} {
	result := utils.NewNavigableMapWithCapacity[string, interface{}](len(keys))
	for i, key := range keys {
		result.Put(key, values[i])
	}
	return result
}

// pathSegment is one step of the JSON path of the value being parsed. The
// textual path is only built when an error is reported.
type pathSegment struct {
	key     string
	index   int
	isIndex bool
}

type parser struct {
	StringReader
	builder ValueBuilder
	keys    []string
	values  []interface{}
	path    []pathSegment
}

// pathString returns the JSON path of the value currently being parsed.
func (str *parser) pathString() string {
	result := "#"
	for _, segment := range str.path {
		if segment.isIndex {
			result = fmt.Sprintf("%s[%d]", result, segment.index)
		} else {
			result = result + "/" + segment.key
		}
	}
	return result
}

func UnmarshallJSONC(str []byte) (interface{}, error) {
	return UnmarshallJSONCWithBuilder(str, defaultBuilder{})
}

// UnmarshallJSONCWithBuilder parses a JSONC document and constructs the values
// with the given builder.
func UnmarshallJSONCWithBuilder(str []byte, builder ValueBuilder) (interface{}, error) {
	str, err := ConvertToUTF8(str)
	if err != nil {
		return nil, burrito.PassError(err)
	}
	reader := &parser{StringReader: *NewStringReader(str), builder: builder}
	err = skipWhitespace(&reader.StringReader)
	if err != nil {
		return nil, err
	}
	object, err := reader.parseValue()
	if err != nil {
		return nil, err
	}
	err = skipWhitespace(&reader.StringReader)
	if err != nil {
		return nil, err
	}
	if reader.Peek() != TokenEof {
		return object, utils.WrappedJsonErrorf("#", ExpectedEofError, reader.line, reader.column)
	}
	return object, nil
}

// parseValue parses the value starting at the current token.
func (str *parser) parseValue() (interface{}, error) {
	switch str.Peek() {
	case TokenOpenObject:
		return str.parseObject()
	case TokenOpenArray:
		return str.parseArray()
	case TokenDoubleQuote:
		value, err := str.parseString()
		if err != nil {
			return nil, err
		}
		return str.builder.String(value), nil
	default:
		return str.parsePrimitive()
	}
}

func MarshalJSONC(object interface{}, pretty bool) ([]rune, error) {
	switch object.(type) {
	case utils.NavigableMap[string, interface{}]:
		return writeObject(object.(utils.NavigableMap[string, interface{}]), pretty, 0)
	case map[string]interface{}:
		return writeObject(utils.MapToNavigableMap(object.(map[string]interface{})), pretty, 0)
	case []interface{}:
		return writeArray(object.([]interface{}), pretty, 0)
	case string:
		return writeString(object.(string)), nil
	case float64:
		// If the float doesn't have a decimal point, force `.0` to be appended
		if object.(float64) == float64(int64(object.(float64))) {
			return []rune(strconv.FormatFloat(object.(float64), 'f', 1, 64)), nil
		}
		return []rune(strconv.FormatFloat(object.(float64), 'f', -1, 64)), nil
	case int:
		return []rune(strconv.FormatInt(int64(object.(int)), 10)), nil
	case int64:
		return []rune(strconv.FormatInt(object.(int64), 10)), nil
	case bool:
		if object.(bool) {
			return []rune("true"), nil
		} else {
			return []rune("false"), nil
		}
	case nil:
		return []rune("null"), nil
	default:
		return nil, burrito.WrappedErrorf("Unsupported type %T", object)
	}
}

func writeObject(object utils.NavigableMap[string, interface{}], pretty bool, indent int) ([]rune, error) {
	var result []rune
	result = append(result, TokenOpenObject)
	if pretty {
		result = append(result, TokenNewline)
	}
	indent++
	keys := object.Keys()
	for i, key := range keys {
		if pretty {
			result = append(result, indentBytes(indent)...)
		}
		result = append(result, TokenDoubleQuote)
		result = append(result, writeString(key)...)
		result = append(result, TokenDoubleQuote)
		result = append(result, TokenColon)
		if pretty {
			result = append(result, TokenSpace)
		}
		value := object.Get(key)
		switch value.(type) {
		case utils.NavigableMap[string, interface{}]:
			bytes, err := writeObject(value.(utils.NavigableMap[string, interface{}]), pretty, indent)
			if err != nil {
				return result, err
			}
			result = append(result, bytes...)
		case []interface{}:
			bytes, err := writeArray(value.([]interface{}), pretty, indent)
			if err != nil {
				return result, err
			}
			result = append(result, bytes...)
		case string:
			result = append(result, TokenDoubleQuote)
			result = append(result, writeString(value.(string))...)
			result = append(result, TokenDoubleQuote)
		case float64:
			// If the float doesn't have a decimal point, force `.0` to be appended
			if value.(float64) == float64(int64(value.(float64))) {
				result = append(result, []rune(strconv.FormatFloat(value.(float64), 'f', 1, 64))...)
			} else {
				result = append(result, []rune(strconv.FormatFloat(value.(float64), 'f', -1, 64))...)
			}
		case int:
			result = append(result, []rune(strconv.FormatInt(int64(value.(int)), 10))...)
		case int32:
			result = append(result, []rune(strconv.FormatInt(int64(value.(int32)), 10))...)
		case bool:
			result = append(result, []rune(strconv.FormatBool(value.(bool)))...)
		case nil:
			result = append(result, []rune("null")...)
		default:
			return result, burrito.WrappedErrorf("Unsupported type %T", value)
		}
		if i < len(keys)-1 {
			result = append(result, TokenComma)
		}
		if pretty {
			result = append(result, TokenNewline)
		}
	}
	indent--
	if pretty {
		result = append(result, indentBytes(indent)...)
	}
	result = append(result, TokenCloseObject)
	return result, nil
}

func writeArray(arr []interface{}, pretty bool, indent int) ([]rune, error) {
	var result []rune
	result = append(result, TokenOpenArray)
	if pretty {
		result = append(result, TokenNewline)
	}
	indent++
	for i, value := range arr {
		if pretty {
			result = append(result, indentBytes(indent)...)
		}
		switch value.(type) {
		case utils.NavigableMap[string, interface{}]:
			bytes, err := writeObject(value.(utils.NavigableMap[string, interface{}]), pretty, indent)
			if err != nil {
				return result, err
			}
			result = append(result, bytes...)
		case []interface{}:
			bytes, err := writeArray(value.([]interface{}), pretty, indent)
			if err != nil {
				return result, err
			}
			result = append(result, bytes...)
		case string:
			result = append(result, TokenDoubleQuote)
			result = append(result, writeString(value.(string))...)
			result = append(result, TokenDoubleQuote)
		case float64:
			// If the float doesn't have a decimal point, force `.0` to be appended
			if value.(float64) == float64(int64(value.(float64))) {
				result = append(result, []rune(strconv.FormatFloat(value.(float64), 'f', 1, 64))...)
			} else {
				result = append(result, []rune(strconv.FormatFloat(value.(float64), 'f', -1, 64))...)
			}
		case int:
			result = append(result, []rune(strconv.FormatInt(int64(value.(int)), 10))...)
		case int32:
			result = append(result, []rune(strconv.FormatInt(int64(value.(int32)), 10))...)
		case bool:
			result = append(result, []rune(strconv.FormatBool(value.(bool)))...)
		case nil:
			result = append(result, []rune("null")...)
		default:
			return result, burrito.WrappedErrorf("Unsupported type %T", value)
		}
		if i < len(arr)-1 {
			result = append(result, TokenComma)
		}
		if pretty {
			result = append(result, TokenNewline)
		}
	}
	indent--
	if pretty {
		result = append(result, indentBytes(indent)...)
	}
	result = append(result, TokenCloseArray)
	return result, nil
}

func writeString(s string) []rune {
	var result []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		switch c {
		case TokenBackslash:
			result = append(result, TokenBackslash)
			result = append(result, TokenBackslash)
		case TokenDoubleQuote:
			result = append(result, TokenBackslash)
			result = append(result, TokenDoubleQuote)
		case TokenNewline:
			result = append(result, TokenBackslash)
			result = append(result, 'n')
		case TokenCarriageReturn:
			result = append(result, TokenBackslash)
			result = append(result, 'r')
		case TokenTab:
			result = append(result, TokenBackslash)
			result = append(result, 't')
		case TokenFormFeed:
			result = append(result, TokenBackslash)
			result = append(result, 'f')
		case TokenBackspace:
			result = append(result, TokenBackslash)
			result = append(result, 'b')
		default:
			result = append(result, c)
		}
	}
	return result
}

func indentBytes(indent int) []rune {
	var result []rune
	for i := 0; i < indent*2; i++ {
		result = append(result, TokenSpace)
	}
	return result
}

func isWhitespace(token rune) bool {
	return token == TokenSpace || token == TokenTab || token == TokenNewline || token == TokenCarriageReturn
}

func skipWhitespace(str *StringReader) error {
	canStartComment := false
	for {
		token := str.Peek()
		if isWhitespace(token) {
			str.Read()
			continue
		} else if token == TokenSlash {
			canStartComment = parseComment(str, canStartComment)
			continue
		} else if token == TokenAsterisk {
			canStartComment = parseComment(str, canStartComment)
			continue
		} else if canStartComment && token != TokenSlash && token != TokenAsterisk {
			return burrito.WrappedErrorf(UnexpectedTokenExpectedError, token, TokenSlash, str.line, str.column)
		}
		break
	}
	return nil
}

func (str *parser) parseObject() (interface{}, error) {
	if str.Read() != TokenOpenObject {
		return nil, utils.WrappedJsonErrorf(str.pathString(), ExpectedTokenError, TokenOpenObject, str.line, str.column)
	}

	keyBase := len(str.keys)
	valueBase := len(str.values)
	defer func() {
		str.keys = str.keys[:keyBase]
		str.values = str.values[:valueBase]
	}()

	canStartComment := false
	comma := false
	open := true

	for {
		token := str.Peek()
		if token == TokenEof {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedEofError, str.line, str.column)
		} else if isWhitespace(token) {
			str.Read()
		} else if token == TokenSlash {
			canStartComment = parseComment(&str.StringReader, canStartComment)
		} else if token == TokenAsterisk {
			canStartComment = parseComment(&str.StringReader, canStartComment)
		} else if canStartComment && token != TokenSlash && token != TokenAsterisk {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenExpectedError, token, TokenSlash, str.line, str.column)
		} else if token == TokenDoubleQuote && (comma || open) {
			comma = false
			open = false
			key, err := str.parseString()
			if err != nil {
				return nil, err
			}
			err = skipWhitespace(&str.StringReader)
			if err != nil {
				return nil, err
			}
			if str.Read() != TokenColon {
				return nil, utils.WrappedJsonErrorf(str.pathString(), ExpectedTokenError, TokenColon, str.line, str.column)
			}
			err = skipWhitespace(&str.StringReader)
			if err != nil {
				return nil, err
			}
			str.path = append(str.path, pathSegment{key: key})
			value, err := str.parseValue()
			if err != nil {
				return nil, err
			}
			str.path = str.path[:len(str.path)-1]
			str.keys = append(str.keys, key)
			str.values = append(str.values, value)
		} else if token == TokenComma && !comma && !open {
			comma = true
			str.Read()
		} else if token == TokenCloseObject && !comma {
			str.Read()
			return str.builder.Object(str.keys[keyBase:], str.values[valueBase:]), nil
		} else {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenError, token, str.line, str.column)
		}
	}
}

func parseComment(str *StringReader, canStartComment bool) bool {
	canEndComment := false
	isSingleLineComment := false

	token := str.Read()

	if token == TokenSlash && !canStartComment {
		return true
	} else if token == TokenAsterisk && !canStartComment {
		return true
	} else if token == TokenSlash && canStartComment {
		isSingleLineComment = true
	} else if token == TokenAsterisk && canStartComment {
		isSingleLineComment = false
	} else {
		return false
	}

	for {
		token = str.Read()
		if token == TokenEof {
			return false
		}
		if token == TokenNewline && isSingleLineComment {
			return false
		}
		if token == TokenAsterisk && !isSingleLineComment {
			canEndComment = true
			continue
		}
		if token == TokenSlash && !isSingleLineComment && canEndComment {
			return false
		}
	}
}

func (str *parser) parsePrimitive() (interface{}, error) {
	token := str.Peek()
	if token >= '0' && token <= '9' || token == '-' || token == '+' || token == '.' {
		number, err := parseNumber(&str.StringReader)
		if err != nil {
			return nil, err
		}
		return str.builder.Number(number), nil
	} else if token == 't' {
		if err := str.parseConstant("true"); err != nil {
			return nil, err
		}
		return str.builder.Bool(true), nil
	} else if token == 'f' {
		if err := str.parseConstant("false"); err != nil {
			return nil, err
		}
		return str.builder.Bool(false), nil
	} else if token == 'n' {
		if err := str.parseConstant("null"); err != nil {
			return nil, err
		}
		return str.builder.Null(), nil
	} else {
		return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenError, token, str.line, str.column)
	}
}

func (str *parser) parseConstant(constant string) error {
	for i := 0; i < len(constant); i++ {
		token := str.Read()
		if token != rune(constant[i]) {
			return utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenExpectedError, token, constant[i], str.line, str.column)
		}
	}
	return nil
}

func parseNumber(str *StringReader) (json.Number, error) {
	start := str.marker

	state := 0
	dot := false
	exp := false
	sign := false
	expSign := false

	for {
		c := str.Peek()
		switch state {
		case 0: // Initial state
			if c >= '0' && c <= '9' {
				state = 1
			} else if c == '-' || c == '+' {
				state = 1
				sign = true
				str.Read()
			} else {
				return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
			}
		case 1: // Reading digits before dot
			if c >= '0' && c <= '9' {
				str.Read()
			} else if c == '.' {
				if dot {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				dot = true
				state = 2
				str.Read()
			} else if c == 'e' || c == 'E' {
				if exp {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				exp = true
				state = 3
				str.Read()
			} else if c == '-' || c == '+' {
				return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
			} else {
				if str.marker-start == 1 && sign {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				state = 4
			}
		case 2: // Reading digits after dot
			if c >= '0' && c <= '9' {
				str.Read()
			} else if c == 'e' || c == 'E' {
				if exp {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				exp = true
				state = 3
				str.Read()
			} else if c == '.' || c == '-' || c == '+' {
				return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
			} else {
				state = 4
			}
		case 3: // Reading digits of exponent
			if c >= '0' && c <= '9' {
				str.Read()
			} else if c == '-' || c == '+' {
				if expSign {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				expSign = true
				str.Read()
			} else {
				if str.str[str.marker-1] == 'e' || str.str[str.marker-1] == 'E' || str.str[str.marker-1] == '-' || str.str[str.marker-1] == '+' {
					return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, c, str.line, str.column)
				}
				state = 4
			}
		case 4: // End state
			// All characters of a number are ASCII, so the byte slice matches the runes read.
			return json.Number(str.str[start:str.marker]), nil
		}
	}
}

func (str *parser) parseArray() (interface{}, error) {
	if str.Read() != TokenOpenArray {
		return nil, utils.WrappedJsonErrorf(str.pathString(), ExpectedTokenError, TokenOpenArray, str.line, str.column)
	}

	base := len(str.values)
	defer func() {
		str.values = str.values[:base]
	}()

	comma := false
	open := true
	canStartComment := false

	for {
		token := str.Peek()
		if token == TokenEof {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenExpectedError, token, TokenEof, str.line, str.column)
		} else if isWhitespace(token) {
			str.Read()
		} else if token == TokenSlash {
			canStartComment = parseComment(&str.StringReader, canStartComment)
		} else if token == TokenAsterisk {
			canStartComment = parseComment(&str.StringReader, canStartComment)
		} else if canStartComment && token != TokenSlash && token != TokenAsterisk {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenExpectedError, token, TokenSlash, str.line, str.column)
		} else if (token == TokenOpenArray || token == TokenOpenObject || token == TokenDoubleQuote) && (comma || open) {
			str.path = append(str.path, pathSegment{index: len(str.values) - base, isIndex: true})
			value, err := str.parseValue()
			if err != nil {
				return nil, err
			}
			str.path = str.path[:len(str.path)-1]
			str.values = append(str.values, value)
			comma = false
			open = false
		} else if token == TokenComma && !comma && !open {
			comma = true
			str.Read()
		} else if token == TokenCloseArray && !comma {
			str.Read()
			return str.builder.Array(str.values[base:]), nil
		} else if comma || open {
			str.path = append(str.path, pathSegment{index: len(str.values) - base, isIndex: true})
			value, err := str.parsePrimitive()
			if err != nil {
				return nil, err
			}
			str.path = str.path[:len(str.path)-1]
			str.values = append(str.values, value)
			comma = false
			open = false
		} else {
			return nil, utils.WrappedJsonErrorf(str.pathString(), UnexpectedTokenError, token, str.line, str.column)
		}
	}
}

func (str *parser) parseString() (string, error) {
	str.Read()
	// Fast path: scan for the closing quote. When the string contains no escape
	// sequences, it can be returned as a single slice of the input.
	column := str.column
	line := str.line
	for i := str.marker; i < len(str.str); i++ {
		c := str.str[i]
		if c == TokenDoubleQuote {
			result := string(str.str[str.marker:i])
			str.marker = i + 1
			str.column = column + 1
			str.line = line
			return result, nil
		}
		if c == TokenBackslash || c == TokenEof {
			break
		}
		if c == TokenNewline {
			line++
			column = 0
		} else if !utf8.RuneStart(c) {
			// Continuation byte of a multi-byte rune; the column was counted at its first byte.
		} else {
			column++
		}
	}
	return str.parseStringSlow()
}

func (str *parser) parseStringSlow() (string, error) {
	var result []byte
	for {
		token := str.Read()
		if token == TokenEof {
			return "", utils.WrappedJsonErrorf(str.pathString(), UnexpectedEofError, str.line, str.column)
		}
		if token == TokenDoubleQuote {
			break
		}
		if token == TokenBackslash {
			token = str.Read()
			if token == TokenEof {
				return "", utils.WrappedJsonErrorf(str.pathString(), UnexpectedEofError, str.line, str.column)
			}
			switch token {
			case TokenDoubleQuote:
				result = append(result, '"')
			case TokenBackslash:
				result = append(result, '\\')
			case TokenSlash:
				result = append(result, '/')
			case 'b':
				result = append(result, '\b')
			case 'f':
				result = append(result, '\f')
			case 'n':
				result = append(result, '\n')
			case 'r':
				result = append(result, '\r')
			case 't':
				result = append(result, '\t')
			case 'u':
				var hex []rune
				for i := 0; i < 4; i++ {
					hex = append(hex, str.Read())
				}
				unicode, err := strconv.ParseInt(string(hex), 16, 32)
				if err != nil {
					return "", utils.WrappedJsonErrorf(str.pathString(), "Invalid unicode escape sequence at line %d, column %d", str.line, str.column)
				}
				result = append(result, string(rune(unicode))...)
			default:
				result = append(result, '\\')
				result = utf8.AppendRune(result, token)
			}
		} else {
			result = utf8.AppendRune(result, token)
		}
	}
	return string(result), nil
}
