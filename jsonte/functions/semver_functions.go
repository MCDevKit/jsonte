package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
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
			Summary: "Converts a string semver to semver object.",
			Arguments: []Argument{
				{
					Name:    "version",
					Summary: "The semver string to parse.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '1.8.0'",
    "test": "{{semver('1.8.0')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "semver",
		Body:  semverArray,
		Docs: Docs{
			Summary: "Converts an array semver to semver object.",
			Arguments: []Argument{
				{
					Name:    "version",
					Summary: "The semver array to parse.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '1.8.0'",
    "test": "{{semver([1, 8, 0])}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "semver",
		Body:  semver,
		Docs: Docs{
			Summary: "Creates a new semver object.",
			Arguments: []Argument{
				{
					Name:    "major",
					Summary: "Major version.",
				},
				{
					Name:    "minor",
					Summary: "Minor version.",
				},
				{
					Name:    "patch",
					Summary: "Patch version.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '1.8.0'",
    "test": "{{semver(1, 8, 0)}}"
  }
}
</code>`,
		},
	})
}

func semverString(str string) (utils.Semver, error) {
	return utils.ParseSemverString(str)
}

func semverArray(arr []interface{}) (utils.Semver, error) {
	return utils.ParseSemverArray(arr)
}

func semver(major, minor, patch utils.JsonNumber) (utils.Semver, error) {
	return utils.Semver{Major: int(major.IntValue()), Minor: int(minor.IntValue()), Patch: int(patch.IntValue())}, nil
}
