package json

import (
	"encoding/json"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"strconv"
)

const UnexpectedTokenError = "Unexpected token '%c' at line %d, column %d"
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
	TokenEof            = 0
)

type StringReader struct {
	str    []byte
	marker int
	column int
	line   int
}

func NewStringReader(str []byte) *StringReader {
	return &StringReader{
		str:    str,
		marker: 0,
		column: 0,
		line:   1,
	}
}

func (sr *StringReader) Peek() byte {
	if sr.marker >= len(sr.str) {
		return TokenEof
	}
	return sr.str[sr.marker]
}

func (sr *StringReader) Read() byte {
	c := sr.Peek()
	if c == TokenNewline {
		sr.line++
		sr.column = 0
	} else {
		sr.column++
	}
	sr.marker++
	return c
}

func UnmarshallJSONC(str []byte) (interface{}, error) {
	reader := NewStringReader(str)
	skipWhitespace(reader)
	token := reader.Peek()
	var object interface{}
	var err error
	if token == TokenOpenObject {
		object, err = parseObject(reader, "#")
		if err != nil {
			return nil, err
		}
	} else if token == TokenOpenArray {
		object, err = parseArray(reader, "#")
		if err != nil {
			return nil, err
		}
	} else if token == TokenDoubleQuote {
		object, err = parseString(reader, "#")
		if err != nil {
			return nil, err
		}
	} else {
		object, err = parsePrimitive(reader, "#")
		if err != nil {
			return nil, err
		}
	}
	skipWhitespace(reader)
	if reader.Peek() != TokenEof {
		return object, utils.WrappedJsonErrorf("#", ExpectedEofError, reader.line, reader.column)
	}
	return object, nil
}

func MarshalJSONC(object interface{}, pretty bool) ([]byte, error) {
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
			return []byte(strconv.FormatFloat(object.(float64), 'f', 1, 64)), nil
		}
		return []byte(strconv.FormatFloat(object.(float64), 'f', -1, 64)), nil
	case int:
		return []byte(strconv.FormatInt(int64(object.(int)), 10)), nil
	case int64:
		return []byte(strconv.FormatInt(object.(int64), 10)), nil
	case bool:
		if object.(bool) {
			return []byte("true"), nil
		} else {
			return []byte("false"), nil
		}
	case nil:
		return []byte("null"), nil
	default:
		return nil, burrito.WrappedErrorf("Unsupported type %T", object)
	}
}

func writeObject(object utils.NavigableMap[string, interface{}], pretty bool, indent int) ([]byte, error) {
	var result []byte
	result = append(result, TokenOpenObject)
	if pretty {
		result = append(result, TokenNewline)
	}
	indent++
	for i, key := range object.Keys() {
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
				result = append(result, []byte(strconv.FormatFloat(value.(float64), 'f', 1, 64))...)
			} else {
				result = append(result, []byte(strconv.FormatFloat(value.(float64), 'f', -1, 64))...)
			}
		case int:
			result = append(result, []byte(strconv.FormatInt(int64(value.(int)), 10))...)
		case int32:
			result = append(result, []byte(strconv.FormatInt(int64(value.(int32)), 10))...)
		case bool:
			result = append(result, []byte(strconv.FormatBool(value.(bool)))...)
		case nil:
			result = append(result, []byte("null")...)
		default:
			return result, burrito.WrappedErrorf("Unsupported type %T", value)
		}
		if i < len(object.Keys())-1 {
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

func writeArray(arr []interface{}, pretty bool, indent int) ([]byte, error) {
	var result []byte
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
				result = append(result, []byte(strconv.FormatFloat(value.(float64), 'f', 1, 64))...)
			} else {
				result = append(result, []byte(strconv.FormatFloat(value.(float64), 'f', -1, 64))...)
			}
		case int:
			result = append(result, []byte(strconv.FormatInt(int64(value.(int)), 10))...)
		case int32:
			result = append(result, []byte(strconv.FormatInt(int64(value.(int32)), 10))...)
		case bool:
			result = append(result, []byte(strconv.FormatBool(value.(bool)))...)
		case nil:
			result = append(result, []byte("null")...)
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

func writeString(s string) []byte {
	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
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

func indentBytes(indent int) []byte {
	var result []byte
	for i := 0; i < indent*2; i++ {
		result = append(result, TokenSpace)
	}
	return result
}

func isWhitespace(token byte) bool {
	return token == TokenSpace || token == TokenTab || token == TokenNewline || token == TokenCarriageReturn
}

func skipWhitespace(str *StringReader) {
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
		}
		break
	}
}

func parseObject(str *StringReader, p string) (utils.NavigableMap[string, interface{}], error) {
	result := utils.NewNavigableMap[string, interface{}]()

	if str.Read() != TokenOpenObject {
		return result, utils.WrappedJsonErrorf(p, ExpectedTokenError, TokenOpenObject, str.line, str.column)
	}

	canStartComment := false
	comma := false
	open := true

	for {
		token := str.Peek()
		if token == TokenEof {
			return result, utils.WrappedJsonErrorf(p, UnexpectedEofError, str.line, str.column)
		} else if isWhitespace(token) {
			str.Read()
		} else if token == TokenSlash {
			canStartComment = parseComment(str, canStartComment)
		} else if token == TokenAsterisk {
			canStartComment = parseComment(str, canStartComment)
		} else if token == TokenDoubleQuote && (comma || open) {
			comma = false
			open = false
			key, err := parseString(str, p)
			if err != nil {
				return result, err
			}
			skipWhitespace(str)
			if str.Read() != TokenColon {
				return result, utils.WrappedJsonErrorf(p, ExpectedTokenError, TokenColon, str.line, str.column)
			}
			skipWhitespace(str)
			peekToken := str.Peek()
			if peekToken == TokenOpenObject {
				value, err := parseObject(str, p+"/"+key)
				if err != nil {
					return result, err
				}
				result.Put(key, value)
			} else if peekToken == TokenOpenArray {
				value, err := parseArray(str, p+"/"+key)
				if err != nil {
					return result, err
				}
				result.Put(key, value)
			} else if peekToken == TokenDoubleQuote {
				value, err := parseString(str, p+"/"+key)
				if err != nil {
					return result, err
				}
				result.Put(key, value)
			} else {
				value, err := parsePrimitive(str, p+"/"+key)
				if err != nil {
					return result, err
				}
				result.Put(key, value)
			}
		} else if token == TokenComma && !comma && !open {
			comma = true
			str.Read()
		} else if token == TokenCloseObject && !comma {
			str.Read()
			return result, nil
		} else {
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
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

func parsePrimitive(str *StringReader, p string) (interface{}, error) {
	token := str.Peek()
	if token >= '0' && token <= '9' || token == '-' {
		return parseNumber(str), nil
	} else if token == 't' {
		return parseTrue(str, p)
	} else if token == 'f' {
		return parseFalse(str, p)
	} else if token == 'n' {
		return parseNull(str, p)
	} else {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
}

func parseTrue(str *StringReader, p string) (interface{}, error) {
	token := str.Read()
	if token != 't' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'r' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'u' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'e' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	return true, nil
}

func parseFalse(str *StringReader, p string) (interface{}, error) {
	token := str.Read()
	if token != 'f' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'a' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'l' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 's' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'e' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	return false, nil
}

func parseNull(str *StringReader, p string) (interface{}, error) {
	token := str.Read()
	if token != 'n' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'u' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'l' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	token = str.Read()
	if token != 'l' {
		return nil, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
	}
	return nil, nil
}

func parseNumber(str *StringReader) json.Number {
	start := str.marker
	for {
		token := str.Peek()
		if token >= '0' && token <= '9' || token == '.' || token == 'e' || token == 'E' || token == '-' {
			str.Read()
		} else {
			break
		}
	}
	return json.Number(str.str[start:str.marker])
}

func parseArray(str *StringReader, p string) ([]interface{}, error) {
	result := make([]interface{}, 0)

	if str.Read() != TokenOpenArray {
		return result, utils.WrappedJsonErrorf(p, ExpectedTokenError, TokenOpenArray, str.line, str.column)
	}

	comma := false
	open := true
	canStartComment := false

	for {
		token := str.Peek()
		if token == TokenEof {
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
		} else if isWhitespace(token) {
			str.Read()
		} else if token == TokenSlash {
			canStartComment = parseComment(str, canStartComment)
		} else if token == TokenAsterisk {
			canStartComment = parseComment(str, canStartComment)
		} else if token == TokenOpenArray && (comma || open) {
			value, err := parseArray(str, fmt.Sprintf("%s[%d]", p, len(result)))
			if err != nil {
				return result, err
			}
			result = append(result, value)
			comma = false
			open = false
		} else if token == TokenOpenObject && (comma || open) {
			value, err := parseObject(str, fmt.Sprintf("%s[%d]", p, len(result)))
			if err != nil {
				return result, err
			}
			result = append(result, value)
			comma = false
			open = false
		} else if token == TokenDoubleQuote && (comma || open) {
			value, err := parseString(str, fmt.Sprintf("%s[%d]", p, len(result)))
			if err != nil {
				return result, err
			}
			result = append(result, value)
			comma = false
			open = false
		} else if token == TokenComma && !comma && !open {
			comma = true
			str.Read()
		} else if token == TokenCloseArray && !comma {
			str.Read()
			return result, nil
		} else if comma || open {
			value, err := parsePrimitive(str, fmt.Sprintf("%s[%d]", p, len(result)))
			if err != nil {
				return result, err
			}
			result = append(result, value)
			comma = false
			open = false
		} else {
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenError, token, str.line, str.column)
		}
	}
}

func parseString(str *StringReader, p string) (string, error) {
	var result []byte
	str.Read()
	for {
		token := str.Read()
		if token == TokenEof {
			return "", utils.WrappedJsonErrorf(p, UnexpectedEofError, str.line, str.column)
		}
		if token == TokenDoubleQuote {
			break
		}
		if token == TokenBackslash {
			token = str.Read()
			if token == TokenEof {
				return "", utils.WrappedJsonErrorf(p, UnexpectedEofError, str.line, str.column)
			}
			switch token {
			case TokenDoubleQuote:
				result = append(result, '"')
			case TokenBackslash:
				result = append(result, '\\')
			case TokenSlash:
				result = append(result, '/')
			case TokenAsterisk:
				result = append(result, '*')
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
				var hex []byte
				for i := 0; i < 4; i++ {
					hex = append(hex, str.Read())
				}
				unicode, err := strconv.ParseInt(string(hex), 16, 32)
				if err != nil {
					return "", utils.WrappedJsonErrorf(p, "Invalid unicode escape sequence at line %d, column %d", str.line, str.column)
				}
				result = append(result, []byte(string(rune(unicode)))...)
			default:
				return "", utils.WrappedJsonErrorf(p, "Invalid escape sequence at line %d, column %d", str.line, str.column)
			}
		} else {
			result = append(result, token)
		}
	}
	return string(result), nil
}
