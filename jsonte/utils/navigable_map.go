package utils

import (
	"sort"
)

// nmEntry holds a key-value pair for NavigableMap.
// Combining key and value in a single slice halves heap allocations vs. separate key/value slices.
type nmEntry[K comparable, V any] struct {
	key   K
	value V
}

type NavigableMap[K comparable, V any] struct {
	entries []nmEntry[K, V]
	data    map[K]V
	// valid is nil when there are no tombstones (common case).
	// Allocated lazily on first Remove; len(valid) == len(entries) when non-nil.
	valid []bool

	live       int // number of valid entries
	tombstones int // number of invalidated slots

	// Auto-compaction thresholds:
	// trigger when tombstones percentage >= autoCompactPct AND tombstones >= autoCompactMin
	autoCompactPct int // 0..100
	autoCompactMin int
}

// NewNavigableMap creates a new NavigableMap.
func NewNavigableMap[K comparable, V any]() NavigableMap[K, V] {
	return NewNavigableMapWithCapacity[K, V](4)
}

// mapThreshold is the number of live entries at which the hash map is created.
// Below this threshold, TryGet/ContainsKey use a linear scan over the slice.
const mapThreshold = 8

// NewNavigableMapWithCapacity creates a new NavigableMap with preallocated storage.
// The internal hash map is not allocated until the number of entries reaches mapThreshold,
// avoiding heap allocations for small maps.
func NewNavigableMapWithCapacity[K comparable, V any](capacity int) NavigableMap[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return NavigableMap[K, V]{
		entries:        make([]nmEntry[K, V], 0, capacity),
		autoCompactPct: 33,
		autoCompactMin: 128,
	}
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
	target := len(m.entries) + n
	if cap(m.entries) < target {
		newE := make([]nmEntry[K, V], len(m.entries), target)
		copy(newE, m.entries)
		m.entries = newE
	}
	if m.valid != nil && cap(m.valid) < target {
		newB := make([]bool, len(m.valid), target)
		copy(newB, m.valid)
		m.valid = newB
	}
}

func (m *NavigableMap[K, V]) Get(key K) V {
	if m.data != nil {
		return m.data[key]
	}
	if m.valid == nil {
		for i := range m.entries {
			if m.entries[i].key == key {
				return m.entries[i].value
			}
		}
	} else {
		for i := range m.entries {
			if m.entries[i].key == key && m.valid[i] {
				return m.entries[i].value
			}
		}
	}
	var zero V
	return zero
}

func (m *NavigableMap[K, V]) TryGet(key K) (V, bool) {
	if m.data != nil {
		v, ok := m.data[key]
		return v, ok
	}
	if m.valid == nil {
		for i := range m.entries {
			if m.entries[i].key == key {
				return m.entries[i].value, true
			}
		}
	} else {
		for i := range m.entries {
			if m.entries[i].key == key && m.valid[i] {
				return m.entries[i].value, true
			}
		}
	}
	var zero V
	return zero, false
}

func (m *NavigableMap[K, V]) ContainsKey(key K) bool {
	if m.data != nil {
		_, ok := m.data[key]
		return ok
	}
	if m.valid == nil {
		for i := range m.entries {
			if m.entries[i].key == key {
				return true
			}
		}
	} else {
		for i := range m.entries {
			if m.entries[i].key == key && m.valid[i] {
				return true
			}
		}
	}
	return false
}

func (m *NavigableMap[K, V]) Put(key K, value V) {
	if m.data == nil {
		if m.valid == nil {
			for i := range m.entries {
				if m.entries[i].key == key {
					m.entries[i].value = value
					return
				}
			}
		} else {
			for i := range m.entries {
				if m.entries[i].key == key && m.valid[i] {
					m.entries[i].value = value
					return
				}
			}
		}
		m.entries = append(m.entries, nmEntry[K, V]{key: key, value: value})
		if m.valid != nil {
			m.valid = append(m.valid, true)
		}
		m.live++
		// Lazy promotion: create hash map after threshold.
		if m.live >= mapThreshold {
			m.data = make(map[K]V, m.live)
			for i := range m.entries {
				if m.valid == nil || m.valid[i] {
					m.data[m.entries[i].key] = m.entries[i].value
				}
			}
		}
		return
	}
	if _, exists := m.data[key]; !exists {
		m.entries = append(m.entries, nmEntry[K, V]{key: key, value: value})
		if m.valid != nil {
			m.valid = append(m.valid, true)
		}
		m.live++
	} else {
		// Update in-place: find the live slot for this key.
		if m.valid == nil {
			for i := range m.entries {
				if m.entries[i].key == key {
					m.entries[i].value = value
					break
				}
			}
		} else {
			for i := range m.entries {
				if m.entries[i].key == key && m.valid[i] {
					m.entries[i].value = value
					break
				}
			}
		}
	}
	m.data[key] = value
	m.autoCompactMaybe()
}

func (m *NavigableMap[K, V]) Remove(key K) {
	if m == nil {
		return
	}
	if m.data != nil {
		if _, ok := m.data[key]; !ok {
			return
		}
		delete(m.data, key)
	} else {
		found := false
		if m.valid == nil {
			for i := range m.entries {
				if m.entries[i].key == key {
					found = true
					break
				}
			}
		} else {
			for i := range m.entries {
				if m.entries[i].key == key && m.valid[i] {
					found = true
					break
				}
			}
		}
		if !found {
			return
		}
	}
	// Lazily allocate valid on first Remove.
	if m.valid == nil {
		m.valid = make([]bool, len(m.entries))
		for i := range m.valid {
			m.valid[i] = true
		}
	}
	var zero nmEntry[K, V]
	for i := range m.entries {
		if m.entries[i].key == key && m.valid[i] {
			m.entries[i] = zero
			m.valid[i] = false
			m.live--
			m.tombstones++
			break
		}
	}
	m.autoCompactMaybe()
}

func (m *NavigableMap[K, V]) Size() int     { return m.live }
func (m *NavigableMap[K, V]) IsEmpty() bool { return m.live == 0 }

func (m *NavigableMap[K, V]) Clear() {
	m.entries = nil
	m.data = nil
	m.valid = nil
	m.live = 0
	m.tombstones = 0
}

// Reset clears all entries while keeping allocated storage for reuse.
func (m *NavigableMap[K, V]) Reset() {
	if m == nil {
		return
	}
	for k := range m.data {
		delete(m.data, k)
	}
	var zero nmEntry[K, V]
	for i := range m.entries {
		m.entries[i] = zero
	}
	m.entries = m.entries[:0]
	m.valid = nil
	m.live = 0
	m.tombstones = 0
}

func (m *NavigableMap[K, V]) Range(fn func(K, V) bool) {
	if m == nil || fn == nil {
		return
	}
	if m.valid == nil {
		for i := range m.entries {
			if fn(m.entries[i].key, m.entries[i].value) {
				return
			}
		}
		return
	}
	for i := range m.entries {
		if m.valid[i] {
			if fn(m.entries[i].key, m.entries[i].value) {
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
			return true
		}
		return false
	})
	return found
}

// ===== Materialization (prefer Append* to avoid allocs) =====

func (m *NavigableMap[K, V]) Keys() []K {
	out := make([]K, 0, m.live)
	if m.valid == nil {
		for i := range m.entries {
			out = append(out, m.entries[i].key)
		}
		return out
	}
	for i := range m.entries {
		if m.valid[i] {
			out = append(out, m.entries[i].key)
		}
	}
	return out
}

func (m *NavigableMap[K, V]) Values() []V {
	out := make([]V, 0, m.live)
	if m.valid == nil {
		for i := range m.entries {
			out = append(out, m.entries[i].value)
		}
		return out
	}
	for i := range m.entries {
		if m.valid[i] {
			out = append(out, m.entries[i].value)
		}
	}
	return out
}

func (m *NavigableMap[K, V]) AppendKeys(dst []K) []K {
	if m.valid == nil {
		for i := range m.entries {
			dst = append(dst, m.entries[i].key)
		}
		return dst
	}
	for i := range m.entries {
		if m.valid[i] {
			dst = append(dst, m.entries[i].key)
		}
	}
	return dst
}

func (m *NavigableMap[K, V]) AppendValues(dst []V) []V {
	if m.valid == nil {
		for i := range m.entries {
			dst = append(dst, m.entries[i].value)
		}
		return dst
	}
	for i := range m.entries {
		if m.valid[i] {
			dst = append(dst, m.entries[i].value)
		}
	}
	return dst
}

// ===== Sorting =====

// Sort orders valid entries by the provided comparer while keeping stability.
func (m *NavigableMap[K, V]) Sort(comparer func(K, K) int) {
	if comparer == nil || m.live <= 1 {
		return
	}

	if m.valid == nil {
		sort.SliceStable(m.entries, func(a, b int) bool {
			return comparer(m.entries[a].key, m.entries[b].key) < 0
		})
		return
	}

	// Build dense list of valid indices
	idxs := make([]int, 0, m.live)
	for i := range m.entries {
		if m.valid[i] {
			idxs = append(idxs, i)
		}
	}

	// Stable sort by keys at those indices
	sort.SliceStable(idxs, func(a, b int) bool {
		return comparer(m.entries[idxs[a]].key, m.entries[idxs[b]].key) < 0
	})

	// Build compacted & sorted entries
	newEntries := make([]nmEntry[K, V], 0, m.live)
	newValid := make([]bool, 0, m.live)
	for _, i := range idxs {
		newEntries = append(newEntries, m.entries[i])
		newValid = append(newValid, true)
	}

	m.entries = newEntries
	m.valid = newValid
	m.tombstones = 0
}

// ===== Compaction =====

func (m *NavigableMap[K, V]) Compact() {
	if m == nil || m.tombstones == 0 {
		return
	}
	newEntries := make([]nmEntry[K, V], 0, m.live)
	for i := range m.entries {
		if m.valid[i] {
			newEntries = append(newEntries, m.entries[i])
		}
	}
	m.entries = newEntries
	m.valid = nil
	m.tombstones = 0
}

func (m *NavigableMap[K, V]) autoCompactMaybe() {
	if m.autoCompactPct == 0 || m.tombstones == 0 || len(m.entries) == 0 {
		return
	}
	if m.tombstones < m.autoCompactMin {
		return
	}
	if m.tombstones*100/len(m.entries) >= m.autoCompactPct {
		m.Compact()
	}
}
