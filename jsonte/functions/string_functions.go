package functions

import (
	"jsonte/jsonte/utils"
)

func RegisterStringFunctions() {
	RegisterFunction(JsonFunction{
		/**
		 * Returns the length of this string.
		 * @param str string: String
		 * @example
		 * <code>
		 * {
		 *   "$template": {
		 *     "$comment": "The field below will be ['h', 'e', 'l', 'l', 'o']",
		 *     "test": "{{chars('hello')}}"
		 *   }
		 * }
		 * </code>
		 */
		Name:       "length",
		Body:       length,
		IsInstance: true,
	})
}

func length(str string) (utils.JsonNumber, error) {
	return utils.ToNumber(len(str)), nil
}
