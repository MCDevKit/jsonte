# jsonte

JSON Templating Engine (jsonte for short) is a simple templating engine for JSON. 
It is designed to be used in Minecraft Bedrock addons, which require large amounts of JSON, 
that is often easy to generate with a bit of script. 
It includes its own query language.

## [Online demo](https://mcdevkit.com/json)

## [Interactive tutorial](https://mcdevkit.com/tutorial)

## [Docs](https://docs.mcdevkit.com/json-templating-engine/)

## Installation

With regolith:

```powershell
regolith install json_templating_engine
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

All changes to `grammar/JsonTemplate.g4` must be compiled with `scripts/compile_antlr.ps1` before they can be used.
Due to the issue with generating the parser in Go, each new rule should be added to method `visit` in `jsonte/expression_visitor.go` manually.
