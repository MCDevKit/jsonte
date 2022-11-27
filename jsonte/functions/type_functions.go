package functions

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func RegisterTypeFunctions() {
	const group = "type"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Type functions",
		Summary: "Type functions are related to checking the actual type of a variable.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isArray",
		Body:  isArray,
		Docs: Docs{
			Summary: "Returns whether the argument is an array.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isArray([])}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isObject",
		Body:  isObject,
		Docs: Docs{
			Summary: "Returns whether the argument is an object.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$scope": {
    "test": {}
  },
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isObject(test)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isNumber",
		Body:  isNumber,
		Docs: Docs{
			Summary: "Returns whether the argument is a number.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isNumber(1)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isBoolean",
		Body:  isBoolean,
		Docs: Docs{
			Summary: "Returns whether the argument is a boolean.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isBoolean(true)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isSemver",
		Body:  isSemver,
		Docs: Docs{
			Summary: "Returns whether the argument is a semantic version object.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isSemver(semver('1.10.0'))}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "isString",
		Body:  isString,
		Docs: Docs{
			Summary: "Returns whether the argument is a string.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to check.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isString('asd')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "asString",
		Body:  asString,
		Docs: Docs{
			Summary: "Returns the argument as a string.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be \"5\"",
    "test": "{{asString(5)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "asNumber",
		Body:  asNumber,
		Docs: Docs{
			Summary: "Returns the argument as a number.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 5",
    "test": "{{asNumber('5')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "asBoolean",
		Body:  asBool,
		Docs: Docs{
			Summary: "Returns the argument as a boolean.",
			Arguments: []Argument{
				{
					Name:    "object",
					Summary: "The value to convert.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{asBoolean('true')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "parseArray",
		Body:  parseArray,
		Docs: Docs{
			Summary: "Returns the JSON string in the argument as an array.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The JSON string to parse.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be [1, 2, 3]",
    "test": "{{parseArray('[1, 2, 3]')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "parseObject",
		Body:  parseObject,
		Docs: Docs{
			Summary: "Returns the JSON string in the argument as an object.",
			Arguments: []Argument{
				{
					Name:    "string",
					Summary: "The JSON string to parse.",
				},
			},
			Example: `
<code>
{
  "$scope": {
	"test": "{\"a\": 1, \"b\": 2}"
  },
  "$template": {
    "$comment": "The field below will be 1",
    "test": "{{parseObject(test).a}}"
  }
}
</code>`,
		},
	})
}

func isArray(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsArray(obj))
}

func isObject(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsObject(obj))
}

func isNumber(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsNumber(obj))
}

func isBoolean(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsBool(obj))
}

func isSemver(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsSemver(obj))
}

func isString(obj types.JsonType) types.JsonBool {
	return types.NewBool(types.IsString(obj))
}

func asString(obj types.JsonType) types.JsonString {
	return types.AsString(obj)
}

func asNumber(obj types.JsonType) types.JsonNumber {
	return types.AsNumber(obj)
}

func asBool(obj types.JsonType) types.JsonBool {
	return types.AsBool(obj)
}

func parseArray(obj types.JsonString) (types.JsonArray, error) {
	jsonc, err := json.UnmarshallJSONC([]byte(obj.StringValue()))
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to parse string as a valid JSON")
	}
	if !types.IsArray(jsonc) {
		return types.NewJsonArray(), burrito.WrappedError("String is not a JSON array")
	}
	return types.AsArray(jsonc), nil
}

func parseObject(obj types.JsonString) (types.JsonObject, error) {
	jsonc, err := json.UnmarshallJSONC([]byte(obj.StringValue()))
	if err != nil {
		return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to parse string as a valid JSON")
	}
	if !types.IsObject(jsonc) {
		return types.NewJsonObject(), burrito.WrappedError("String is not a JSON object")
	}
	return types.AsObject(jsonc), nil
}
