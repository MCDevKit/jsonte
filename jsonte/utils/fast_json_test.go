package utils

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/zeebo/xxh3"
)

func TestFastJSONChildrenAreDirectOnly(t *testing.T) {
	src := []byte(`{"a":{"x":1},"b":[{"y":2},3]}`)
	idx, err := Build(src)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	rootChildren := idx.Children(0)
	if len(rootChildren) != 2 {
		t.Fatalf("root children count = %d, want 2", len(rootChildren))
	}
	if idx.Nodes[rootChildren[0]].Kind != Object {
		t.Fatalf("root child[0] kind = %v, want Object", idx.Nodes[rootChildren[0]].Kind)
	}
	if idx.Nodes[rootChildren[1]].Kind != Array {
		t.Fatalf("root child[1] kind = %v, want Array", idx.Nodes[rootChildren[1]].Kind)
	}

	aChildren := idx.Children(rootChildren[0])
	if len(aChildren) != 1 || idx.Nodes[aChildren[0]].Kind != Number {
		t.Fatalf("object a direct children mismatch: %#v", aChildren)
	}

	bChildren := idx.Children(rootChildren[1])
	if len(bChildren) != 2 {
		t.Fatalf("array b children count = %d, want 2", len(bChildren))
	}
	if idx.Nodes[bChildren[0]].Kind != Object || idx.Nodes[bChildren[1]].Kind != Number {
		t.Fatalf("array b children kinds mismatch: got %v, %v", idx.Nodes[bChildren[0]].Kind, idx.Nodes[bChildren[1]].Kind)
	}

	nestedChildren := idx.Children(bChildren[0])
	if len(nestedChildren) != 1 || idx.Nodes[nestedChildren[0]].Kind != Number {
		t.Fatalf("nested object children mismatch: %#v", nestedChildren)
	}
}

func TestFastJSONBuildKeyTableUsesDirectChildrenAndUnescapedHash(t *testing.T) {
	src := []byte(`{"a":{"nested":1},"c":2,"\u0064":3,"\u0024sys":4}`)
	idx, err := Build(src)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	table, err := idx.BuildKeyTable(0, nil)
	if err != nil {
		t.Fatalf("BuildKeyTable failed: %v", err)
	}
	if len(table) != 4 {
		t.Fatalf("table size = %d, want 4", len(table))
	}

	entries := make(map[string]KeyTableEntry, len(table))
	for _, ent := range table {
		raw := idx.Src[ent.KeyLo:ent.KeyHi]
		key, err := decodeKey(raw)
		if err != nil {
			t.Fatalf("decode key failed: %v", err)
		}
		entries[key] = ent
		if key == "nested" {
			t.Fatalf("unexpected nested key present in root key table")
		}
	}

	if _, ok := entries["a"]; !ok {
		t.Fatalf("missing key a")
	}
	if _, ok := entries["c"]; !ok {
		t.Fatalf("missing key c")
	}

	d, ok := entries["d"]
	if !ok {
		t.Fatalf("missing key d")
	}
	if !d.HasEscape {
		t.Fatalf("key d should have HasEscape=true")
	}
	if d.Hash != xxh3.Hash([]byte("d")) {
		t.Fatalf("hash for d mismatch")
	}

	sys, ok := entries["$sys"]
	if !ok {
		t.Fatalf("missing key $sys")
	}
	if !sys.HasEscape {
		t.Fatalf("key $sys should have HasEscape=true")
	}
	if !sys.HasDollar {
		t.Fatalf("key $sys should have HasDollar=true")
	}
	if sys.Hash != xxh3.Hash([]byte("$sys")) {
		t.Fatalf("hash for $sys mismatch")
	}

	table2, err := idx.BuildKeyTable(0, make([]KeyTableEntry, 0, 8))
	if err != nil {
		t.Fatalf("BuildKeyTable (cached) failed: %v", err)
	}
	if len(table2) != len(table) {
		t.Fatalf("cached table size mismatch")
	}
	for i := range table {
		if table[i] != table2[i] {
			t.Fatalf("cached table differs at %d", i)
		}
	}
}

func decodeKey(raw []byte) (string, error) {
	if bytes.IndexByte(raw, '\\') < 0 {
		return string(raw), nil
	}
	quoted := make([]byte, 0, len(raw)+2)
	quoted = append(quoted, '"')
	quoted = append(quoted, raw...)
	quoted = append(quoted, '"')
	return strconv.Unquote(string(quoted))
}
