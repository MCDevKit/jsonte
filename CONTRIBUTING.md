# Contributing

When contributing to this repository, it's worth discussing first the change you wish 
to make via issue, email, Discord (stirante#0001), or any other method with the owners 
of this repository before making a change.

## All code changes happen through pull requests

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've added code that should be documented, add documentation.
4. Try your best to follow the [effective go](https://go.dev/doc/effective_go) guidelines as much as possible.
5. Be sure to test your modifications.
6. Format your code with `go fmt`.
7. Write a good commit message.
8. Issue that pull request!

## Any contributions you make will be under the MIT Software License

In short, when you submit code changes, your submissions are understood to be
under the same [MIT License](http://choosealicense.com/licenses/mit/) that
covers the project. Feel free to contact the maintainers if that's a concern.

## Report bugs using Github's [issues](https://github.com/MCDevKit/jsonte/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new
issue](https://github.com/MCDevKit/jsonte/issues/new); it's that easy!

## How to Implement Common Changes

### Instructions for Adding a New Function Group

1. Create a new file in the `jsonte/functions` directory. The file should be named `<group>_functions.go`.
2. Write a function that registers all functions within that group. This function should be named `Register<group>Functions`.
3. Use the `RegisterGroup` function to register the group. You will need to input the following structure:
   - `Name` - The group's name
   - `Title` - The group's title
   - `Summary` - A brief description of the group
4. Register all functions to the newly created group (refer to the [Adding a New Function](#instructions-for-adding-a-new-function) section for more details).
5. Include a call to the newly created function in the `Init` function, located in `jsonte/functions/function_definition.go`.
6. Write tests for all the functions in the new group. These should be placed in a new file named `test/<group>_functions_test.go`.

### Instructions for Adding a New Function

1. Create a new function within the `jsonte/functions` directory. The appropriate file for this function will be determined by the group your function belongs to, e.g., `array_functions.go`.
2. Ensure the parameters and the return value of your function conform to the following types:
   - `utils.JsonArray` - array
   - `utils.JsonObject` - object
   - `utils.JsonString` - string
   - `utils.JsonNumber` - number
   - `utils.JsonBool` - boolean
   - `utils.JsonLambda` - lambda
   - `[]utils.JsonType` - varargs of any type
3. The function may also return an additional value of the `error` type. If this value is not `nil`, the error will be returned to the user.
4. Register the function by calling the `RegisterFunction` function within `Register<group>Functions`. You will need to input a structure with the following fields:
   - `Group` - The group's name
   - `Name` - The function's name
   - `Body` - The function itself
   - `IsInstance` - A flag indicating whether this function can be invoked on an instance of an object (currently only supported for `array` and `string` types)
   - `IsUnsafe` - A flag indicating whether this function should be marked as unsafe and consequently disabled in safe mode (applicable to file manipulation, network access, etc.)
   - `Docs` - The function's documentation
5. Write a test for your new function. This should be placed in the `test` directory, in the file associated with your function's group, e.g., `array_functions_test.go`.

### Instructions for Modifying the Grammar

1. Execute the `scripts/setup_env.ps1` script to configure the environment. This only needs to be done once. The script will establish the environment in `C:\antlr` and add it to the path.
2. Make necessary changes to the `grammar/JsonTemplate.g4` file (refer to the [official Antlr4 documentation](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)).
3. Run the `scripts/compile_antlr.ps1` script to compile the parser.
4. Implement the new grammar by writing corresponding code in the `jsonte/expression_visitor.go` file.
5. When creating a new rule, provide an interface implementation in the `Visit<rule>` function and include a call to this function in the `Visit` function, located in the `jsonte/expression_visitor.go` file.
6. Write tests for your grammar changes. These should be included in the `test/eval_test.go` file.
7. Execute the `go test` command to verify that all modifications are working as expected.