package functions

import (
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
  "$template": {
    "$comment": "The field below will be true",
    "test": "{{isObject({})}}"
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
