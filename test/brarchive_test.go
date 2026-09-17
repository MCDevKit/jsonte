package test

import (
	"bytes"
	"encoding/binary"
	"github.com/MCDevKit/jsonte/jsonte/brarchive"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	json2 "github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"io"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// writeBrarchive writes files into a .brarchive file at p using the Minecraft format.
func writeBrarchive(t *testing.T, p string, files map[string][]byte) {
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, buildBrarchive(files), 0644); err != nil {
		t.Fatal(err)
	}
}

// buildBrarchive encodes files in the Minecraft .brarchive format.
func buildBrarchive(files map[string][]byte) []byte {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	var header, data bytes.Buffer
	header.Write([]byte{0x7d, 0x27, 0x25, 0xb1, 0xa0, 0x52, 0x70, 0x26})
	_ = binary.Write(&header, binary.LittleEndian, uint32(len(names)))
	_ = binary.Write(&header, binary.LittleEndian, uint32(1))
	for _, name := range names {
		entry := make([]byte, 256)
		entry[0] = byte(len(name))
		copy(entry[1:], name)
		if len(files[name]) > 0 {
			binary.LittleEndian.PutUint32(entry[248:], uint32(data.Len()))
			binary.LittleEndian.PutUint32(entry[252:], uint32(len(files[name])))
			data.Write(files[name])
		}
		header.Write(entry)
	}
	return append(header.Bytes(), data.Bytes()...)
}

func TestBrarchiveFakeFS(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(map[string]interface{}{
		"packs/BP/__brarchive/entities.brarchive": buildBrarchive(map[string][]byte{
			"axolotl.json": []byte(`{"name":"axolotl"}`),
		}),
	}, false)
	defer func() { safeio.Resolver = CacheFS }()
	pseudo := "packs/BP/__brarchive/entities.brarchive#/axolotl.json"
	eval := evaluate(t, "load('"+pseudo+"').name")
	assertString(t, eval, "axolotl")
	eval = evaluate(t, "fileExists('"+pseudo+"')")
	assertBool(t, eval, true)
	eval = evaluate(t, "fileExists('packs/BP/__brarchive/entities.brarchive#/missing.json')")
	assertBool(t, eval, false)
}

func TestBrarchiveOpenAndStat(t *testing.T) {
	safeio.Resolver = safeio.DefaultIOResolver
	archive := filepath.Join(t.TempDir(), "entities.brarchive")
	writeBrarchive(t, archive, map[string][]byte{
		"axolotl.json": []byte(`{"name":"axolotl"}`),
		"cow.json":     []byte(`{"name":"cow"}`),
		"empty.json":   {},
	})

	f, err := safeio.Resolver.Open(brarchive.JoinPath(archive, "cow.json"))
	if err != nil {
		t.Fatal(err)
	}
	all, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	_ = f.Close()
	if string(all) != `{"name":"cow"}` {
		t.Fatalf("Unexpected content: %s", all)
	}

	stat, err := safeio.Resolver.Stat(brarchive.JoinPath(archive, "axolotl.json"))
	if err != nil {
		t.Fatal(err)
	}
	if stat.IsDir() || stat.Size() != int64(len(`{"name":"axolotl"}`)) {
		t.Fatalf("Unexpected stat: isDir=%v size=%d", stat.IsDir(), stat.Size())
	}

	f, err = safeio.Resolver.Open(brarchive.JoinPath(archive, "empty.json"))
	if err != nil {
		t.Fatal(err)
	}
	all, err = io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 0 {
		t.Fatalf("Expected empty file, got %s", all)
	}

	_, err = safeio.Resolver.Stat(brarchive.JoinPath(archive, "missing.json"))
	if !os.IsNotExist(err) {
		t.Fatalf("Expected not exist error, got %v", err)
	}
	_, err = safeio.Resolver.Open(brarchive.JoinPath(filepath.Join(t.TempDir(), "missing.brarchive"), "cow.json"))
	if !os.IsNotExist(err) {
		t.Fatalf("Expected not exist error, got %v", err)
	}
}

func TestBrarchiveLoad(t *testing.T) {
	safeio.Resolver = safeio.DefaultIOResolver
	archive := filepath.Join(t.TempDir(), "entities.brarchive")
	writeBrarchive(t, archive, map[string][]byte{
		"axolotl.json": []byte(`{"name":"axolotl"}`),
	})
	pseudo := filepath.ToSlash(brarchive.JoinPath(archive, "axolotl.json"))
	eval := evaluate(t, "load('"+pseudo+"').name")
	assertString(t, eval, "axolotl")
	eval = evaluate(t, "fileExists('"+pseudo+"')")
	assertBool(t, eval, true)
}

func TestGetLatestBPFileFromBrarchive(t *testing.T) {
	safeio.Resolver = safeio.DefaultIOResolver
	// Force the pack lookup to use the cache directory instead of the Minecraft install
	functions.ServerMode = true
	json2.CacheDir = t.TempDir()
	functions.ClearPackCache()
	defer func() {
		functions.ServerMode = false
		json2.CacheDir = ""
		functions.ClearPackCache()
	}()
	pack := filepath.Join(json2.CacheDir, "1.26.50", "BP")
	archive := filepath.Join(pack, "__brarchive", "entities.brarchive")
	writeBrarchive(t, archive, map[string][]byte{
		"axolotl.json": []byte(`{"name":"axolotl"}`),
		"cow.json":     []byte(`{"name":"archived cow"}`),
	})
	nested := filepath.Join(pack, "__brarchive", "entities", "hostile.brarchive")
	writeBrarchive(t, nested, map[string][]byte{
		"zombie.json": []byte(`{"name":"zombie"}`),
	})
	loose := filepath.Join(pack, "entities", "cow.json")
	if err := os.MkdirAll(filepath.Dir(loose), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(loose, []byte(`{"name":"cow"}`), 0644); err != nil {
		t.Fatal(err)
	}

	assertPseudoPath := func(actual, expectedArchive, expectedName string) {
		a, name, ok := brarchive.SplitPath(actual)
		if !ok || filepath.Clean(a) != filepath.Clean(expectedArchive) || name != expectedName {
			t.Fatalf("Expected %s#/%s, got %s", expectedArchive, expectedName, actual)
		}
	}

	eval := evaluate(t, "getLatestBPFile('entities/axolotl.json')")
	assertPseudoPath(eval.Value.StringValue(), archive, "axolotl.json")

	eval = evaluate(t, "load(getLatestBPFile('entities/axolotl.json')).name")
	assertString(t, eval, "axolotl")

	// Loose files take precedence over archived ones
	eval = evaluate(t, "getLatestBPFile('entities/cow.json')")
	if filepath.Clean(eval.Value.StringValue()) != filepath.Clean(loose) {
		t.Fatalf("Expected %s, got %s", loose, eval.Value.StringValue())
	}

	eval = evaluate(t, "getLatestBPFile('entities/missing.json', false)")
	assertNull(t, eval)

	eval = evaluate(t, "listLatestBPFiles('entities')")
	arr, ok := eval.Value.(*types.JsonArray)
	if !ok {
		t.Fatalf("Expected array, got %s", types.ToString(eval.Value))
	}
	if len(arr.Value) != 3 {
		t.Fatalf("Expected 3 files, got %s", types.ToString(eval.Value))
	}
	found := map[string]bool{}
	for _, v := range arr.Value {
		s := v.StringValue()
		if filepath.Clean(s) == filepath.Clean(loose) {
			found["cow"] = true
		} else if a, name, ok := brarchive.SplitPath(s); ok {
			found[filepath.Base(a)+"/"+name] = true
		}
	}
	for _, key := range []string{"cow", "entities.brarchive/axolotl.json", "hostile.brarchive/zombie.json"} {
		if !found[key] {
			t.Fatalf("Missing %s in %s", key, types.ToString(eval.Value))
		}
	}
}
