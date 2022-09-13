package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"testing"
)

func prepareFS() {
	safeio.Resolver = safeio.CreateFakeFS(
		map[string][]byte{
			"test.png":           img,
			"test.txt":           []byte("Hello World!"),
			"test2.txt":          []byte("Hello World!"),
			"dir/test3.txt":      []byte("Hello World!"),
			"dir/dir2/test4.txt": []byte("Hello World!"),
			"test.json":          []byte(`{"test": "Hello World!"}`),
		},
	)
}

func TestFileList(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('.').sort()`)
	assertArray(t, eval, utils.JsonArray{"test.json", "test.png", "test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList2(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('dir').sort()`)
	assertArray(t, eval, utils.JsonArray{"test3.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList3(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('.', '*.txt').sort()`)
	assertArray(t, eval, utils.JsonArray{"test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList4(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileListRecurse('.').sort()`)
	assertArray(t, eval, utils.JsonArray{"dir\\dir2\\test4.txt", "dir\\test3.txt", "test.json", "test.png", "test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList5(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileListRecurse('.', '*.txt').sort()`)
	assertArray(t, eval, utils.JsonArray{"dir\\dir2\\test4.txt", "dir\\test3.txt", "test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestLoadFile(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `load('test.json').test`)
	assertString(t, eval, "Hello World!")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileExtension(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileExtension('test.png')`)
	assertString(t, eval, ".png")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileExtension2(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileExtension('test')`)
	assertString(t, eval, "")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileBaseName(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileBaseName('test.png')`)
	assertString(t, eval, "test")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileName(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileName('dir/test.png')`)
	assertString(t, eval, "test.png")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFilePath(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `filePath('dir/test.png')`)
	assertString(t, eval, "dir")
	safeio.Resolver = safeio.DefaultIOResolver
}
