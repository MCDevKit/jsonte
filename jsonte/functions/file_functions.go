package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gobwas/glob"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func RegisterFileFunctions() {
	RegisterFunction(JsonFunction{
		Name:     "load",
		Body:     load,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "fileList",
		Body:     fileList,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "fileList",
		Body:     fileListFilter,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "fileListRecurse",
		Body:     fileListRecurse,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "fileListRecurse",
		Body:     fileListRecurseFilter,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name: "fileExtension",
		Body: fileExtension,
	})
	RegisterFunction(JsonFunction{
		Name: "fileBaseName",
		Body: fileBaseName,
	})
	RegisterFunction(JsonFunction{
		Name: "fileName",
		Body: fileName,
	})
	RegisterFunction(JsonFunction{
		Name: "filePath",
		Body: filePath,
	})
	RegisterFunction(JsonFunction{
		Name:     "isDir",
		Body:     isDir,
		IsUnsafe: true,
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
	return result, err
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
	return result, err
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
	return result, err
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
	return result, err
}

func fileExtension(s string) (string, error) {
	return filepath.Ext(s), nil
}

func fileName(s string) (string, error) {
	return filepath.Base(s), nil
}

func fileBaseName(s string) (string, error) {
	return strings.TrimSuffix(filepath.Base(s), filepath.Ext(s)), nil
}

func filePath(s string) (string, error) {
	return filepath.Dir(s), nil
}

func isDir(s string) (bool, error) {
	return safeio.Resolver.IsDir(s)
}
