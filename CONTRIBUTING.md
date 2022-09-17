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

## Common changes and how to make them

### Adding a new function group

1. Create a new file in `jsonte/functions` directory named `<group>_functions.go`.
2. Create a function for registering all functions in that group. It should be named `Register<group>Functions`.
3. Register all functions by to that group (refer to the next section for more info).
4. Add a call to that function in `Init` function in `jsonte/functions/function_definition.go`.
5. Add tests for all functions in that group in a new file called `test/<group>_functions_test.go`.

### Adding a new function

1. Create a new function in `jsonte/functions` directory in file, in which group your function will be e.g. `array_functions.go`.
2. Make sure, that parameters to this function as well as return value are of following types:
    - `utils.JsonArray` - array
    - `utils.JsonObject` - object
    - `string` - string
    - `utils.JsonNumber` - number
    - `bool` - boolean
    - `utils.JsonLambda` - lambda
3. The function can also return an additional value of type `error`. If it's not `nil`, the error will be returned to the user.
4. Register the function by calling `RegisterFunction` function in `Register<group>Functions` with struct of following fields:
    - `Name` - name of the function
    - `Body` - function itself
    - `IsInstance` - could this function be called on an instance of an object (currently supported only for `array` and `string` types) 
    - `IsUnsafe` - should this function be marked as unsafe and thus disabled in safe mode (file manipulation, network access, etc.)
    - `Docs` - The docs for the function
5. Add a test for your function in `test` directory in file, in which group your function will be e.g. `array_functions_test.go`.

### Making changes to the grammar

1. Run script `scripts/setup_env.ps1` to set up the environment (you need to run it only once, it will set it up in `C:\antlr` and add it to the path).
2. Make changes to `grammar/JsonTemplate.g4` file (refer to the [docs](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)).
3. Run script `scripts/compile_antlr.ps1` to generate the parser.
4. Write code handling the new grammar in `jsonte/expression_visitor.go` file.
5. When making a new rule, add create an interface implementation in function `Visit<rule>` and add a call to it in `Visit` function in `jsonte/expression_visitor.go` file.
6. Add tests for your changes in `test/eval_test.go` file.
7. Run `go test` to make sure, that everything works.