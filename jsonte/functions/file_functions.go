package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gobwas/glob"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func RegisterFileFunctions() {
	const group = "file"
	RegisterGroup(Group{
		Name:    group,
		Title:   "File functions",
		Summary: "Functions related to files and file paths.",
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "load",
		Body:     load,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Loads a JSON file from the given path.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file to load.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be an object from the file data.json",
    "test": "{{load('data.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "fileList",
		Body:     fileList,
		IsUnsafe: true, Docs: Docs{
			Summary: "Lists all files in a directory.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory to list.",
				},
				{
					Name:     "filter",
					Summary:  "A glob filter to match files against.",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be an array of all files with .json extension in the data directory",
    "test": "{{fileList('data', "*.json")}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "fileList",
		Body:     fileListFilter,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "fileListRecurse",
		Body:     fileListRecurse,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Lists all files in a directory, recursively.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory to list.",
				},
				{
					Name:     "filter",
					Summary:  "A glob filter to match files against.",
					Optional: true,
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be an array of all files with .json extension in the data directory and its subdirectories",
    "test": "{{fileListRecurse('data', "*.json")}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "fileListRecurse",
		Body:     fileListRecurseFilter,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "fileExtension",
		Body:  fileExtension,
		Docs: Docs{
			Summary: "Gets the extension of a file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be '.json'",
    "test": "{{fileExtension('data.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "fileBaseName",
		Body:  fileBaseName,
		Docs: Docs{
			Summary: "Gets the base name of a file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 'data'",
    "test": "{{fileBaseName('data.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "fileName",
		Body:  fileName,
		Docs: Docs{
			Summary: "Gets the name of a file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be the name of the file data.json",
    "test": "{{fileName('dir/data.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "filePath",
		Body:  filePath,
		Docs: Docs{
			Summary: "Gets the path of a file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 'data'",
    "test": "{{filePath('data.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "isDir",
		Body:     isDir,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Checks if the given path is a directory.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be true if the path is a directory",
    "test": "{{isDir('data')}}"
  }
}
</code>`,
		},
	})
}

func load(s string) (utils.JsonObject, error) {
	resolver, err := safeio.Resolver.Open(s)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to resolve path %s", s)
	}
	readAll, err := ioutil.ReadAll(resolver)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to read file %s", s)
	}
	err = resolver.Close()
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to close file %s", s)
	}
	return utils.ParseJson(readAll)
}

func fileList(s string) (utils.JsonArray, error) {
	resolved, err := safeio.Resolver.OpenDir(s)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to resolve path %s", s)
	}
	result := make(utils.JsonArray, len(resolved))
	for i, file := range resolved {
		result[i] = file
	}
	return result, nil
}

func fileListFilter(s string, filter string) (utils.JsonArray, error) {
	compile, err := glob.Compile(filter)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to compile glob %s", filter)
	}
	resolved, err := safeio.Resolver.OpenDir(s)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to resolve path %s", s)
	}
	result := make(utils.JsonArray, 0)
	for _, file := range resolved {
		if compile.Match(file) {
			result = append(result, file)
		}
	}
	return result, nil
}

func fileListRecurse(s string) (utils.JsonArray, error) {
	resolved, err := safeio.Resolver.OpenDirRecursive(s)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to resolve path %s", s)
	}
	result := make(utils.JsonArray, len(resolved))
	for i, file := range resolved {
		result[i] = file
	}
	return result, nil
}

func fileListRecurseFilter(s string, filter string) (utils.JsonArray, error) {
	compile, err := glob.Compile(filter)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to compile glob %s", filter)
	}
	resolved, err := safeio.Resolver.OpenDirRecursive(s)
	if err != nil {
		return nil, utils.WrapErrorf(err, "Failed to resolve path %s", s)
	}
	result := make(utils.JsonArray, 0)
	for _, file := range resolved {
		if compile.Match(file) {
			result = append(result, file)
		}
	}
	return result, nil
}

func fileExtension(s string) string {
	return filepath.Ext(s)
}

func fileName(s string) string {
	return filepath.Base(s)
}

func fileBaseName(s string) string {
	return strings.TrimSuffix(filepath.Base(s), filepath.Ext(s))
}

func filePath(s string) string {
	return filepath.Dir(s)
}

func isDir(s string) (bool, error) {
	stat, err := safeio.Resolver.Stat(s)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, utils.WrapErrorf(err, "Failed to stat path %s", s)
	}
	return stat.IsDir(), nil
}
