package bench

import (
	stdjson "encoding/json"
	"fmt"
	"strings"
	"testing"

	legacyjson "github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/utils"
)

var benchIterValuesPayload = buildIterValuesPayload(96)

func BenchmarkJSONIterateAllValuesParseAndIterate(b *testing.B) {
	payload := benchIterValuesPayload
	b.SetBytes(int64(len(payload)))

	b.Run("stdlib_encoding_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var root interface{}
			if err := stdjson.Unmarshal(payload, &root); err != nil {
				b.Fatal(err)
			}
			count := stdlibCountLeafValues(root)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})

	b.Run("legacy_json_file", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			raw, err := legacyjson.UnmarshallJSONC(payload)
			if err != nil {
				b.Fatal(err)
			}
			count := legacyCountLeafValues(raw)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})

	b.Run("fast_json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx, err := utils.Build(payload)
			if err != nil {
				b.Fatal(err)
			}
			count := fastCountLeafValues(idx)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})
}

func BenchmarkJSONIterateAllValuesIterateOnly(b *testing.B) {
	payload := benchIterValuesPayload

	b.Run("stdlib_encoding_json", func(b *testing.B) {
		var root interface{}
		if err := stdjson.Unmarshal(payload, &root); err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := stdlibCountLeafValues(root)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})

	b.Run("legacy_json_file", func(b *testing.B) {
		raw, err := legacyjson.UnmarshallJSONC(payload)
		if err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := legacyCountLeafValues(raw)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})

	b.Run("fast_json", func(b *testing.B) {
		idx, err := utils.Build(payload)
		if err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := fastCountLeafValues(idx)
			if count == 0 {
				b.Fatal("unexpected empty count")
			}
		}
	})
}

func stdlibCountLeafValues(v interface{}) int {
	switch t := v.(type) {
	case map[string]interface{}:
		total := 0
		for _, child := range t {
			total += stdlibCountLeafValues(child)
		}
		return total
	case []interface{}:
		total := 0
		for _, child := range t {
			total += stdlibCountLeafValues(child)
		}
		return total
	default:
		return 1
	}
}

func legacyCountLeafValues(v interface{}) int {
	switch t := v.(type) {
	case utils.NavigableMap[string, interface{}]:
		total := 0
		t.ForEach(func(_ string, child interface{}) {
			total += legacyCountLeafValues(child)
		})
		return total
	case []interface{}:
		total := 0
		for _, child := range t {
			total += legacyCountLeafValues(child)
		}
		return total
	default:
		return 1
	}
}

func fastCountLeafValues(idx *utils.Index) int {
	if idx == nil || len(idx.Nodes) == 0 {
		return 0
	}
	total := 0
	for _, n := range idx.Nodes {
		switch n.Kind {
		case utils.Object, utils.Array:
		default:
			total++
		}
	}
	return total
}

func buildIterValuesPayload(n int) []byte {
	var sb strings.Builder
	sb.Grow(64 * 1024)
	sb.WriteString("{\"root\":{")
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(fmt.Sprintf("\"k%d\":{\"num\":%d,\"flag\":%t,\"arr\":[%d,%d,{\"s\":\"v%d\",\"nil\":null}]}", i, i, i%2 == 0, i, i+1, i))
	}
	sb.WriteString("},\"tail\":[1,2,3,{\"x\":\"y\"}],\"ok\":true}")
	return []byte(sb.String())
}
