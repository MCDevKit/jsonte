package functions

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/types"
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this_is_a_test'",
    "test": "{{replace('this is a test', ' ', '_')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this,is,a,test'",
    "test": "{{join(['this', 'is', 'a', 'test'], ',')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{contains('this is a test', 'test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be ['this', 'is', 'a', 'test']",
    "test": "{{split('this is a test', ' ')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 8",
    "test": "{{indexOf('this is a test', 'test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 8",
    "test": "{{lastIndexOf('this is a test', 'test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 2139996864",
    "test": "{{hash('this is a test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'THIS IS A TEST'",
    "test": "{{toUpperCase('this is a test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this is a test'",
    "test": "{{toLowerCase('THIS IS A TEST')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this'",
    "test": "{{substring('this is a test', 0, 4)}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'This is a test'",
    "test": "{{capitalize('this is a test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{startsWith('this is a test', 'this')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{endsWith('this is a test', 'test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this-is-a-test'",
    "test": "{{regexReplace('this is a test', '\s', '-')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be ['t', 'h', 'i', 's']",
    "test": "{{chars('this')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 4",
    "test": "{{length('this')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'this is a test'",
    "test": "{{trim('  this is a test  ')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'This Is A Test'",
    "test": "{{title('this is a test')}}"
  }
}
</code>`,
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
<code>
{
  "$template": {
    "$comment": "The field below will be 'tHIS IS A TEST'",
    "test": "{{swapCase('This is a test')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "number",
		Body:       number,
		IsInstance: true,
		Deprecated: true,
		Docs: Docs{
			Summary: "Converts a string to a number.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1.0",
    "test": "{{=number('1.0')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "format",
		Body:       formatString,
		IsInstance: true,
		Docs: Docs{
			Summary: "Formats a string using the given arguments.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to format.",
				},
				{
					Name:    "args",
					VarArgs: true,
					Summary: "The arguments to use when formatting the string.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 1.0",
    "test": "{{=formatString('%s %s', 'hello', 'world')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "regexMatch",
		Body:       regexMatch,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns an array of matches. Each match is an array containing the submatches (groups) derived from the regular expression pattern. The first submatch in each array is always the complete match found in the target string. If no matches are found, the function returns an empty array.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The string to match.",
				},
				{
					Name:    "args",
					Summary: "The regular expression to match against.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{='hello world'.regexMatch('^hello ')}}"
  }
}
</code>`,
		},
	})
}

func replace(str, old, new types.JsonString) types.JsonString {
	return types.NewString(strings.Replace(str.StringValue(), old.StringValue(), new.StringValue(), -1))
}

func join(strs types.JsonArray, sep types.JsonString) types.JsonString {
	arr := make([]string, len(strs.Value))
	for i, s := range strs.Value {
		arr[i] = types.ToString(s)
	}
	return types.NewString(strings.Join(arr, sep.StringValue()))
}

func stringContains(str, substr types.JsonString) types.JsonBool {
	return types.AsBool(strings.Contains(str.StringValue(), substr.StringValue()))
}

func split(str, sep types.JsonString) types.JsonArray {
	strs := strings.Split(str.StringValue(), sep.StringValue())
	arr := make([]types.JsonType, len(strs))
	for i, s := range strs {
		arr[i] = types.NewString(s)
	}
	return types.JsonArray{Value: arr}
}

func stringIndexOf(str, substr types.JsonString) types.JsonNumber {
	return types.AsNumber(strings.Index(str.StringValue(), substr.StringValue()))
}

func stringLastIndexOf(str, substr types.JsonString) types.JsonNumber {
	return types.AsNumber(strings.LastIndex(str.StringValue(), substr.StringValue()))
}

func hash(str types.JsonString) (types.JsonNumber, error) {
	a := fnv.New32a()
	_, err := a.Write([]byte(str.StringValue()))
	if err != nil {
		return types.AsNumber(0), burrito.WrapErrorf(err, "Failed to hash string")
	}
	return types.AsNumber(a.Sum32()), nil
}

func toUpperCase(str types.JsonString) types.JsonString {
	return types.NewString(cases.Upper(language.Und).String(str.StringValue()))
}

func toLowerCase(str types.JsonString) types.JsonString {
	return types.NewString(cases.Lower(language.Und).String(str.StringValue()))
}

func substring(str types.JsonString, start, end types.JsonNumber) types.JsonString {
	return types.NewString(str.StringValue()[start.IntValue():end.IntValue()])
}

func substringFrom(str types.JsonString, start types.JsonNumber) types.JsonString {
	return types.NewString(str.StringValue()[start.IntValue():])
}

func captialize(str types.JsonString) types.JsonString {
	return types.NewString(cases.Upper(language.Und).String(str.StringValue()[:1]) + cases.Lower(language.Und).String(str.StringValue()[1:]))
}

func title(str types.JsonString) types.JsonString {
	return types.NewString(cases.Title(language.Und).String(str.StringValue()))
}

func swapCase(s types.JsonString) types.JsonString {
	str := s.StringValue()
	for i, c := range str {
		if unicode.IsUpper(c) {
			str = str[:i] + strings.ToLower(string(c)) + str[i+1:]
		} else if unicode.IsLower(c) {
			str = str[:i] + strings.ToUpper(string(c)) + str[i+1:]
		}
	}
	return types.NewString(str)
}

func startsWith(str, substr types.JsonString) types.JsonBool {
	return types.AsBool(strings.HasPrefix(str.StringValue(), substr.StringValue()))
}

func endsWith(str, substr types.JsonString) types.JsonBool {
	return types.AsBool(strings.HasSuffix(str.StringValue(), substr.StringValue()))
}

func regexReplace(str, pattern, repl types.JsonString) (types.JsonString, error) {
	compile, err := regexp.Compile(pattern.StringValue())
	if err != nil {
		return types.EmptyString, burrito.WrapErrorf(err, "Failed to compile regex pattern")
	}
	return types.NewString(compile.ReplaceAllString(str.StringValue(), repl.StringValue())), nil
}

func chars(str types.JsonString) types.JsonArray {
	arr := make([]types.JsonType, len([]rune(str.StringValue())))
	for i, c := range str.StringValue() {
		arr[i] = types.NewString(string(c))
	}
	return types.JsonArray{Value: arr}
}

func length(str types.JsonString) types.JsonNumber {
	return types.AsNumber(len([]rune(str.StringValue())))
}

func trim(str types.JsonString) types.JsonString {
	return types.NewString(strings.Trim(str.StringValue(), " \t\n\r"))
}

var floatPattern = regexp.MustCompile(`^[+-]?([0-9]+([.][0-9]+)?)$`)

func number(str types.JsonString) (types.JsonNumber, error) {
	if str.StringValue() == "" {
		return types.AsNumber(0), nil
	}
	if !floatPattern.MatchString(str.StringValue()) {
		return types.AsNumber(0), burrito.WrappedErrorf("String '%s' is not a valid number", str.StringValue())
	}
	return types.AsNumber(str.StringValue()), nil
}

func formatString(str types.JsonString, args ...types.JsonType) (types.JsonString, error) {
	if len(args) == 0 {
		return str, nil
	}
	fmtArgs := make([]interface{}, len(args))
	for i, arg := range args {
		switch arg.(type) {
		case types.JsonNumber:
			fmtArgs[i] = arg.(types.JsonNumber).FloatValue()
		case types.JsonBool:
			fmtArgs[i] = arg.BoolValue()
		default:
			fmtArgs[i] = arg.StringValue()
		}
	}

	return types.NewString(fmt.Sprintf(str.StringValue(), fmtArgs...)), nil
}

func regexMatch(str, pattern types.JsonString) (types.JsonArray, error) {
	compile, err := regexp.Compile(pattern.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to compile regex pattern")
	}
	return types.AsArray(compile.FindAllStringSubmatch(str.StringValue(), -1)), nil
}
