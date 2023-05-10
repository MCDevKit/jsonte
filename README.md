# jsonte

The JSON Templating Engine (abbreviated as jsonte) is a streamlined 
templating engine specifically designed for JSON. 
It is optimized for use in Minecraft Bedrock addons, 
which frequently demand extensive JSON that can often be conveniently 
generated using a modest amount of scripting. Notably, jsonte incorporates its own query language.

### [Online demo](https://mcdevkit.com/json)

### [Interactive tutorial](https://mcdevkit.com/tutorial)

### [Docs](https://docs.mcdevkit.com/json-templating-engine/)

## Installation

With regolith:

```powershell
regolith install jsonte
```

From source (requires [Go v1.18 or later](https://golang.org/)):

```powershell
go install github.com/MCDevKit/jsonte@latest
```

## Development

### Prerequisites

- [Go v1.18 or later](https://golang.org/)

### Setup

```powershell
# Clone the repository
git clone https://github.com/MCDevKit/jsonte
cd jsonte
# Install dependencies
go mod vendor
cd scripts
# Setup ANTLR
./setup_env.ps1
# Compile grammar
./compile_antlr.ps1
# Go back to the root directory
cd ..
```

### Building

```powershell
go build github.com/MCDevKit/jsonte
```

### Notes

Before any modifications to `grammar/JsonTemplate.g4` can be made, 
they must be compiled using the `scripts/compile_antlr.ps1` script. 
Due to an existing problem with the parser generation in Go, 
you will need to manually add each new rule to the `visit` method 
found in the `jsonte/expression_visitor.go` file.