package test

import (
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"path/filepath"
	"testing"
)

func prepareFS() {
	safeio.Resolver = safeio.CreateFakeFS(
		map[string]interface{}{
			"test.png":           img,
			"test.txt":           "Hello World!",
			"test2.txt":          "Hello World!",
			"dir/test3.txt":      "Hello World!",
			"dir/dir2/test4.txt": "Hello World!",
			"test.json":          `{"test": "Hello World!"}`,
		},
		false,
	)
}

func assertFileList(t *testing.T, eval jsonte.Result, expected []interface{}) {
	t.Helper()
	for i, i2 := range expected {
		expected[i] = filepath.Clean(i2.(string))
	}
	assertArray(t, eval, types.Box(expected).(*types.JsonArray))
}

func TestFileList(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('.').sort()`)
	assertFileList(t, eval, []interface{}{"test.json", "test.png", "test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList2(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('dir').sort()`)
	assertFileList(t, eval, []interface{}{"test3.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList3(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileList('.', '*.txt').sort()`)
	assertFileList(t, eval, []interface{}{"test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList4(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileListRecurse('.').sort()`)
	assertFileList(t, eval, []interface{}{"dir/dir2/test4.txt", "dir/test3.txt", "test.json", "test.png", "test.txt", "test2.txt"})
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileList5(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileListRecurse('.', '*.txt').sort()`)
	assertFileList(t, eval, []interface{}{"dir/dir2/test4.txt", "dir/test3.txt", "test.txt", "test2.txt"})
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
	eval := evaluate(t, `filePath('dir/dir2/test.png')`)
	assertString(t, eval, "dir"+string(filepath.Separator)+"dir2")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFilePathSeparator(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `filePath('dir/dir2/test.png', '/')`)
	assertString(t, eval, "dir/dir2")
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileExists(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileExists('test.txt')`)
	assertBool(t, eval, true)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestFileDoesntExist(t *testing.T) {
	prepareFS()
	eval := evaluate(t, `fileExists('aaaaa.txt')`)
	assertBool(t, eval, false)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestLoadCaching(t *testing.T) {
	prepareFS()
	evaluate(t, `load('test.json')`)
	cached := utils.GetCache("fileCache", "test.json")
	if cached == nil {
		t.Fatalf("cache should contain entry after load")
	}
	// Clean up
	utils.EvictCache("fileCache", "test.json")
	safeio.Resolver = safeio.DefaultIOResolver
}
