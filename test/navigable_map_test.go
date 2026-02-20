package test

import (
	"strconv"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/utils"
)

func TestNavigableMapInitialState(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	if !m.IsEmpty() {
		t.Fatalf("expected empty map, got size %d", m.Size())
	}
	if len(m.Keys()) != 0 {
		t.Fatalf("expected no keys, got %d", len(m.Keys()))
	}
}

func TestNavigableMapPutAndGet(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	m.Put("alpha", 1)
	m.Put("beta", 2)

	if v, ok := m.TryGet("alpha"); !ok || v != 1 {
		t.Fatalf("expected (1, true), got (%d, %v)", v, ok)
	}
	if v := m.Get("missing"); v != 0 {
		t.Fatalf("expected zero value for missing key, got %d", v)
	}
	if !m.ContainsKey("beta") {
		t.Fatalf("expected key 'beta' to be present")
	}
}

func TestNavigableMapReplaceValue(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	m.Put("k", 1)
	m.Put("k", 2)
	if v, _ := m.TryGet("k"); v != 2 {
		t.Fatalf("expected value to be replaced with 2, got %d", v)
	}
	if len(m.Keys()) != 1 {
		t.Fatalf("expected single key after replace, got %d", len(m.Keys()))
	}
}

func TestNavigableMapRemove(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Remove("a")
	if m.ContainsKey("a") {
		t.Fatalf("expected key 'a' to be removed")
	}
	if m.Size() != 1 {
		t.Fatalf("expected size 1 after remove, got %d", m.Size())
	}
	keys := m.Keys()
	if len(keys) != 1 || keys[0] != "b" {
		t.Fatalf("unexpected keys after remove: %v", keys)
	}
}

func TestNavigableMapKeysAndValuesOrder(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	for _, k := range []string{"first", "second", "third"} {
		m.Put(k, len(k))
	}
	keys := m.Keys()
	if want := []string{"first", "second", "third"}; len(keys) != len(want) {
		t.Fatalf("expected %d keys, got %d", len(want), len(keys))
	} else {
		for i, k := range want {
			if keys[i] != k {
				t.Fatalf("expected key %q at %d, got %q", k, i, keys[i])
			}
		}
	}
	values := m.Values()
	if len(values) != 3 || values[0] != len("first") || values[1] != len("second") || values[2] != len("third") {
		t.Fatalf("unexpected values order: %v", values)
	}
}

func TestNavigableMapForEachUntil(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	for i := 0; i < 5; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	count := 0
	m.ForEachUntil(func(k string, v int) bool {
		if k != strconv.Itoa(v) {
			t.Fatalf("mismatched key/value pair %q/%d", k, v)
		}
		count++
		return count == 3
	})
	if count != 3 {
		t.Fatalf("expected early stop at 3 iterations, got %d", count)
	}
}

func TestNavigableMapContainsMatchingKey(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	m.Put("foo", 1)
	m.Put("bar", 2)
	if !m.ContainsMatchingKey(func(k string) bool { return k == "bar" }) {
		t.Fatalf("expected ContainsMatchingKey to find 'bar'")
	}
	if m.ContainsMatchingKey(func(k string) bool { return k == "baz" }) {
		t.Fatalf("did not expect to find key 'baz'")
	}
}

func TestNavigableMapClear(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	m.Put("a", 1)
	m.Clear()
	if !m.IsEmpty() {
		t.Fatalf("expected map to be empty after Clear")
	}
	if len(m.Keys()) != 0 {
		t.Fatalf("expected no keys after Clear")
	}
}

func TestNavigableMapSort(t *testing.T) {
	m := utils.NewNavigableMap[string, int]()
	for _, k := range []string{"c", "a", "b"} {
		m.Put(k, int(k[0]))
	}
	m.Sort(func(a, b string) int {
		return int(a[0]) - int(b[0])
	})
	keys := m.Keys()
	if want := []string{"a", "b", "c"}; len(keys) != len(want) {
		t.Fatalf("expected %d keys, got %d", len(want), len(keys))
	} else {
		for i, k := range want {
			if keys[i] != k {
				t.Fatalf("expected key %q at %d, got %q", k, i, keys[i])
			}
		}
	}
}

func TestNavigableMapResizes(t *testing.T) {
	m := utils.NewNavigableMapWithCapacity[string, int](1)
	for i := 0; i < 32; i++ {
		m.Put(strconv.Itoa(i), i)
	}
	if m.Size() != 32 {
		t.Fatalf("expected size 32 after automatic resize, got %d", m.Size())
	}
	for i := 0; i < 32; i++ {
		if v, ok := m.TryGet(strconv.Itoa(i)); !ok || v != i {
			t.Fatalf("unexpected value for key %d: (%d, %v)", i, v, ok)
		}
	}
	m.Remove("0")
	m.Remove("1")
	if m.ContainsKey("0") || m.ContainsKey("1") {
		t.Fatalf("expected removed keys to be absent after resize operations")
	}
	if len(m.Keys()) != 30 {
		t.Fatalf("expected 30 keys after removals, got %d", len(m.Keys()))
	}
}

func TestNavigableMapConversions(t *testing.T) {
	m := utils.ToNavigableMap("x", 1, "y", 2)
	if m.Size() != 2 {
		t.Fatalf("expected size 2 for ToNavigableMap, got %d", m.Size())
	}
	converted := utils.MapToNavigableMap(map[string]interface{}{"a": 1, "b": 2})
	if converted.Size() != 2 {
		t.Fatalf("expected size 2 for MapToNavigableMap, got %d", converted.Size())
	}
}
