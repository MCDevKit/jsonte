package jsonte

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/gammazero/deque"
	"strings"
)

// ProcessLangFile processes a lang file replacing all the jsonte expressions with their values
func ProcessLangFile(input string, scope *types.JsonObject) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	lines := strings.Split(string(str), "\n")
	for i, line := range lines {
		lines[i], err = ProcessString(line, scope, "##", "")
		if err != nil {
			return "", burrito.WrapErrorf(err, "Failed to process line %d", i+1)
		}
	}
	return strings.Join(lines, "\n"), nil
}

// ProcessMCFunction processes a file replacing all the jsonte expressions with their values
func ProcessMCFunction(input string, scope *types.JsonObject) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	output, err := ProcessString(string(str), scope, "#", "")
	if err != nil {
		return "", burrito.PassError(err)
	}
	return output, nil
}

// ProcessString processes a string replacing all the jsonte expressions with their values
func ProcessString(input string, scope *types.JsonObject, startToken, endToken string) (string, error) {
	runes := []rune(input)
	templateMatches, err := FindTemplateMatches(input, startToken, endToken)
	if err != nil {
		return "", burrito.PassError(err)
	}
	globalScope := deque.Deque[*types.JsonObject]{}
	globalScope.PushBack(scope)
	var sb strings.Builder
	lastMatchEnd := 0
	for _, match := range templateMatches {
		if match.Start > lastMatchEnd {
			sb.WriteString(string(runes[lastMatchEnd:match.Start]))
		}
		result, err := Eval(match.Match, globalScope, "#")
		if err != nil {
			return "", burrito.WrapErrorf(err, "Failed to evaluate expression '%s'", match.EscapedMatch)
		}
		if result.Value == nil {
			return "", burrito.WrappedErrorf("The expression '%s' evaluated to null.", match.EscapedMatch)
		}
		if result.Action == types.Value {
			sb.WriteString(types.ToString(result.Value))
			lastMatchEnd = match.Start + match.Length + 1
		} else {
			return "", burrito.WrappedErrorf("The expression '%s' evaluated to an action.", match.EscapedMatch)
		}
	}
	if lastMatchEnd < len(runes) {
		sb.WriteString(string(runes[lastMatchEnd:]))
	}

	return sb.String(), nil
}

type TemplateMatch struct {
	Match        string
	EscapedMatch string
	Start        int
	Length       int
}

// FindTemplateMatches processes a string replacing all the jsonte expressions with their values
func FindTemplateMatches(input, startToken, endToken string) ([]TemplateMatch, error) {
	inputRunes := []rune(input)
	startTokenRunes := []rune(startToken)
	endTokenRunes := []rune(endToken)
	inputLen := len(inputRunes)
	startLen := len(startTokenRunes)
	endLen := len(endTokenRunes)

	matches := make([]TemplateMatch, 0)
	started := false
	bracketCount := 0
	startIndex := 0
	var currentMatch strings.Builder
	var debugMatch strings.Builder
outerFor:
	for i := 0; i < inputLen; i++ {
		char := inputRunes[i]
		if !started {
			if i+startLen < inputLen && inputRunes[i+startLen] == '{' {
				for j := 0; j < startLen; j++ {
					if inputRunes[i+j] != startTokenRunes[j] {
						continue outerFor
					}
				}
				started = true
				bracketCount = 1
				startIndex = i
				i += startLen
				currentMatch.Reset()
				debugMatch.Reset()
				continue
			}
		} else {
			if char == '"' || char == '\'' {
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
				i++
				ended, debug := UnescapeStringToBuffer(inputRunes, &currentMatch, &i, char)
				if !ended {
					return matches, burrito.WrappedErrorf("The string '%s' is not closed.", debug)
				}
				debugMatch.WriteString(debug)
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '{' {
				bracketCount++
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			} else if char == '}' {
				bracketCount--
				if bracketCount == 0 && i+endLen < inputLen {
					for j := 0; j < endLen; j++ {
						if inputRunes[i+j] != endTokenRunes[j] {
							return matches, burrito.WrappedErrorf("The expression '%s' is not closed.", debugMatch.String())
						}
					}
					started = false
					i += endLen
					match := currentMatch.String()
					matches = append(matches, TemplateMatch{
						Match:        match,
						EscapedMatch: startToken + "{" + debugMatch.String() + "}" + endToken,
						Start:        startIndex,
						Length:       i - startIndex,
					})
				} else {
					currentMatch.WriteRune(char)
					debugMatch.WriteRune(char)
				}
			} else {
				currentMatch.WriteRune(char)
				debugMatch.WriteRune(char)
			}
		}
	}
	if started {
		return matches, burrito.WrappedErrorf("The expression '%s' is not closed.", debugMatch.String())
	}

	return matches, nil
}
