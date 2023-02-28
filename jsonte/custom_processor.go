package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/gammazero/deque"
	"strings"
)

// ProcessLangFile processes a lang file replacing all the jsonte expressions with their values
func ProcessLangFile(input string, scope types.JsonObject) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	input = string(str)
	inputLen := len(input)
	globalScope := deque.Deque[types.JsonObject]{}
	globalScope.PushBack(scope)

	matches := map[string]string{}
	started := false
	startedString := false
	stringType := '"'
	bracketCount := 0
	var currentMatch strings.Builder
	var debugMatch strings.Builder
	for i := 0; i < inputLen; i++ {
		char := rune(input[i])
		if !started {
			if i+2 < inputLen && char == '#' && input[i+1] == '#' && input[i+2] == '{' {
				started = true
				bracketCount = 1
				i += 2
				currentMatch.Reset()
				debugMatch.Reset()
				continue
			}
		} else {
			if char == '\n' {
				return "", burrito.WrappedErrorf("The expression '%s' is not closed.", debugMatch.String())
			}
			if char == '"' || char == '\'' {
				if !startedString {
					startedString = true
					stringType = char
				} else if char == stringType {
					startedString = false
				}
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '{' && !startedString {
				bracketCount++
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '}' && !startedString {
				bracketCount--
				if bracketCount == 0 {
					started = false
					match := currentMatch.String()
					result, err := Eval(match, globalScope, "#")
					if err != nil {
						return "", burrito.WrapErrorf(err, "Failed to evaluate expression '%s'", debugMatch.String())
					}
					if result.Value == nil {
						return "", burrito.WrappedErrorf("The expression '%s' evaluated to null.", debugMatch.String())
					}
					if result.Action == types.Value {
						matches[debugMatch.String()] = types.ToString(result.Value)
					} else {
						return "", burrito.WrappedErrorf("The expression '%s' evaluated to an action.", debugMatch.String())
					}
				} else {
					currentMatch.WriteRune(char)
					debugMatch.WriteRune(char)
				}
			} else if char == '\\' {
				nextChar := input[i+1]
				if nextChar == 'n' {
					currentMatch.WriteRune('\n')
				} else if nextChar == 't' {
					currentMatch.WriteRune('\t')
				} else if nextChar == 'r' {
					currentMatch.WriteRune('\r')
				} else if nextChar == 'b' {
					currentMatch.WriteRune('\b')
				} else {
					currentMatch.WriteRune(rune(nextChar))
				}
				debugMatch.WriteRune(char)
				debugMatch.WriteRune(rune(nextChar))
				i++
			} else {
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			}
		}
	}

	for s, s2 := range matches {
		input = strings.ReplaceAll(input, "##{"+s+"}", s2)
	}

	return input, nil
}

// ProcessMCFunction processes a file replacing all the jsonte expressions with their values
func ProcessMCFunction(input string, scope types.JsonObject) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	input = string(str)
	inputLen := len(input)
	globalScope := deque.Deque[types.JsonObject]{}
	globalScope.PushBack(scope)

	matches := map[string]string{}
	started := false
	startedString := false
	stringType := '"'
	bracketCount := 0
	var currentMatch strings.Builder
	var debugMatch strings.Builder
	for i := 0; i < inputLen; i++ {
		char := rune(input[i])
		if !started {
			if i+1 < inputLen && char == '#' && input[i+1] == '{' {
				started = true
				bracketCount = 1
				i++
				currentMatch.Reset()
				debugMatch.Reset()
				continue
			}
		} else {
			if char == '\n' {
				return "", burrito.WrappedErrorf("The expression '%s' is not closed.", debugMatch.String())
			}
			if char == '"' || char == '\'' {
				if !startedString {
					startedString = true
					stringType = char
				} else if char == stringType {
					startedString = false
				}
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '{' && !startedString {
				bracketCount++
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '}' && !startedString {
				bracketCount--
				if bracketCount == 0 {
					started = false
					match := currentMatch.String()
					result, err := Eval(match, globalScope, "#")
					if err != nil {
						return "", burrito.WrapErrorf(err, "Failed to evaluate expression '%s'", debugMatch.String())
					}
					if result.Value == nil {
						return "", burrito.WrappedErrorf("The expression '%s' evaluated to null.", debugMatch.String())
					}
					if result.Action == types.Value {
						matches[debugMatch.String()] = types.ToString(result.Value)
					} else {
						return "", burrito.WrappedErrorf("The expression '%s' evaluated to an action.", debugMatch.String())
					}
				} else {
					currentMatch.WriteRune(char)
					debugMatch.WriteRune(char)
				}
			} else if char == '\\' {
				nextChar := input[i+1]
				if nextChar == 'n' {
					currentMatch.WriteRune('\n')
				} else if nextChar == 't' {
					currentMatch.WriteRune('\t')
				} else if nextChar == 'r' {
					currentMatch.WriteRune('\r')
				} else if nextChar == 'b' {
					currentMatch.WriteRune('\b')
				} else {
					currentMatch.WriteRune(rune(nextChar))
				}
				debugMatch.WriteRune(char)
				debugMatch.WriteRune(rune(nextChar))
				i++
			} else {
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			}
		}
	}

	for s, s2 := range matches {
		input = strings.ReplaceAll(input, "#{"+s+"}", s2)
	}

	return input, nil
}