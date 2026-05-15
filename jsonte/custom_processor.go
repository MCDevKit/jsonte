package jsonte

import (
	"strings"
	"sync"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
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

// stripMolangLineComment removes a `# ` comment from a single line without
// touching `#{` template prefixes or content inside string literals.
func stripMolangLineComment(line string) string {
	runes := []rune(line)
	inString := false
	var quote rune
	for j, ch := range runes {
		if inString {
			if ch == quote && (j == 0 || runes[j-1] != '\\') {
				inString = false
			}
		} else {
			switch ch {
			case '"', '\'':
				inString = true
				quote = ch
			case '#':
				if j+1 < len(runes) && runes[j+1] == ' ' {
					return strings.TrimRight(string(runes[:j]), " \t")
				}
			}
		}
	}
	return line
}

// stripMolangComments removes `# ` line comments from molang source line by line.
func stripMolangComments(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = stripMolangLineComment(line)
	}
	return strings.Join(lines, "\n")
}

// ProcessMolangFile processes a molang file replacing all the jsonte expressions with their values, then minifies it
func ProcessMolangFile(input string, scope *types.JsonObject) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	stripped := stripMolangComments(string(str))
	output, err := ProcessString(stripped, scope, "#", "")
	if err != nil {
		return "", burrito.PassError(err)
	}
	return utils.MinifyWithShortAccessors(output), nil
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
	templateMatches, err := FindTemplateMatches(input, startToken, endToken)
	if err != nil {
		return "", burrito.PassError(err)
	}
	if len(templateMatches) == 0 {
		return input, nil
	}
	globalScope := deque.Deque[*types.JsonObject]{}
	globalScope.PushBack(scope)
	var sb strings.Builder
	sb.Grow(len(input))
	lastMatchEnd := 0
	for _, match := range templateMatches {
		if match.StartByte > lastMatchEnd {
			sb.WriteString(input[lastMatchEnd:match.StartByte])
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
			lastMatchEnd = match.EndByte
		} else {
			return "", burrito.WrappedErrorf("The expression '%s' evaluated to an action.", match.EscapedMatch)
		}
	}
	if lastMatchEnd < len(input) {
		sb.WriteString(input[lastMatchEnd:])
	}

	return sb.String(), nil
}

type TemplateMatch struct {
	Match        string
	EscapedMatch string
	Start        int
	Length       int
	StartByte    int
	EndByte      int
}

// templateMatchCache caches FindTemplateMatches results for static strings.
// Keys are the raw input string (for the common "{"/"}" token pair).
var templateMatchCache sync.Map

// FindTemplateMatches processes a string replacing all the jsonte expressions with their values
func FindTemplateMatches(input, startToken, endToken string) ([]TemplateMatch, error) {
	// Fast-path: avoid expensive rune conversion if no template marker is present.
	marker := startToken + "{"
	if !strings.Contains(input, marker) {
		return nil, nil
	}
	// Cache lookup for the common case (standard JSON template tokens).
	isStandard := startToken == "{" && endToken == "}"
	if isStandard {
		if cached, ok := templateMatchCache.Load(input); ok {
			return cached.([]TemplateMatch), nil
		}
	}
	inputRunes := []rune(input)
	runeByteOffsets := make([]int, 0, len(inputRunes)+1)
	for byteOffset := range input {
		runeByteOffsets = append(runeByteOffsets, byteOffset)
	}
	runeByteOffsets = append(runeByteOffsets, len(input))
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
						StartByte:    runeByteOffsets[startIndex],
						EndByte:      runeByteOffsets[i+1],
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

	if isStandard {
		templateMatchCache.Store(input, matches)
	}
	return matches, nil
}
