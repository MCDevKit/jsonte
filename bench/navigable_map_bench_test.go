package bench

import (
	"testing"

	"github.com/MCDevKit/jsonte/jsonte/utils"
)

const benchmarkMapSize = 2048

var benchmarkKeys = func() []string {
	keys := make([]string, benchmarkMapSize)
	for i := range keys {
		keys[i] = fastItoa(i)
	}
	return keys
}()

// fastItoa avoids fmt to keep the benchmark focused on map performance.
func fastItoa(v int) string {
	if v == 0 {
		return "0"
	}
	digits := [20]byte{}
	pos := len(digits)
	for v > 0 {
		pos--
		digits[pos] = byte('0' + v%10)
		v /= 10
	}
	return string(digits[pos:])
}

func BenchmarkNavigableMap(b *testing.B) {
	b.Run("Put", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
			for _, k := range benchmarkKeys {
				m.Put(k, n)
			}
		}
	})

	b.Run("PutReplace", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 0)
		}
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			for _, k := range benchmarkKeys {
				m.Put(k, n)
			}
		}
	})

	b.Run("Get", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		var sink int
		for n := 0; n < b.N; n++ {
			for _, k := range benchmarkKeys {
				sink += m.Get(k)
			}
		}
		_ = sink
	})

	b.Run("TryGet", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		missing := "missing"
		b.ResetTimer()
		b.ReportAllocs()
		var sink bool
		for n := 0; n < b.N; n++ {
			for _, k := range benchmarkKeys {
				_, sink = m.TryGet(k)
			}
			_, sink = m.TryGet(missing)
		}
		_ = sink
	})

	b.Run("Remove", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
			for _, k := range benchmarkKeys {
				m.Put(k, n)
			}
			for _, k := range benchmarkKeys {
				m.Remove(k)
			}
		}
	})

	b.Run("Keys", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			if len(m.Keys()) != benchmarkMapSize {
				b.Fatalf("unexpected key count: %d", len(m.Keys()))
			}
		}
	})

	b.Run("Values", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		var sink int
		for n := 0; n < b.N; n++ {
			vals := m.Values()
			for _, v := range vals {
				sink += v
			}
		}
		_ = sink
	})

	b.Run("ForEachUntil", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			count := 0
			m.ForEachUntil(func(_ string, _ int) bool {
				count++
				return false
			})
			if count != benchmarkMapSize {
				b.Fatalf("unexpected iteration count: %d", count)
			}
		}
	})

	b.Run("ContainsKey", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		missing := "missing"
		b.ResetTimer()
		b.ReportAllocs()
		var sink bool
		for n := 0; n < b.N; n++ {
			for _, k := range benchmarkKeys {
				sink = m.ContainsKey(k)
			}
			sink = m.ContainsKey(missing)
		}
		_ = sink
	})

	b.Run("ContainsMatchingKey", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		predicate := func(target string) func(string) bool {
			return func(key string) bool {
				return key == target
			}
		}
		b.ResetTimer()
		b.ReportAllocs()
		var sink bool
		for n := 0; n < b.N; n++ {
			sink = m.ContainsMatchingKey(predicate("nonexistent"))
			sink = m.ContainsMatchingKey(predicate(benchmarkKeys[len(benchmarkKeys)/2]))
		}
		_ = sink
	})

	b.Run("Sort", func(b *testing.B) {
		m := utils.NewNavigableMapWithCapacity[string, int](benchmarkMapSize)
		for _, k := range benchmarkKeys {
			m.Put(k, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m.Sort(func(a, b string) int {
				if a == b {
					return 0
				}
				if a < b {
					return -1
				}
				return 1
			})
		}
	})

	b.Run("Resize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := utils.NewNavigableMapWithCapacity[string, int](1)
			for _, k := range benchmarkKeys {
				m.Put(k, n)
			}
		}
	})
}
