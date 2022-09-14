package jsonte

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"regexp"
	"strings"
)

var mcFunctionPattern, _ = regexp.Compile("#\\{((?:\\\\.|[^{}])+)}")

func ProcessMCFunction(input string, scope utils.JsonObject) (string, error) {
	globalScope := deque.Deque[interface{}]{}
	globalScope.PushBack(scope)
	matches := mcFunctionPattern.FindAllString(input, -1)
	replacements := make(map[string]string, len(matches))
	for _, match := range matches {
		result, err := Eval(strings.TrimPrefix(strings.TrimSuffix(match, "}"), "#{"), globalScope, "#")
		if err != nil {
			return "", err
		}
		if result.Value == nil {
			return "", utils.WrappedErrorf("The expression '%s' evaluated to null.", match)
		}
		if result.Action == utils.Value {
			replacements[match] = utils.ToString(result.Value)
		} else {
			return "", utils.WrappedErrorf("The expression '%s' evaluated to an action.", match)
		}
	}
	result := templatePattern.ReplaceAllStringFunc(input, func(match string) string {
		return replacements[match]
	})
	return result, nil
}
