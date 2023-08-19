package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func RegisterJsonPathFunctions() {
	const group = "JsonPath"
	RegisterGroup(Group{
		Name:    group,
		Title:   "JsonPath functions",
		Summary: "JsonPath functions are related to working with Json paths.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "JsonPath",
		Body:  parseJsonPath,
		Docs: Docs{
			Summary: "Creates a new JsonPath object from a string.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to parse.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '#/test/path[1]'",
    "test": "{{JsonPath('#/test/path[1]')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:      group,
		Name:       "parent",
		Body:       pathParent,
		IsInstance: true,
		Docs: Docs{
			Summary: "Returns the parent of the given path.",
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '#/test/path'",
    "test": "{{JsonPath('#/test/path[1]').parent()}}"
  }
}
</code>`,
		},
	})
}

func parseJsonPath(str types.JsonString) (types.JsonPath, error) {
	return types.ParseJsonPath(str.StringValue())
}

func pathParent(path types.JsonPath) types.JsonPath {
	return path.Parent()
}
