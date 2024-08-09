package functions

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gobwas/glob"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const fileCache = "fileCache"

func RegisterFileFunctions() {
	const group = "file"
	RegisterGroup(Group{
		Name:    group,
		Title:   "File functions",
		Summary: "Functions related to files and file paths.",
	})
	utils.CreateCacheBucket(fileCache)
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
    "test": "{{fileList('data', '*.json')}}"
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
		Body:  filePathSeparator,
		Docs: Docs{
			Summary: "Gets the path of a file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file.",
				},
				{
					Name:    "separator",
					Summary: "Separator to use for the path.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will be 'data'",
    "test": "{{filePath('data/file.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "filePath",
		Body:  filePath,
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
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "fileExists",
		Body:     fileExists,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Checks if the file under given path exists.",
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
    "$comment": "The field below will be true if the file test.txt exists",
    "test": "{{fileExists('test.txt')}}"
  }
}
</code>`,
		},
	})
}

func fileExists(s *types.JsonString) (*types.JsonBool, error) {
	_, err := safeio.Resolver.Stat(s.StringValue())
	if err != nil {
		return types.NewBool(false), nil
	}
	return types.NewBool(true), nil
}

func load(s *types.JsonString) (*types.JsonObject, error) {
	cache := utils.GetCache(fileCache, s.StringValue())
	if cache != nil {
		return (*cache).(*types.JsonObject), nil
	} else {
		resolver, err := safeio.Resolver.Open(s.StringValue())
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
		}
		readAll, err := ioutil.ReadAll(resolver)
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to read file %s", s.StringValue())
		}
		err = resolver.Close()
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to close file %s", s.StringValue())
		}
		object, err := types.ParseJsonObject(readAll)
		if err != nil {
			utils.PutCache(fileCache, s.StringValue(), object)
		}
		return object, err
	}
}

func fileList(s *types.JsonString) (*types.JsonArray, error) {
	resolved, err := safeio.Resolver.OpenDir(s.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
	}
	result := make([]types.JsonType, len(resolved))
	for i, file := range resolved {
		result[i] = types.NewString(file)
	}
	return &types.JsonArray{Value: result}, nil
}

func fileListFilter(s *types.JsonString, filter *types.JsonString) (*types.JsonArray, error) {
	compile, err := glob.Compile(filter.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to compile glob %s", filter.StringValue())
	}
	resolved, err := safeio.Resolver.OpenDir(s.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
	}
	result := make([]types.JsonType, 0)
	for _, file := range resolved {
		if compile.Match(file) {
			result = append(result, types.NewString(file))
		}
	}
	return &types.JsonArray{Value: result}, nil
}

func fileListRecurse(s *types.JsonString) (*types.JsonArray, error) {
	resolved, err := safeio.Resolver.OpenDirRecursive(s.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
	}
	result := make([]types.JsonType, len(resolved))
	for i, file := range resolved {
		result[i] = types.NewString(file)
	}
	return &types.JsonArray{Value: result}, nil
}

func fileListRecurseFilter(s *types.JsonString, filter *types.JsonString) (*types.JsonArray, error) {
	compile, err := glob.Compile(filter.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to compile glob %s", filter.StringValue())
	}
	resolved, err := safeio.Resolver.OpenDirRecursive(s.StringValue())
	if err != nil {
		return types.NewJsonArray(), burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
	}
	result := make([]types.JsonType, 0)
	for _, file := range resolved {
		if compile.Match(file) {
			result = append(result, types.NewString(file))
		}
	}
	return &types.JsonArray{Value: result}, nil
}

func fileExtension(s *types.JsonString) *types.JsonString {
	return types.NewString(filepath.Ext(s.StringValue()))
}

func fileName(s *types.JsonString) *types.JsonString {
	return types.NewString(filepath.Base(s.StringValue()))
}

func fileBaseName(s *types.JsonString) *types.JsonString {
	return types.NewString(strings.TrimSuffix(filepath.Base(s.StringValue()), filepath.Ext(s.StringValue())))
}

func filePath(s *types.JsonString) *types.JsonString {
	return types.NewString(filepath.Dir(s.StringValue()))
}

func filePathSeparator(s, separator *types.JsonString) *types.JsonString {
	return types.NewString(strings.ReplaceAll(filepath.Dir(s.StringValue()), string(filepath.Separator), separator.Value))
}

func isDir(s *types.JsonString) (*types.JsonBool, error) {
	stat, err := safeio.Resolver.Stat(s.StringValue())
	if err != nil {
		if os.IsNotExist(err) {
			return types.False(), nil
		}
		return types.False(), burrito.WrapErrorf(err, "Failed to stat path %s", s.StringValue())
	}
	return types.AsBool(stat.IsDir()), nil
}
