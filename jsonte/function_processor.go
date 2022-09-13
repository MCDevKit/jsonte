package jsonte

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"regexp"
)

var mcFunctionPattern, _ = regexp.Compile("#\\{((?:\\\\.|[^{}])+)}")

func ProcessMCFunction(input string, scope utils.JsonObject) (string, error) {
	globalScope := deque.Deque[interface{}]{}
	globalScope.PushBack(scope)
	matches := mcFunctionPattern.FindAllString(input, -1)
	replacements := make(map[string]string, len(matches))
	for _, match := range matches {
		result, err := Eval(match, globalScope, "#")
		if err != nil {
			return "", err
		}
		if result.Value == nil {
			return "", &utils.TemplatingError{
				Message: "The expression evaluated to null",
			}
		}
		if result.Action == utils.Value {
			replacements[match] = utils.ToString(result.Value)
		} else {
			return "", &utils.TemplatingError{
				Message: "Cannot execute action here",
			}
		}
	}
	result := templatePattern.ReplaceAllStringFunc(input, func(match string) string {
		return replacements[match]
	})
	return result, nil
}
