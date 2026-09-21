package jsonte

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	molang "github.com/stirante/molang-go"
	"github.com/stirante/molang-go/lexer"
	"github.com/stirante/molang-go/parser"
	"github.com/stirante/molang-go/printer"
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

// ProcessMolangFile processes a molang file replacing all the jsonte expressions with their values, then minifies it.
// It returns an *InvalidMolangError if the result is not valid Molang.
func ProcessMolangFile(input string, scope *types.JsonObject) (string, error) {
	return ProcessMolangFileWithValidation(input, scope, true)
}

// ProcessMolangFileWithValidation is ProcessMolangFile with validation that can be turned off.
// Valid Molang is minified by molang-go. When validate is false, anything else (e.g. a fragment with
// placeholders replaced later) is minified token by token instead of returning an *InvalidMolangError.
func ProcessMolangFileWithValidation(input string, scope *types.JsonObject, validate bool) (string, error) {
	return processMolangFile(input, validate, scopeEvaluator(scope))
}

// EnableMolangLoading makes the loadMolang function process molang files with scope as the global scope.
// scopeFiles are the .molang files found in scope directories: loadText returns them processed.
// Call it after the JSON scope files are merged and before scope scripts run, so scripts can load molang too.
func EnableMolangLoading(scope *types.JsonObject, scopeFiles []string, validate bool) {
	// Results of expressions that do not mention any parameter are the same for every set of parameters,
	// because lambdas from the scope cannot see the parameters. Caching them avoids evaluating heavy
	// helper functions again for every loadMolang call with different parameters.
	var independent sync.Map
	functions.SetMolangLoader(func(source string, params *types.JsonObject) (string, error) {
		scopes := deque.Deque[*types.JsonObject]{}
		scopes.PushBack(scope)
		scopes.PushBack(params)
		keys := params.Keys()
		eval := func(expr string) (Result, error) {
			for _, key := range keys {
				// Conservative: any occurrence, even inside a string like this['key'], counts as a use
				if strings.Contains(expr, key) {
					return Eval(expr, scopes, "#")
				}
			}
			if cached, ok := independent.Load(expr); ok {
				return cached.(Result), nil
			}
			result, err := Eval(expr, scopes, "#")
			if err == nil {
				independent.Store(expr, result)
			}
			return result, err
		}
		return processMolangFile(source, validate, eval)
	}, scopeFiles)
}

// scopeEvaluator evaluates expressions with scope as the global scope.
func scopeEvaluator(scope *types.JsonObject) func(expr string) (Result, error) {
	var scopes deque.Deque[*types.JsonObject]
	return func(expr string) (Result, error) {
		if scopes.Len() == 0 {
			scopes.PushBack(scope)
		}
		return Eval(expr, scopes, "#")
	}
}

// processMolangFile processes a molang file, evaluating its expressions with eval.
func processMolangFile(input string, validate bool, eval func(expr string) (Result, error)) (string, error) {
	str, err := json.ConvertToUTF8([]byte(input))
	if err != nil {
		return input, burrito.PassError(err)
	}
	stripped := stripMolangComments(string(str))
	output, err := processString(stripped, "#", "", eval)
	if err != nil {
		return "", burrito.PassError(err)
	}
	if strings.TrimSpace(output) == "" {
		return "", nil
	}
	prog, parseErr := molang.ParseWith(output, molang.Extensions{Comments: true})
	if parseErr != nil {
		if validate {
			return "", newInvalidMolangError(parseErr, output)
		}
		return utils.MinifyWithShortAccessors(output), nil
	}
	return printer.Minify(prog), nil
}

// InvalidMolangError reports a processed molang file that is not valid Molang.
// Line and Column point into the processed file, after comments are removed and expressions are evaluated.
type InvalidMolangError struct {
	Msg     string
	Line    int
	Column  int
	Snippet string
}

func (e *InvalidMolangError) Error() string {
	return fmt.Sprintf("%s (line %d, column %d of the processed file: %s)", e.Msg, e.Line, e.Column, e.Snippet)
}

// molangParseError returns the message and the position in the source of a molang-go parse or lex error.
func molangParseError(err error) (string, int, bool) {
	var parseErr *parser.Error
	var lexErr *lexer.Error
	switch {
	case errors.As(err, &parseErr):
		return parseErr.Msg, parseErr.Pos, true
	case errors.As(err, &lexErr):
		return lexErr.Msg, lexErr.Pos, true
	}
	return "", 0, false
}

// newInvalidMolangError replaces molang-go's error text, which embeds the whole source, with the message,
// the line and column in the processed source, and that line.
func newInvalidMolangError(err error, source string) error {
	msg, pos, ok := molangParseError(err)
	if !ok {
		return err
	}
	line, column, snippet := molangErrorLocation(source, pos)
	return &InvalidMolangError{Msg: msg, Line: line, Column: column, Snippet: snippet}
}

// molangErrorLocation returns the line and column of pos in source and the line around it,
// shortened to keep long generated lines readable.
func molangErrorLocation(source string, pos int) (int, int, string) {
	pos = min(max(pos, 0), len(source))
	lineStart := strings.LastIndexByte(source[:pos], '\n') + 1
	lineEnd := strings.IndexByte(source[pos:], '\n')
	if lineEnd < 0 {
		lineEnd = len(source)
	} else {
		lineEnd += pos
	}
	line := strings.Count(source[:pos], "\n") + 1
	column := pos - lineStart + 1
	// Keep long generated lines readable: show up to 60 bytes on each side of the error
	from, to := max(lineStart, pos-60), min(lineEnd, pos+60)
	snippet := strings.TrimSpace(source[from:to])
	if from > lineStart {
		snippet = "..." + snippet
	}
	if to < lineEnd {
		snippet += "..."
	}
	return line, column, snippet
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
	return processString(input, startToken, endToken, scopeEvaluator(scope))
}

// processString is ProcessString with a custom way to evaluate expressions.
func processString(input string, startToken, endToken string, eval func(expr string) (Result, error)) (string, error) {
	templateMatches, err := FindTemplateMatches(input, startToken, endToken)
	if err != nil {
		return "", burrito.PassError(err)
	}
	if len(templateMatches) == 0 {
		return input, nil
	}
	var sb strings.Builder
	sb.Grow(len(input))
	lastMatchEnd := 0
	for _, match := range templateMatches {
		if match.StartByte > lastMatchEnd {
			sb.WriteString(input[lastMatchEnd:match.StartByte])
		}
		result, err := eval(match.Match)
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
