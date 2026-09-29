package utils

import (
	"math"
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

	// live is the number of valid entries. The number of invalidated slots
	// (tombstones) is always len(entries) - live, so it is not stored.
	live int32

	// Auto-compaction thresholds:
	// trigger when tombstones percentage >= autoCompactPct AND tombstones >= autoCompactMin.
	// Kept small so they pack next to live and the header stays compact.
	autoCompactPct uint8  // 0..100; 0 disables auto-compaction
	autoCompactMin uint16 // saturates at math.MaxUint16

	// inline backs entries for small maps initialized via InitInline,
	// avoiding a separate heap allocation. Unused otherwise.
	inline [NavigableMapInlineCap]nmEntry[K, V]
}

// NavigableMapInlineCap is the number of entries InitInline can store without a heap allocation.
const NavigableMapInlineCap = 2

// InitInline points entry storage at the map's inline array, so maps that stay
// at or below nmInlineCap entries need no separate heap allocation.
// Only call it through a pointer whose pointee will not be copied or moved
// (e.g. a freshly heap-allocated struct embedding the map): entries aliases
// the inline array, so copying the map by value leaves the copy's slice
// pointing into the original.
func (m *NavigableMap[K, V]) InitInline() {
	m.entries = m.inline[:0]
	m.autoCompactPct = defaultAutoCompactPct
	m.autoCompactMin = defaultAutoCompactMin
}

// Default auto-compaction thresholds used by the constructors.
const (
	defaultAutoCompactPct = 33
	defaultAutoCompactMin = 128
)

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
		autoCompactPct: defaultAutoCompactPct,
		autoCompactMin: defaultAutoCompactMin,
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
	if minTombstones > math.MaxUint16 {
		minTombstones = math.MaxUint16
	}
	m.autoCompactPct = uint8(percent)
	m.autoCompactMin = uint16(minTombstones)
}

// tombstones returns the number of invalidated slots.
func (m *NavigableMap[K, V]) tombstones() int {
	return len(m.entries) - int(m.live)
}

// CloneInto overwrites dst with a compacted copy of m, mapping every value
// through f (values are copied as-is when f is nil). Entry order is kept.
// dst gets the default auto-compaction thresholds. Small results use dst's
// inline storage, so dst must not be copied or moved afterwards (e.g. a
// freshly heap-allocated struct embedding the map).
func (m *NavigableMap[K, V]) CloneInto(dst *NavigableMap[K, V], f func(V) V) {
	live := int(m.live)
	if live <= NavigableMapInlineCap {
		dst.inline = [NavigableMapInlineCap]nmEntry[K, V]{}
		dst.entries = dst.inline[:0]
	} else {
		dst.entries = make([]nmEntry[K, V], 0, live)
	}
	for i := range m.entries {
		if m.valid != nil && !m.valid[i] {
			continue
		}
		e := m.entries[i]
		if f != nil {
			e.value = f(e.value)
		}
		dst.entries = append(dst.entries, e)
	}
	dst.valid = nil
	dst.live = m.live
	dst.autoCompactPct = defaultAutoCompactPct
	dst.autoCompactMin = defaultAutoCompactMin
	dst.data = nil
	if live >= mapThreshold {
		dst.data = make(map[K]V, live)
		for i := range dst.entries {
			dst.data[dst.entries[i].key] = dst.entries[i].value
		}
	}
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
			m.data = make(map[K]V, max(int(m.live), cap(m.entries)))
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
			break
		}
	}
	m.autoCompactMaybe()
}

func (m *NavigableMap[K, V]) Size() int     { return int(m.live) }
func (m *NavigableMap[K, V]) IsEmpty() bool { return m.live == 0 }

func (m *NavigableMap[K, V]) Clear() {
	m.entries = nil
	m.data = nil
	m.valid = nil
	m.live = 0
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

// At returns the key and value of the i-th live entry in insertion order.
// ok is false when i is out of range.
func (m *NavigableMap[K, V]) At(i int) (key K, value V, ok bool) {
	if i < 0 || i >= int(m.live) {
		return key, value, false
	}
	if m.valid == nil {
		e := &m.entries[i]
		return e.key, e.value, true
	}
	for j := range m.entries {
		if m.valid[j] {
			if i == 0 {
				return m.entries[j].key, m.entries[j].value, true
			}
			i--
		}
	}
	return key, value, false
}

// ===== Materialization (prefer Append* to avoid allocs) =====

func (m *NavigableMap[K, V]) Keys() []K {
	out := make([]K, 0, int(m.live))
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
	out := make([]V, 0, int(m.live))
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
	idxs := make([]int, 0, int(m.live))
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
	newEntries := make([]nmEntry[K, V], 0, int(m.live))
	newValid := make([]bool, 0, int(m.live))
	for _, i := range idxs {
		newEntries = append(newEntries, m.entries[i])
		newValid = append(newValid, true)
	}

	m.entries = newEntries
	m.valid = newValid
}

// ===== Compaction =====

func (m *NavigableMap[K, V]) Compact() {
	if m == nil || m.tombstones() == 0 {
		return
	}
	newEntries := make([]nmEntry[K, V], 0, int(m.live))
	for i := range m.entries {
		if m.valid[i] {
			newEntries = append(newEntries, m.entries[i])
		}
	}
	m.entries = newEntries
	m.valid = nil
}

func (m *NavigableMap[K, V]) autoCompactMaybe() {
	tombstones := m.tombstones()
	if m.autoCompactPct == 0 || tombstones == 0 || len(m.entries) == 0 {
		return
	}
	if tombstones < int(m.autoCompactMin) {
		return
	}
	if tombstones*100/len(m.entries) >= int(m.autoCompactPct) {
		m.Compact()
	}
}
