package test

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"io/ioutil"
	"testing"
)

func TestFakeOpen(t *testing.T) {
	expected := "Hello World!"
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt": []byte(expected),
	})
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt":           []byte("Hello World!"),
		"test2.txt":          []byte("Hello World!"),
		"dir/test3.txt":      []byte("Hello World!"),
		"dir/dir2/test4.txt": []byte("Hello World!"),
	})
	open, err := safeio.Resolver.OpenDir(".")
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

func TestFakeOpenDir2(t *testing.T) {
	expected := []string{"test3.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt":           []byte("Hello World!"),
		"test2.txt":          []byte("Hello World!"),
		"dir/test3.txt":      []byte("Hello World!"),
		"dir/dir2/test4.txt": []byte("Hello World!"),
	})
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
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt":           []byte("Hello World!"),
		"test2.txt":          []byte("Hello World!"),
		"dir/test3.txt":      []byte("Hello World!"),
		"dir/dir2/test4.txt": []byte("Hello World!"),
	})
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
	expected := []string{"test3.txt", "dir2\\test4.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt":           []byte("Hello World!"),
		"test2.txt":          []byte("Hello World!"),
		"dir/test3.txt":      []byte("Hello World!"),
		"dir/dir2/test4.txt": []byte("Hello World!"),
	})
	open, err := safeio.Resolver.OpenDirRecursive("dir")
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

func TestFakeOpenDirRecursive2(t *testing.T) {
	expected := []string{"test.txt", "test2.txt", "dir\\test3.txt", "dir\\dir2\\test4.txt"}
	safeio.Resolver = safeio.CreateFakeFS(map[string][]byte{
		"test.txt":           []byte("Hello World!"),
		"test2.txt":          []byte("Hello World!"),
		"dir/test3.txt":      []byte("Hello World!"),
		"dir/dir2/test4.txt": []byte("Hello World!"),
	})
	open, err := safeio.Resolver.OpenDirRecursive(".")
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
