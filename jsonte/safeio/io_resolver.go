package safeio

import (
	"bytes"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type IOResolver struct {
	Open             func(path string) (io.ReadCloser, error)
	OpenDir          func(path string) ([]string, error)
	OpenDirRecursive func(path string) ([]string, error)
	IsDir            func(path string) (bool, error)
}

var Resolver IOResolver = DefaultIOResolver

var DefaultIOResolver = IOResolver{
	Open: func(path string) (io.ReadCloser, error) {
		return os.Open(path)
	},
	OpenDirRecursive: func(path string) ([]string, error) {
		var files []string
		err := filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
			if err != nil {
				return utils.WrapErrorf(err, "Failed to walk path %s", p)
			}
			rel, err := filepath.Rel(path, p)
			if err != nil {
				return utils.WrapErrorf(err, "Failed to get relative path of %s", p)
			}
			files = append(files, rel)
			return nil
		})
		if err != nil {
			return nil, utils.WrapErrorf(err, "Failed to walk path %s", path)
		}
		return files, nil
	},
	OpenDir: func(path string) ([]string, error) {
		var files []string
		dir, err := os.ReadDir(path)
		if err != nil {
			return nil, utils.WrapErrorf(err, "Failed to read dir %s", path)
		}
		for _, file := range dir {
			files = append(files, file.Name())
		}
		return files, nil
	},
	IsDir: func(path string) (bool, error) {
		info, err := os.Stat(path)
		if err != nil {
			return false, utils.WrapErrorf(err, "Failed to stat path %s", path)
		}
		return info.IsDir(), nil
	},
}

var NoIOResolver = IOResolver{
	Open: func(path string) (io.ReadCloser, error) {
		return nil, utils.WrappedErrorf("IO is disabled")
	},
	OpenDir: func(path string) ([]string, error) {
		return nil, utils.WrappedErrorf("IO is disabled")
	},
	IsDir: func(path string) (bool, error) {
		return false, utils.WrappedErrorf("IO is disabled")
	},
}

func CreateFakeFS(files map[string][]byte) IOResolver {
	return IOResolver{
		Open: func(path string) (io.ReadCloser, error) {
			if data, ok := files[path]; ok {
				return io.NopCloser(bytes.NewReader(data)), nil
			}
			return nil, utils.WrappedErrorf("File '%s' does not exist", path)
		},
		OpenDir: func(path string) ([]string, error) {
			result := make([]string, 0)
			for p := range files {
				rel, err := filepath.Rel(path, p)
				if err != nil || filepath.Dir(rel) != "." {
					continue
				}
				result = append(result, rel)
			}
			return result, nil
		},
		OpenDirRecursive: func(path string) ([]string, error) {
			result := make([]string, 0)
			for p := range files {
				rel, err := filepath.Rel(path, p)
				if err != nil || strings.HasPrefix(rel, "..") {
					continue
				}
				result = append(result, rel)
			}
			return result, nil
		},
		IsDir: func(path string) (bool, error) {
			for p := range files {
				if strings.HasPrefix(p, path) && p != path {
					return true, nil
				}
			}
			return false, nil
		},
	}
}
