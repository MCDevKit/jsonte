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
	const group = "string"
	RegisterGroup(Group{
		Name:    group,
		Title:   "String functions",
		Summary: "String functions are related to manipulating string values.",
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "replace",
		Body:       replace,
		IsInstance: true,
		Docs: Docs{
			Summary: "Replaces all occurrences of a substring within a string with another string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "old",
					Summary: "The substring to replace.",
				},
				{
					Name:    "new",
					Summary: "The string to replace with.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this_is_a_test'",
    "test": "{{replace('this is a test', ' ', '_')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "join",
		Body:       join,
		IsInstance: true,
		Docs: Docs{
			Summary: "Joins an array of strings together into a single string, separated by a separator.",
			Arguments: []Argument{
				{
					Name:    "strings",
					Summary: "The array of strings to join together.",
				},
				{
					Name:    "separator",
					Summary: "The separator to use.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this,is,a,test'",
    "test": "{{join(['this', 'is', 'a', 'test'], ',')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "contains",
		Body:       stringContains,
		IsInstance: true,
		Docs: Docs{
			Summary: "Checks if a string contains a substring.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "substring",
					Summary: "The substring to search for.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{contains('this is a test', 'test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "split",
		Body:       split,
		IsInstance: true,
		Docs: Docs{
			Summary: "Splits a string into an array of strings, using a separator.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to split.",
				},
				{
					Name:    "separator",
					Summary: "The separator to use.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be ['this', 'is', 'a', 'test']",
    "test": "{{split('this is a test', ' ')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "indexOf",
		Body:       stringIndexOf,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the index of the first occurrence of a substring within a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "substring",
					Summary: "The substring to search for.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 8",
    "test": "{{indexOf('this is a test', 'test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "lastIndexOf",
		Body:       stringLastIndexOf,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the index of the last occurrence of a substring within a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "substring",
					Summary: "The substring to search for.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 8",
    "test": "{{lastIndexOf('this is a test', 'test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "hash",
		Body:       hash,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns an integer hash of a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to hash.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 2139996864",
    "test": "{{hash('this is a test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "toUpperCase",
		Body:       toUpperCase,
		IsInstance: true,
		Docs: Docs{
			Summary: "Converts a string to uppercase.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to convert.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'THIS IS A TEST'",
    "test": "{{toUpperCase('this is a test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "toLowerCase",
		Body:       toLowerCase,
		IsInstance: true,
		Docs: Docs{
			Summary: "Converts a string to lowercase.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to convert.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this is a test'",
    "test": "{{toLowerCase('THIS IS A TEST')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "substring",
		Body:       substring,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns a substring of a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to get the substring from.",
				},
				{
					Name:    "start",
					Summary: "The starting index of the substring.",
				},
				{
					Name:     "end",
					Summary:  "The ending index of the substring.",
					Optional: true,
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this'",
    "test": "{{substring('this is a test', 0, 4)}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "substring",
		Body:       substringFrom,
		IsInstance: true,
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "capitalize",
		Body:       captialize,
		IsInstance: true,
		Docs: Docs{
			Summary: "Capitalizes the first letter of a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to capitalize.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'This is a test'",
    "test": "{{capitalize('this is a test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "startsWith",
		Body:       startsWith,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns true if a string starts with a substring.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "substring",
					Summary: "The substring to search for.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{startsWith('this is a test', 'this')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "endsWith",
		Body:       endsWith,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns true if a string ends with a substring.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "substring",
					Summary: "The substring to search for.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{endsWith('this is a test', 'test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "regexReplace",
		Body:       regexReplace,
		IsInstance: true,
		Docs: Docs{
			Summary: "Replaces the first occurrence of a substring that matches a regular expression.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to search.",
				},
				{
					Name:    "regex",
					Summary: "The regular expression to search for.",
				},
				{
					Name:    "replacement",
					Summary: "The string to replace the substring with.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this-is-a-test'",
    "test": "{{regexReplace('this is a test', '\s', '-')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "chars",
		Body:       chars,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the characters of a string as an array.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to split into characters.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be ['t', 'h', 'i', 's']",
    "test": "{{chars('this')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "length",
		Body:       length,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the length of a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to get the length of.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 4",
    "test": "{{length('this')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "trim",
		Body:       trim,
		IsInstance: true,
		Docs: Docs{
			Summary: "Removes leading and trailing whitespace from a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to trim.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'this is a test'",
    "test": "{{trim('  this is a test  ')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "title",
		Body:       title,
		IsInstance: true,
		Docs: Docs{
			Summary: "Capitalizes the first letter of each word in a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to capitalize.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'This Is A Test'",
    "test": "{{title('this is a test')}}"
  }
}`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "swapCase",
		Body:       swapCase,
		IsInstance: true,
		Docs: Docs{
			Summary: "Swaps the case of each letter in a string.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to swap the case of.",
				},
			},
			Example: `
{
  "$template": {
    "$comment": "The field below will be 'tHIS IS A TEST'",
    "test": "{{swapCase('This is a test')}}"
  }
}`,
		},
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
