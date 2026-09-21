package functions

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gobwas/glob"
)

const fileCache = "fileCache"
const textFileCache = "textFileCache"

func RegisterFileFunctions() {
	const group = "file"
	RegisterGroup(Group{
		Name:    group,
		Title:   "File functions",
		Summary: "Functions related to files and file paths.",
	})
	utils.CreateCacheBucket(fileCache)
	utils.CreateCacheBucket(textFileCache)
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
		Name:     "loadText",
		Body:     loadText,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Loads a text file from the given path.",
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
    "$comment": "The field below will be a string from the file test.molang",
    "test": "{{loadText('test.molang')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "loadMolang",
		Body:     loadMolangWithParams,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Loads a Molang file, evaluates its `#{...}` expressions and returns the minified Molang. The file is processed when the function is first called, so it can be used in `.jsonte` scripts. Results are cached for each path and parameters.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the Molang file to load.",
				},
				{
					Name:     "params",
					Optional: true,
					Summary:  "An object whose fields are available as variables in the file's expressions, in addition to the scope.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "rect.molang contains 'v.i = #{index};'. The field below will be 'v.i=3;'",
    "test": "{{loadMolang('rect.molang', {'index': 3})}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "loadMolang",
		Body:     loadMolang,
		IsUnsafe: true,
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

// CacheTextFile pre-populates the text file cache so that loadText returns
// the given content for the given path without reading from disk.
func CacheTextFile(path, content string) {
	utils.PutCache(textFileCache, path, types.NewString(content))
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
		readAll, err := readAllFile(resolver)
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to read file %s", s.StringValue())
		}
		err = resolver.Close()
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to close file %s", s.StringValue())
		}
		object, err := types.ParseJsonObject(readAll)
		if err == nil {
			utils.PutCache(fileCache, s.StringValue(), object)
		}
		return object, err
	}
}

// MolangLoader evaluates the jsonte expressions in Molang source, with params as an extra scope, and returns
// the minified Molang.
type MolangLoader func(source string, params *types.JsonObject) (string, error)

var (
	molangLoader     MolangLoader
	scopeMolangFiles map[string]struct{}
	molangCache      sync.Map
)

// SetMolangLoader sets how loadMolang processes files and clears its cache. loadText returns scopeFiles
// processed with no parameters instead of their raw content.
func SetMolangLoader(loader MolangLoader, scopeFiles []string) {
	molangLoader = loader
	scopeMolangFiles = make(map[string]struct{}, len(scopeFiles))
	for _, f := range scopeFiles {
		scopeMolangFiles[molangPathKey(f)] = struct{}{}
	}
	molangCache.Clear()
}

func molangPathKey(p string) string {
	return filepath.ToSlash(filepath.Clean(p))
}

func loadMolang(s *types.JsonString) (*types.JsonString, error) {
	return loadMolangWithParams(s, types.NewJsonObject())
}

func loadMolangWithParams(s *types.JsonString, params *types.JsonObject) (*types.JsonString, error) {
	if molangLoader == nil {
		return types.EmptyString, burrito.WrappedErrorf("Molang files can only be loaded when compiling with a scope")
	}
	key := molangPathKey(s.StringValue()) + "\x00" + params.StringValue()
	if cached, ok := molangCache.Load(key); ok {
		return cached.(*types.JsonString), nil
	}
	resolver, err := safeio.Resolver.Open(s.StringValue())
	if err != nil {
		return types.EmptyString, burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
	}
	source, err := readAllFile(resolver)
	if err != nil {
		return types.EmptyString, burrito.WrapErrorf(err, "Failed to read file %s", s.StringValue())
	}
	err = resolver.Close()
	if err != nil {
		return types.EmptyString, burrito.WrapErrorf(err, "Failed to close file %s", s.StringValue())
	}
	name := s.StringValue()
	if !params.IsEmpty() {
		name += " " + params.StringValue()
	}
	release, err := enterMolangLoad(key, name)
	if err != nil {
		return types.EmptyString, err
	}
	defer release()
	// Concurrent first calls may process the same file twice; the result is the same, so that is harmless.
	output, err := molangLoader(string(source), params)
	if err != nil {
		return types.EmptyString, burrito.WrapErrorf(err, "Failed to process the molang file %s", s.StringValue())
	}
	result := types.NewString(output)
	molangCache.Store(key, result)
	return result, nil
}

// maxMolangLoadDepth limits nested loadMolang calls, which catches endless recursion with changing parameters.
const maxMolangLoadDepth = 100

type molangLoad struct {
	key, name string
}

// molangLoadChains holds the loadMolang calls in progress for each goroutine, innermost last.
var molangLoadChains sync.Map

// enterMolangLoad records that the goroutine starts processing a molang file and returns a function that
// records the end. Nested loadMolang calls run synchronously in the goroutine of the outer call, so the
// per-goroutine chain shows exactly which files are loading each other.
func enterMolangLoad(key, name string) (release func(), err error) {
	id := goroutineID()
	var chain []molangLoad
	if c, ok := molangLoadChains.Load(id); ok {
		chain = c.([]molangLoad)
	}
	for i, load := range chain {
		if load.key == key {
			return nil, burrito.WrappedErrorf("Circular loadMolang: %s", formatMolangChain(chain[i:], name))
		}
	}
	if len(chain) >= maxMolangLoadDepth {
		return nil, burrito.WrappedErrorf("loadMolang is nested more than %d levels deep: %s", maxMolangLoadDepth, formatMolangChain(chain[len(chain)-3:], name))
	}
	chain = append(chain, molangLoad{key: key, name: name})
	molangLoadChains.Store(id, chain)
	return func() {
		if len(chain) == 1 {
			molangLoadChains.Delete(id)
		} else {
			molangLoadChains.Store(id, chain[:len(chain)-1])
		}
	}, nil
}

func formatMolangChain(chain []molangLoad, next string) string {
	names := make([]string, 0, len(chain)+1)
	for _, load := range chain {
		names = append(names, load.name)
	}
	return strings.Join(append(names, next), " -> ")
}

// goroutineID returns the id of the current goroutine, parsed from the "goroutine N [...]" stack header.
func goroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	header := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	id, _ := strconv.ParseUint(header[:strings.IndexByte(header, ' ')], 10, 64)
	return id
}

func loadText(s *types.JsonString) (*types.JsonString, error) {
	cache := utils.GetCache(textFileCache, s.StringValue())
	if cache != nil {
		return (*cache).(*types.JsonString), nil
	} else {
		if _, ok := scopeMolangFiles[molangPathKey(s.StringValue())]; ok && molangLoader != nil {
			return loadMolang(s)
		}
		resolver, err := safeio.Resolver.Open(s.StringValue())
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "Failed to resolve path %s", s.StringValue())
		}
		readAll, err := readAllFile(resolver)
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "Failed to read file %s", s.StringValue())
		}
		err = resolver.Close()
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "Failed to close file %s", s.StringValue())
		}
		str := types.NewString(string(readAll))
		utils.PutCache(textFileCache, s.StringValue(), str)
		return str, err
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

// readAllFile reads the whole file. When the size is known up front (e.g. entries of a .brarchive),
// the buffer is allocated once instead of being grown by io.ReadAll.
func readAllFile(r io.Reader) ([]byte, error) {
	if s, ok := r.(interface{ Size() int64 }); ok {
		buf := make([]byte, s.Size())
		_, err := io.ReadFull(r, buf)
		return buf, err
	}
	return io.ReadAll(r)
}
