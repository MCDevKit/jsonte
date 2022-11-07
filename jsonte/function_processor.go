package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"strings"
)

// ProcessMCFunction processes an mcfunction file replacing all the jsonte expressions with their values
func ProcessMCFunction(input string, scope utils.NavigableMap[string, interface{}]) (string, error) {
	globalScope := deque.Deque[interface{}]{}
	globalScope.PushBack(scope)

	matches := map[string]string{}
	started := false
	startedString := false
	stringType := '"'
	bracketCount := 0
	var currentMatch strings.Builder
	var debugMatch strings.Builder
	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		if char == '#' && !started {
			if input[i+1] == '{' {
				started = true
				bracketCount = 1
				i++
				currentMatch.Reset()
				debugMatch.Reset()
				continue
			}
		} else if started {
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
					if result.Action == utils.Value {
						matches[debugMatch.String()] = utils.ToString(result.Value)
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
