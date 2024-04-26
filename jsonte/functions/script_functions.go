package functions

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func RegisterScriptFunctions() {
	const group = "script"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Script functions",
		Summary: "Script functions are related to working with scripts.",
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "println",
		Body:  scriptPrintln,
		Docs: Docs{
			Summary: "Prints the given value to the console and appends a newline.",
			Arguments: []Argument{
				{
					Name:    "value",
					Summary: "The value to print.",
				},
			},
			Example: `
<code>
println('Hello world!')
</code>`,
		},
	})
}

func scriptPrintln(str types.JsonType) (*types.JsonNull, error) {
	fmt.Println(str.StringValue())
	return types.Null, nil
}
