package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"
)

func TestFakeOpen(t *testing.T) {
	expected := "Hello World!"
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt": expected,
	}, false)
	open, err := safeio.Resolver.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		t.Fatal(err)
	}
	if string(all) != expected {
		t.Fatalf("Invalid content. Expected '%s', got '%s'", expected, string(all))
	}
}

func TestFakeOpenDir(t *testing.T) {
	expected := []string{"test.txt", "test2.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	open, err := safeio.Resolver.OpenDir(".")
	if err != nil {
		t.Fatal(err)
	}
	if len(open) != len(expected) {
		t.Fatalf("Invalid length. Expected %d, got %d", len(expected), len(open))
	}
	sort.SliceStable(open, func(i, j int) bool {
		return open[i] < open[j]
	})
	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	for i, file := range open {
		if file != expected[i] {
			t.Fatalf("Invalid file. Expected %s, got %s", expected[i], file)
		}
	}
}

func TestFakeOpenDir2(t *testing.T) {
	expected := []string{"test3.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	open, err := safeio.Resolver.OpenDir("dir")
	if err != nil {
		t.Fatal(err)
	}
	if len(open) != len(expected) {
		t.Fatalf("Invalid length. Expected %d, got %d", len(expected), len(open))
	}
	for i, file := range open {
		if file != expected[i] {
			t.Fatalf("Invalid file. Expected %s, got %s", expected[i], file)
		}
	}
}

func TestFakeOpenDir3(t *testing.T) {
	expected := []string{"test4.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	open, err := safeio.Resolver.OpenDir("dir/dir2")
	if err != nil {
		t.Fatal(err)
	}
	if len(open) != len(expected) {
		t.Fatalf("Invalid length. Expected %d, got %d", len(expected), len(open))
	}
	for i, file := range open {
		if file != expected[i] {
			t.Fatalf("Invalid file. Expected %s, got %s", expected[i], file)
		}
	}
}

func TestFakeOpenDirRecursive(t *testing.T) {
	expected := []string{"test3.txt", "dir2/test4.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	open, err := safeio.Resolver.OpenDirRecursive("dir")
	if err != nil {
		t.Fatal(err)
	}
	if len(open) != len(expected) {
		t.Fatalf("Invalid length. Expected %d, got %d", len(expected), len(open))
	}
	sort.SliceStable(open, func(i, j int) bool {
		return open[i] < open[j]
	})
	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	for i, file := range open {
		if file != filepath.Clean(expected[i]) {
			t.Fatalf("Invalid file. Expected %s, got %s", expected[i], file)
		}
	}
}

func TestFakeOpenDirRecursive2(t *testing.T) {
	expected := []string{"test.txt", "test2.txt", "dir/test3.txt", "dir/dir2/test4.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	open, err := safeio.Resolver.OpenDirRecursive(".")
	if err != nil {
		t.Fatal(err)
	}
	if len(open) != len(expected) {
		t.Fatalf("Invalid length. Expected %d, got %d", len(expected), len(open))
	}
	sort.SliceStable(open, func(i, j int) bool {
		return open[i] < open[j]
	})
	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	for i, file := range open {
		if file != filepath.Clean(expected[i]) {
			t.Fatalf("Invalid file. Expected %s, got %s", expected[i], file)
		}
	}
}

func TestFakeCreate(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	out, err := safeio.Resolver.Create("test5.txt")
	if err != nil {
		t.Fatal(err)
	}
	n, err := out.Write([]byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}
	if n != 12 {
		t.Fatalf("Invalid written length. Expected %d, got %d", 12, n)
	}
	in, err := safeio.Resolver.Open("test5.txt")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(in)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "Hello World!" {
		t.Fatalf("Invalid data. Expected %s, got %s", "Hello World!", string(data))
	}
}

func TestFakeCreate2(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	out, err := safeio.Resolver.Create("dir/test5.txt")
	if err != nil {
		t.Fatal(err)
	}
	n, err := out.Write([]byte("Hello World!"))
	if err != nil {
		t.Fatal(err)
	}
	if n != 12 {
		t.Fatalf("Invalid written length. Expected %d, got %d", 12, n)
	}
	in, err := safeio.Resolver.Open("dir/test5.txt")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(in)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "Hello World!" {
		t.Fatalf("Invalid data. Expected %s, got %s", "Hello World!", string(data))
	}
}

func TestFakeRemove(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"test.txt":           "Hello World!",
		"test2.txt":          "Hello World!",
		"dir/test3.txt":      "Hello World!",
		"dir/dir2/test4.txt": "Hello World!",
	}, false)
	err := safeio.Resolver.Remove("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = safeio.Resolver.Open("test.txt")
	if err == nil {
		t.Fatal("File should not exist")
	}
}
