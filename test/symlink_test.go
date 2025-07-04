package test

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
)

func TestWalkDirFollowSymlinks(t *testing.T) {
	dir := t.TempDir()
	realDir := filepath.Join(dir, "real")
	if err := os.Mkdir(realDir, 0o755); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(realDir, "file.txt")
	if err := os.WriteFile(file, []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(realDir, link); err != nil {
		t.Skip("symlinks not supported")
	}
	got := make([]string, 0)
	err := utils.WalkDirFollowSymlinks(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			rel, _ := filepath.Rel(dir, p)
			got = append(got, rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(got)
	expected := []string{
		strings.ReplaceAll("link/file.txt", "/", string(filepath.Separator)),
		strings.ReplaceAll("real/file.txt", "/", string(filepath.Separator)),
	}
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestOpenDirRecursiveSymlink(t *testing.T) {
	dir := t.TempDir()
	realDir := filepath.Join(dir, "real")
	if err := os.Mkdir(realDir, 0o755); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(realDir, "file.txt")
	if err := os.WriteFile(file, []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(realDir, link); err != nil {
		t.Skip("symlinks not supported")
	}
	safeio.Resolver = safeio.DefaultIOResolver
	files, err := safeio.Resolver.OpenDirRecursive(dir)
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(files)
	expected := []string{
		strings.ReplaceAll("link/file.txt", "/", string(filepath.Separator)),
		strings.ReplaceAll("real/file.txt", "/", string(filepath.Separator)),
	}
	sort.Strings(expected)
	if !reflect.DeepEqual(files, expected) {
		t.Fatalf("expected %v, got %v", expected, files)
	}
}
