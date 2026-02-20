package utils

import (
	"sort"
)

type NavigableMap[K comparable, V any] struct {
	keys   []K
	values []V
	data   map[K]V
	index  map[K]int
	valid  []bool

	live       int // number of valid entries
	tombstones int // number of invalidated slots

	// Auto-compaction thresholds:
	// trigger when tombstones percentage >= autoCompactPct AND tombstones >= autoCompactMin
	autoCompactPct int // 0..100
	autoCompactMin int
}

// NewNavigableMap creates a new NavigableMap.
func NewNavigableMap[K comparable, V any]() NavigableMap[K, V] {
	return NewNavigableMapWithCapacity[K, V](8)
}

// NewNavigableMapWithCapacity creates a new NavigableMap with preallocated storage.
func NewNavigableMapWithCapacity[K comparable, V any](capacity int) NavigableMap[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	m := NavigableMap[K, V]{
		data:           make(map[K]V, capacity),
		keys:           make([]K, 0, capacity),
		values:         make([]V, 0, capacity),
		index:          make(map[K]int, capacity),
		valid:          make([]bool, 0, capacity),
		autoCompactPct: 33,
		autoCompactMin: 128,
	}
	return m
}

// ToNavigableMap creates a new NavigableMap with given keys and values.
func ToNavigableMap(entries ...interface{}) NavigableMap[string, interface{}] {
	n := NewNavigableMap[string, interface{}]()
	for i := 0; i+1 < len(entries); i += 2 {
		n.Put(entries[i].(string), entries[i+1])
	}
	return n
}

// MapToNavigableMap creates a new NavigableMap from a golang map.
func MapToNavigableMap(entries map[string]interface{}) NavigableMap[string, interface{}] {
	n := NewNavigableMapWithCapacity[string, interface{}](len(entries))
	for k, v := range entries {
		n.Put(k, v)
	}
	return n
}

// SetAutoCompact sets the automatic compaction thresholds.
// percent is the minimum percentage of tombstones (0..100) to trigger compaction.
// minTombstones is the minimum absolute number of tombstones to trigger compaction.
// SetAutoCompact(0, 0) disables automatic compaction.
func (m *NavigableMap[K, V]) SetAutoCompact(percent, minTombstones int) {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	if minTombstones < 0 {
		minTombstones = 0
	}
	m.autoCompactPct = percent
	m.autoCompactMin = minTombstones
}

// Reserve grows internal capacity to at least n additional free slots.
func (m *NavigableMap[K, V]) Reserve(n int) {
	if n <= 0 {
		return
	}
	target := len(m.keys) + n
	if cap(m.keys) < target {
		newK := make([]K, len(m.keys), target)
		copy(newK, m.keys)
		m.keys = newK
	}
	if cap(m.values) < target {
		newV := make([]V, len(m.values), target)
		copy(newV, m.values)
		m.values = newV
	}
	if cap(m.valid) < target {
		newB := make([]bool, len(m.valid), target)
		copy(newB, m.valid)
		m.valid = newB
	}
}

func (m *NavigableMap[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *NavigableMap[K, V]) TryGet(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *NavigableMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *NavigableMap[K, V]) Put(key K, value V) {
	if m.data == nil {
		*m = NewNavigableMapWithCapacity[K, V](8)
	}
	if _, exists := m.data[key]; !exists {
		m.keys = append(m.keys, key)
		m.values = append(m.values, value)
		m.valid = append(m.valid, true)
		m.index[key] = len(m.keys) - 1
		m.live++
	} else {
		if idx, ok := m.index[key]; ok && idx < len(m.values) && m.valid[idx] {
			m.values[idx] = value
		}
	}
	m.data[key] = value
	m.autoCompactMaybe()
}

func (m *NavigableMap[K, V]) Remove(key K) {
	if m == nil {
		return
	}
	if _, ok := m.data[key]; !ok {
		return
	}
	delete(m.data, key)
	if idx, ok := m.index[key]; ok {
		delete(m.index, key)
		var zk K
		var zv V
		if idx < len(m.keys) {
			m.keys[idx] = zk
		}
		if idx < len(m.values) {
			m.values[idx] = zv
		}
		if idx < len(m.valid) && m.valid[idx] {
			m.valid[idx] = false
			m.live--
			m.tombstones++
		}
	}
	m.autoCompactMaybe()
}

func (m *NavigableMap[K, V]) Size() int     { return m.live }
func (m *NavigableMap[K, V]) IsEmpty() bool { return m.live == 0 }

func (m *NavigableMap[K, V]) Clear() {
	m.keys = nil
	m.values = nil
	m.data = nil
	m.index = nil
	m.valid = nil
	m.live = 0
	m.tombstones = 0
}

func (m *NavigableMap[K, V]) Range(fn func(K, V) bool) {
	if m == nil || fn == nil {
		return
	}
	for i, k := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			if fn(k, m.values[i]) {
				return
			}
		}
	}
}

// For compatibility with existing code:
func (m *NavigableMap[K, V]) ForEach(fn func(K, V)) {
	if m == nil || fn == nil {
		return
	}
	m.Range(func(k K, v V) bool { fn(k, v); return false })
}

func (m *NavigableMap[K, V]) ForEachUntil(fn func(K, V) bool) {
	if m == nil || fn == nil {
		return
	}
	m.Range(fn)
}

// ContainsMatchingKey returns true if match(key) is true for any key.
// It uses Range (slice-based, no hashing) and short-circuits on first hit.
func (m *NavigableMap[K, V]) ContainsMatchingKey(match func(K) bool) bool {
	if m == nil || match == nil {
		return false
	}
	found := false
	m.Range(func(k K, _ V) bool {
		if match(k) {
			found = true
			return true // stop iteration
		}
		return false
	})
	return found
}

// ===== Materialization (prefer Append* to avoid allocs) =====

func (m *NavigableMap[K, V]) Keys() []K {
	out := make([]K, 0, m.live)
	for i, k := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			out = append(out, k)
		}
	}
	return out
}

func (m *NavigableMap[K, V]) Values() []V {
	out := make([]V, 0, m.live)
	for i := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			out = append(out, m.values[i])
		}
	}
	return out
}

func (m *NavigableMap[K, V]) AppendKeys(dst []K) []K {
	for i, k := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			dst = append(dst, k)
		}
	}
	return dst
}

func (m *NavigableMap[K, V]) AppendValues(dst []V) []V {
	for i := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			dst = append(dst, m.values[i])
		}
	}
	return dst
}

// ===== Sorting =====

// Sort orders valid entries by the provided comparer while keeping stability.
// It avoids touching tombstones during compare/swap by sorting a dense index list,
// then applies the permutation once to keys/values/valid and rebuilds index.
func (m *NavigableMap[K, V]) Sort(comparer func(K, K) int) {
	if comparer == nil || m.live <= 1 {
		return
	}

	// Build dense list of valid indices
	idxs := make([]int, 0, m.live)
	for i := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			idxs = append(idxs, i)
		}
	}

	// Stable sort by keys at those indices
	sort.SliceStable(idxs, func(a, b int) bool {
		ia, ib := idxs[a], idxs[b]
		return comparer(m.keys[ia], m.keys[ib]) < 0
	})

	// Apply permutation in one pass by building new arrays (dense, compact)
	newKeys := make([]K, 0, m.live)
	newVals := make([]V, 0, m.live)
	newValid := make([]bool, 0, m.live)
	for _, i := range idxs {
		newKeys = append(newKeys, m.keys[i])
		newVals = append(newVals, m.values[i])
		newValid = append(newValid, true)
	}

	// Swap in compacted & sorted storage
	m.keys = newKeys
	m.values = newVals
	m.valid = newValid
	m.tombstones = 0

	// Rebuild index quickly
	if m.index == nil {
		m.index = make(map[K]int, m.live)
	} else {
		for k := range m.index {
			delete(m.index, k)
		}
	}
	for i, k := range m.keys {
		m.index[k] = i
	}
	// m.data is already correct; no change needed
}

// ===== Compaction =====

func (m *NavigableMap[K, V]) Compact() {
	if m == nil || m.tombstones == 0 {
		return
	}
	newKeys := make([]K, 0, m.live)
	newVals := make([]V, 0, m.live)
	newValid := make([]bool, 0, m.live)

	// Rebuild index fresh
	if m.index == nil {
		m.index = make(map[K]int, m.live)
	} else {
		for k := range m.index {
			delete(m.index, k)
		}
	}

	for i, k := range m.keys {
		if i < len(m.valid) && m.valid[i] {
			pos := len(newKeys)
			newKeys = append(newKeys, k)
			newVals = append(newVals, m.values[i])
			newValid = append(newValid, true)
			m.index[k] = pos
		}
	}
	m.keys = newKeys
	m.values = newVals
	m.valid = newValid
	m.tombstones = 0
	// m.data already only contains live entries
}

func (m *NavigableMap[K, V]) autoCompactMaybe() {
	if m.tombstones == 0 || len(m.keys) == 0 {
		return
	}
	if m.tombstones < m.autoCompactMin {
		return
	}
	// percentage check without float math
	if m.tombstones*100/len(m.keys) >= m.autoCompactPct {
		m.Compact()
	}
}
