package bench

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	legacyjson "github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/zeebo/xxh3"
)

var benchABCPayload = buildABCPayload(128)

func BenchmarkJSONABCParseAndAccess(b *testing.B) {
	payload := benchABCPayload
	b.SetBytes(int64(len(payload)))

	b.Run("stdlib_encoding_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, err := stdlibParseAndAccessABC(payload)
			if err != nil {
				b.Fatal(err)
			}
			if v != 123 {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})

	b.Run("legacy_json_file", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, err := legacyParseAndAccessABC(payload)
			if err != nil {
				b.Fatal(err)
			}
			if v != "123" {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})

	b.Run("fast_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, err := fastParseAndAccessABC(payload)
			if err != nil {
				b.Fatal(err)
			}
			if v != "123" {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})
}

func BenchmarkJSONABCAccessOnly(b *testing.B) {
	payload := benchABCPayload

	b.Run("stdlib_encoding_json", func(b *testing.B) {
		var root map[string]interface{}
		if err := stdjson.Unmarshal(payload, &root); err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v, err := stdlibAccessABC(root)
			if err != nil {
				b.Fatal(err)
			}
			if v != 123 {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})

	b.Run("legacy_json_file", func(b *testing.B) {
		raw, err := legacyjson.UnmarshallJSONC(payload)
		if err != nil {
			b.Fatal(err)
		}
		root, ok := raw.(utils.NavigableMap[string, interface{}])
		if !ok {
			b.Fatalf("unexpected root type: %T", raw)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v, err := legacyAccessABC(root)
			if err != nil {
				b.Fatal(err)
			}
			if v != "123" {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})

	b.Run("fast_json", func(b *testing.B) {
		idx, err := utils.Build(payload)
		if err != nil {
			b.Fatal(err)
		}
		_, _ = idx.BuildKeyTable(0, nil)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v, err := fastAccessABC(idx)
			if err != nil {
				b.Fatal(err)
			}
			if v != "123" {
				b.Fatalf("unexpected value: %v", v)
			}
		}
	})
}

func stdlibParseAndAccessABC(payload []byte) (float64, error) {
	var root map[string]interface{}
	if err := stdjson.Unmarshal(payload, &root); err != nil {
		return 0, err
	}
	return stdlibAccessABC(root)
}

func stdlibAccessABC(root map[string]interface{}) (float64, error) {
	a, ok := root["a"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("missing/invalid key a")
	}
	bObj, ok := a["b"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("missing/invalid key b")
	}
	c, ok := bObj["c"].(float64)
	if !ok {
		return 0, fmt.Errorf("missing/invalid key c")
	}
	return c, nil
}

func legacyParseAndAccessABC(payload []byte) (string, error) {
	raw, err := legacyjson.UnmarshallJSONC(payload)
	if err != nil {
		return "", err
	}
	root, ok := raw.(utils.NavigableMap[string, interface{}])
	if !ok {
		return "", fmt.Errorf("unexpected root type: %T", raw)
	}
	return legacyAccessABC(root)
}

func legacyAccessABC(root utils.NavigableMap[string, interface{}]) (string, error) {
	aRaw, ok := root.TryGet("a")
	if !ok {
		return "", fmt.Errorf("missing key a")
	}
	a, ok := aRaw.(utils.NavigableMap[string, interface{}])
	if !ok {
		return "", fmt.Errorf("invalid key a type: %T", aRaw)
	}

	bRaw, ok := a.TryGet("b")
	if !ok {
		return "", fmt.Errorf("missing key b")
	}
	bObj, ok := bRaw.(utils.NavigableMap[string, interface{}])
	if !ok {
		return "", fmt.Errorf("invalid key b type: %T", bRaw)
	}

	cRaw, ok := bObj.TryGet("c")
	if !ok {
		return "", fmt.Errorf("missing key c")
	}
	n, ok := cRaw.(stdjson.Number)
	if !ok {
		return "", fmt.Errorf("invalid key c type: %T", cRaw)
	}
	return n.String(), nil
}

func fastParseAndAccessABC(payload []byte) (string, error) {
	idx, err := utils.Build(payload)
	if err != nil {
		return "", err
	}
	return fastAccessABC(idx)
}

func fastAccessABC(idx *utils.Index) (string, error) {
	if len(idx.Nodes) == 0 || idx.Nodes[0].Kind != utils.Object {
		return "", fmt.Errorf("root is not object")
	}

	aIdx, ok, err := fastFindObjectChildByKey(idx, 0, []byte("a"), nil)
	if err != nil {
		return "", err
	}
	if !ok || idx.Nodes[aIdx].Kind != utils.Object {
		return "", fmt.Errorf("missing/invalid key a")
	}

	bIdx, ok, err := fastFindObjectChildByKey(idx, aIdx, []byte("b"), nil)
	if err != nil {
		return "", err
	}
	if !ok || idx.Nodes[bIdx].Kind != utils.Object {
		return "", fmt.Errorf("missing/invalid key b")
	}

	cIdx, ok, err := fastFindObjectChildByKey(idx, bIdx, []byte("c"), nil)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("missing key c")
	}
	n := idx.Nodes[cIdx]
	if n.Kind != utils.Number {
		return "", fmt.Errorf("invalid key c kind: %v", n.Kind)
	}
	if !bytes.Equal(idx.Src[n.ValLo:n.ValHi], []byte("123")) {
		return "", fmt.Errorf("unexpected key c value")
	}
	return "123", nil
}

func fastFindObjectChildByKey(idx *utils.Index, objIdx int, key []byte, buf []utils.KeyTableEntry) (int, bool, error) {
	table, err := idx.BuildKeyTable(objIdx, buf)
	if err != nil {
		return -1, false, err
	}
	if len(table) == 0 {
		return -1, false, nil
	}

	h := xxh3.Hash(key)
	start := sort.Search(len(table), func(i int) bool {
		return table[i].Hash >= h
	})
	for i := start; i < len(table) && table[i].Hash == h; i++ {
		if fastKeyEquals(idx, table[i], key) {
			return int(table[i].ValIndex), true, nil
		}
	}
	return -1, false, nil
}

func fastKeyEquals(idx *utils.Index, ent utils.KeyTableEntry, key []byte) bool {
	raw := idx.Src[ent.KeyLo:ent.KeyHi]
	if !ent.HasEscape {
		return bytes.Equal(raw, key)
	}
	decoded, err := decodeJSONStringKey(raw)
	if err != nil {
		return false
	}
	return bytes.Equal(decoded, key)
}

func decodeJSONStringKey(raw []byte) ([]byte, error) {
	if bytes.IndexByte(raw, '\\') < 0 {
		return raw, nil
	}
	quoted := make([]byte, 0, len(raw)+2)
	quoted = append(quoted, '"')
	quoted = append(quoted, raw...)
	quoted = append(quoted, '"')
	unquoted, err := strconv.Unquote(string(quoted))
	if err != nil {
		return nil, err
	}
	return []byte(unquoted), nil
}

func buildABCPayload(noise int) []byte {
	var sb strings.Builder
	sb.Grow(16 * 1024)
	sb.WriteByte('{')

	for i := 0; i < noise; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(fmt.Sprintf("\"root_noise_%d\":%d", i, i))
	}

	sb.WriteString(",\"a\":{")
	for i := 0; i < noise; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(fmt.Sprintf("\"a_noise_%d\":%d", i, i))
	}

	sb.WriteString(",\"b\":{")
	for i := 0; i < noise; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(fmt.Sprintf("\"b_noise_%d\":%d", i, i))
	}

	sb.WriteString(",\"c\":123,\"tail\":true}")
	sb.WriteString(",\"a_tail\":false}")
	sb.WriteString(",\"root_tail\":\"done\"}")

	return []byte(sb.String())
}
