package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"hash/fnv"
	"regexp"
	"strings"
	"unicode"
)

func RegisterStringFunctions() {
	RegisterFunction(JsonFunction{
		Name:       "replace",
		Body:       replace,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "join",
		Body:       join,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "contains",
		Body:       stringContains,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "split",
		Body:       split,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "indexOf",
		Body:       stringIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "lastIndexOf",
		Body:       stringLastIndexOf,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "hash",
		Body:       hash,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "toUpperCase",
		Body:       toUpperCase,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "toLowerCase",
		Body:       toLowerCase,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "substring",
		Body:       substring,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "substring",
		Body:       substringFrom,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "capitalize",
		Body:       captialize,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "startsWith",
		Body:       startsWith,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "endsWith",
		Body:       endsWith,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "regexReplace",
		Body:       regexReplace,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "chars",
		Body:       chars,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "length",
		Body:       length,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "trim",
		Body:       trim,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "title",
		Body:       title,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Name:       "swapCase",
		Body:       swapCase,
		IsInstance: true,
	})
}

func replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func join(strs utils.JsonArray, sep string) string {
	arr := make([]string, len(strs))
	for i, s := range strs {
		arr[i] = utils.ToString(s)
	}
	return strings.Join(arr, sep)
}

func stringContains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func split(str, sep string) utils.JsonArray {
	strs := strings.Split(str, sep)
	arr := make(utils.JsonArray, len(strs))
	for i, s := range strs {
		arr[i] = s
	}
	return arr
}

func stringIndexOf(str, substr string) utils.JsonNumber {
	return utils.ToNumber(strings.Index(str, substr))
}

func stringLastIndexOf(str, substr string) utils.JsonNumber {
	return utils.ToNumber(strings.LastIndex(str, substr))
}

func hash(str string) (utils.JsonNumber, error) {
	a := fnv.New32a()
	_, err := a.Write([]byte(str))
	if err != nil {
		return utils.ToNumber(0), utils.WrapErrorf(err, "Failed to hash string")
	}
	return utils.ToNumber(a.Sum32()), nil
}

func toUpperCase(str string) string {
	return cases.Upper(language.Und).String(str)
}

func toLowerCase(str string) string {
	return cases.Lower(language.Und).String(str)
}

func substring(str string, start, end utils.JsonNumber) string {
	return str[start.IntValue():end.IntValue()]
}

func substringFrom(str string, start utils.JsonNumber) string {
	return str[start.IntValue():]
}

func captialize(str string) string {
	return cases.Upper(language.Und).String(str[:1]) + cases.Lower(language.Und).String(str[1:])
}

func title(str string) string {
	return cases.Title(language.Und).String(str)
}

func swapCase(str string) string {
	for i, c := range str {
		if unicode.IsUpper(c) {
			str = str[:i] + strings.ToLower(string(c)) + str[i+1:]
		} else if unicode.IsLower(c) {
			str = str[:i] + strings.ToUpper(string(c)) + str[i+1:]
		}
	}
	return str
}

func startsWith(str, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func endsWith(str, substr string) bool {
	return strings.HasSuffix(str, substr)
}

func regexReplace(str, pattern, repl string) (string, error) {
	compile, err := regexp.Compile(pattern)
	if err != nil {
		return "", utils.WrapErrorf(err, "Failed to compile regex pattern")
	}
	return compile.ReplaceAllString(str, repl), nil
}

func chars(str string) utils.JsonArray {
	arr := make(utils.JsonArray, len(str))
	for i, c := range str {
		arr[i] = string(c)
	}
	return arr
}

func length(str string) utils.JsonNumber {
	return utils.ToNumber(len(str))
}

func trim(str string) string {
	return strings.Trim(str, " \t\n\r")
}
