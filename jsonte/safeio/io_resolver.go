package safeio

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type IOResolver struct {
	Open             func(path string) (ReadAtWriteCloser, error)
	OpenDir          func(path string) ([]string, error)
	OpenDirRecursive func(path string) ([]string, error)
	MkdirAll         func(path string) error
	Remove           func(path string) error
	HttpGet          func(url string) (io.ReadCloser, http.Header, error)
	Create           func(path string) (io.ReadWriteCloser, error)
	Stat             func(path string) (fs.FileInfo, error)
	ExecCommand      func(name string, arg ...string) ([]byte, error)
}

// Resolver Current Resolver used by jsonte to access IO
var Resolver = DefaultIOResolver

// DefaultIOResolver Default IO Resolver, that uses the os package
var DefaultIOResolver = IOResolver{
	Open: func(path string) (ReadAtWriteCloser, error) {
		return os.Open(path)
	},
	OpenDirRecursive: func(path string) ([]string, error) {
		var files []string
		err := utils.WalkDirFollowSymlinks(path, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return burrito.WrapErrorf(err, "Failed to walk path %s", p)
			}
			if d.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(path, p)
			if err != nil {
				return burrito.WrapErrorf(err, "Failed to get relative path of %s", p)
			}
			files = append(files, rel)
			return nil
		})
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to walk path %s", path)
		}
		return files, nil
	},
	OpenDir: func(path string) ([]string, error) {
		var files []string
		dir, err := os.ReadDir(path)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to read dir %s", path)
		}
		for _, file := range dir {
			files = append(files, file.Name())
		}
		return files, nil
	},
	MkdirAll: func(path string) error {
		return os.MkdirAll(path, 0755)
	},
	Remove: func(path string) error {
		return os.RemoveAll(path)
	},
	HttpGet: func(url string) (io.ReadCloser, http.Header, error) {
		open, err := http.Get(url)
		if err != nil {
			return nil, nil, burrito.WrapErrorf(err, "Failed to open url %s", url)
		}
		return open.Body, open.Header, nil
	},
	Create: func(path string) (io.ReadWriteCloser, error) {
		return os.Create(path)
	},
	Stat: func(path string) (fs.FileInfo, error) {
		return os.Stat(path)
	},
	ExecCommand: func(name string, arg ...string) ([]byte, error) {
		return exec.Command(name, arg...).Output()
	},
}

// NoIOResolver Resolver that does not allow any IO
var NoIOResolver = IOResolver{
	Open: func(path string) (ReadAtWriteCloser, error) {
		return nil, burrito.WrappedErrorf("IO is disabled")
	},
	OpenDir: func(path string) ([]string, error) {
		return nil, burrito.WrappedErrorf("IO is disabled")
	},
	MkdirAll: func(path string) error {
		return burrito.WrappedErrorf("IO is disabled")
	},
	Remove: func(path string) error {
		return burrito.WrappedErrorf("IO is disabled")
	},
	HttpGet: func(url string) (io.ReadCloser, http.Header, error) {
		return nil, nil, burrito.WrappedErrorf("IO is disabled")
	},
	Create: func(path string) (io.ReadWriteCloser, error) {
		return nil, burrito.WrappedErrorf("IO is disabled")
	},
	Stat: func(path string) (fs.FileInfo, error) {
		return nil, burrito.WrappedErrorf("IO is disabled")
	},
	ExecCommand: func(name string, arg ...string) ([]byte, error) {
		return nil, burrito.WrappedErrorf("IO is disabled")
	},
}

// CreateFakeFS Creates a fake filesystem from a map of byte slices
func CreateFakeFS(fs map[string]interface{}, withNetwork bool) IOResolver {
	files := make(map[string]*FakeFile)
	for path, data := range fs {
		if data == nil {
			continue
		}
		path = filepath.Clean(path)
		if d, ok := data.([]byte); ok {
			files[path] = CreateFakeFile(d)
		} else if d, ok := data.(string); ok {
			files[path] = CreateFakeFile([]byte(d))
		} else {
			panic("Invalid data type")
		}
	}
	return IOResolver{
		Open: func(path string) (ReadAtWriteCloser, error) {
			path = filepath.Clean(path)
			if data, ok := files[path]; ok {
				return data, nil
			}
			return nil, os.ErrNotExist
		},
		OpenDir: func(path string) ([]string, error) {
			path = filepath.Clean(path)
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
			path = filepath.Clean(path)
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
		MkdirAll: func(path string) error {
			return nil
		},
		Remove: func(path string) error {
			path = filepath.Clean(path)
			delete(files, path)
			return nil
		},
		HttpGet: func(url string) (io.ReadCloser, http.Header, error) {
			if withNetwork {
				return DefaultIOResolver.HttpGet(url)
			}
			return nil, nil, burrito.WrappedErrorf("Network is disabled")
		},
		Create: func(path string) (io.ReadWriteCloser, error) {
			path = filepath.Clean(path)
			files[path] = CreateFakeFile(make([]byte, 0))
			data := files[path]
			return data, nil
		},
		Stat: func(path string) (os.FileInfo, error) {
			path = filepath.Clean(path)
			if data, ok := files[path]; ok {
				return &FakeFileInfo{
					f: data,
					p: path,
				}, nil
			}
			for s := range files {
				if strings.HasPrefix(s, path) {
					return &FakeFileInfo{
						f: nil,
						p: path,
					}, nil
				}
			}
			return nil, os.ErrNotExist
		},
		ExecCommand: func(name string, arg ...string) ([]byte, error) {
			return nil, burrito.WrappedErrorf("Command execution is disabled")
		},
	}
}

type ReadAtWriteCloser interface {
	io.ReadWriteCloser
	io.ReaderAt
}

type FakeFileInfo struct {
	os.FileInfo
	f *FakeFile
	p string
}

func (f *FakeFileInfo) Name() string {
	return f.p
}

func (f *FakeFileInfo) Size() int64 {
	return int64(len(f.f.data))
}

func (f *FakeFileInfo) IsDir() bool {
	return f.f == nil
}

func (f *FakeFileInfo) Mode() os.FileMode {
	return 0644
}

func (f *FakeFileInfo) ModTime() time.Time {
	return time.Now()
}

func (f *FakeFileInfo) Sys() interface{} {
	return nil
}

type FakeFile struct {
	ReadAtWriteCloser
	data    []byte
	pointer int
}

func (f *FakeFile) Close() error {
	return nil
}

func (f *FakeFile) Write(p []byte) (int, error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *FakeFile) Read(p []byte) (n int, err error) {
	if f.pointer >= len(f.data) {
		return 0, io.EOF
	}
	n = copy(p, f.data[f.pointer:])
	f.pointer += n
	return n, nil
}

func (f *FakeFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.pointer = int(offset)
	case io.SeekCurrent:
		f.pointer += int(offset)
	case io.SeekEnd:
		f.pointer = len(f.data) - int(offset)
	}
	return int64(f.pointer), nil
}

func (f *FakeFile) Bytes() []byte {
	return f.data
}

func (f *FakeFile) ReadAt(p []byte, off int64) (n int, err error) {
	if off >= int64(len(f.data)) {
		return 0, io.EOF
	}
	n = copy(p, f.data[off:])
	return n, nil
}

// CreateFakeFile Creates a fake file from a byte slice
func CreateFakeFile(data []byte) *FakeFile {
	return &FakeFile{
		data: data,
	}
}
