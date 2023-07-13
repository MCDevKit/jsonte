package json

import (
	"encoding/json"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"strconv"
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

type StringReader struct {
	str    []rune
	marker int
	column int
	line   int
}

func NewStringReader(str []byte) *StringReader {
	return &StringReader{
		str:    []rune(string(str)),
		marker: 0,
		column: 0,
		line:   1,
	}
}

func (sr *StringReader) Peek() rune {
	if sr.marker >= len(sr.str) {
		return TokenEof
	}
	return sr.str[sr.marker]
}

func (sr *StringReader) Read() rune {
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
	str, err := ConvertToUTF8(str)
	if err != nil {
		return nil, burrito.PassError(err)
	}
	reader := NewStringReader(str)
	err = skipWhitespace(reader)
	if err != nil {
		return nil, err
	}
	token := reader.Peek()
	var object interface{}
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
	err = skipWhitespace(reader)
	if err != nil {
		return nil, err
	}
	if reader.Peek() != TokenEof {
		return object, utils.WrappedJsonErrorf("#", ExpectedEofError, reader.line, reader.column)
	}
	return object, nil
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
		} else if canStartComment && token != TokenSlash && token != TokenAsterisk {
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenExpectedError, token, TokenSlash, str.line, str.column)
		} else if token == TokenDoubleQuote && (comma || open) {
			comma = false
			open = false
			key, err := parseString(str, p)
			if err != nil {
				return result, err
			}
			err = skipWhitespace(str)
			if err != nil {
				return result, err
			}
			if str.Read() != TokenColon {
				return result, utils.WrappedJsonErrorf(p, ExpectedTokenError, TokenColon, str.line, str.column)
			}
			err = skipWhitespace(str)
			if err != nil {
				return result, err
			}
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
	if token >= '0' && token <= '9' || token == '-' || token == '+' || token == '.' {
		return parseNumber(str)
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

func parseConstant(str *StringReader, p string, constant string) error {
	runes := []rune(constant)
	for i := 0; i < len(runes); i++ {
		token := str.Read()
		if token != runes[i] {
			return utils.WrappedJsonErrorf(p, UnexpectedTokenExpectedError, token, constant[i], str.line, str.column)
		}
	}
	return nil
}

func parseTrue(str *StringReader, p string) (interface{}, error) {
	return true, parseConstant(str, p, "true")
}

func parseFalse(str *StringReader, p string) (interface{}, error) {
	return false, parseConstant(str, p, "false")
}

func parseNull(str *StringReader, p string) (interface{}, error) {
	return nil, parseConstant(str, p, "null")
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
			return json.Number(str.str[start:str.marker]), nil
		}
	}
	//hadDecimalPoint := false
	//hadSign := false
	//for {
	//	token := str.Peek()
	//	if token >= '0' && token <= '9' || token == '.' || token == 'e' || token == 'E' || token == '-' || token == '+' {
	//		if token == '.' {
	//			if hadDecimalPoint {
	//				return "", utils.WrappedJsonErrorf("", UnexpectedTokenError, token, str.line, str.column)
	//			}
	//			hadDecimalPoint = true
	//		}
	//		str.Read()
	//	} else {
	//		break
	//	}
	//}
	//return json.Number(str.str[start:str.marker]), nil
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
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenExpectedError, token, TokenEof, str.line, str.column)
		} else if isWhitespace(token) {
			str.Read()
		} else if token == TokenSlash {
			canStartComment = parseComment(str, canStartComment)
		} else if token == TokenAsterisk {
			canStartComment = parseComment(str, canStartComment)
		} else if canStartComment && token != TokenSlash && token != TokenAsterisk {
			return result, utils.WrappedJsonErrorf(p, UnexpectedTokenExpectedError, token, TokenSlash, str.line, str.column)
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
	var result []rune
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
					return "", utils.WrappedJsonErrorf(p, "Invalid unicode escape sequence at line %d, column %d", str.line, str.column)
				}
				result = append(result, []rune(string(rune(unicode)))...)
			default:
				result = append(result, '\\')
				result = append(result, token)
			}
		} else {
			result = append(result, token)
		}
	}
	return string(result), nil
}
