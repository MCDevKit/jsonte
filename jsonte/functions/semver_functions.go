package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func RegisterSemverFunctions() {
	const group = "semver"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Semver functions",
		Summary: "Semver functions are related to working with Semantic Version.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "semver",
		Body:  semverString,
		Docs: Docs{
			Summary: "Creates a new semver object from string, array or major, minor and patch numbers as separate arguments.",
			Arguments: []Argument{
				{
					Name:    "version",
					Summary: "The semver string or array to parse. Optionally the major, minor and patch can be specified as separate arguments.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '1.8.0'",
    "test": "{{semver('1.8.0')}}",
    "$comment1": "The field below will be '1.9.0'",
    "test1": "{{semver([1, 9, 0])}}",
    "$comment2": "The field below will be '1.10.0'",
    "test2": "{{semver(1, 10, 0)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "semver",
		Body:  semverArray,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "semver",
		Body:  semver,
	})
}

func semverString(str types.JsonString) (types.Semver, error) {
	return types.ParseSemverString(str.StringValue())
}

func semverArray(arr types.JsonArray) (types.Semver, error) {
	return types.ParseSemverArray(arr.Unbox().([]interface{}))
}

func semver(major, minor, patch types.JsonNumber) (types.Semver, error) {
	return types.Semver{Major: int(major.IntValue()), Minor: int(minor.IntValue()), Patch: int(patch.IntValue())}, nil
}
