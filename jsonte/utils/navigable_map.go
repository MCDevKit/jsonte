package utils

import "sort"

// NavigableMap is a map that can be navigated in order of keys. It is not thread-safe.
type NavigableMap[K comparable, V any] struct {
	keys  []K
	data  map[K]V
	index map[K]int
}

// NewNavigableMap creates a new NavigableMap.
func NewNavigableMap[K comparable, V any]() NavigableMap[K, V] {
	return NewNavigableMapWithCapacity[K, V](0)
}

// NewNavigableMapWithCapacity creates a new NavigableMap with preallocated storage.
func NewNavigableMapWithCapacity[K comparable, V any](capacity int) NavigableMap[K, V] {
	if capacity < 0 {
		capacity = 0
	}
	return NavigableMap[K, V]{
		data:  make(map[K]V, capacity),
		keys:  make([]K, 0, capacity),
		index: make(map[K]int, capacity),
	}
}

// ToNavigableMap creates a new NavigableMap with given keys and values.
func ToNavigableMap(entries ...interface{}) NavigableMap[string, interface{}] {
	n := NavigableMap[string, interface{}]{
		data:  map[string]interface{}{},
		keys:  []string{},
		index: map[string]int{},
	}
	for i := 0; i < len(entries); i += 2 {
		n.Put(entries[i].(string), entries[i+1])
	}
	return n
}

// MapToNavigableMap creates a new NavigableMap from a golang map.
func MapToNavigableMap(entries map[string]interface{}) NavigableMap[string, interface{}] {
	n := NavigableMap[string, interface{}]{
		data:  map[string]interface{}{},
		keys:  []string{},
		index: map[string]int{},
	}
	for key, value := range entries {
		n.Put(key, value)
	}
	return n
}

// Get returns the value associated with the key.
func (m *NavigableMap[K, V]) Get(key K) V {
	return m.data[key]
}

// TryGet returns the value and whether it exists.
func (m *NavigableMap[K, V]) TryGet(key K) (V, bool) {
	value, ok := m.data[key]
	return value, ok
}

// Put puts the value associated with the key.
func (m *NavigableMap[K, V]) Put(key K, value V) {
	if m.data == nil {
		m.data = make(map[K]V)
	}
	if m.index == nil {
		m.index = make(map[K]int)
	}
	if _, exists := m.data[key]; !exists {
		m.keys = append(m.keys, key)
		m.index[key] = len(m.keys) - 1
	}
	m.data[key] = value
}

// Remove removes the value associated with the key.
func (m *NavigableMap[K, V]) Remove(key K) {
	if !m.ContainsKey(key) {
		return
	}
	delete(m.data, key)
	if idx, ok := m.index[key]; ok {
		delete(m.index, key)
		var zero K
		m.keys[idx] = zero
	}
}

// ContainsKey returns true if the key exists.
func (m *NavigableMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.data[key]
	return ok
}

// ContainsMatchingKey returns true if the matching key exists.
func (m *NavigableMap[K, V]) ContainsMatchingKey(matchFunc func(K) bool) bool {
	if m == nil || matchFunc == nil {
		return false
	}
	found := false
	m.ForEachUntil(func(key K, _ V) bool {
		if matchFunc(key) {
			found = true
			return true
		}
		return false
	})
	return found
}

// Keys returns the keys in order.
func (m *NavigableMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for idx, k := range m.keys {
		if i, ok := m.index[k]; ok && i == idx {
			keys = append(keys, k)
		}
	}
	return keys
}

// Values returns the values in order of keys.
func (m *NavigableMap[K, V]) Values() []V {
	var values []V
	for _, k := range m.Keys() {
		values = append(values, m.data[k])
	}
	return values
}

// ForEach executes fn for every key/value pair that is currently present.
func (m *NavigableMap[K, V]) ForEach(fn func(K, V)) {
	if m == nil || fn == nil {
		return
	}
	m.ForEachUntil(func(key K, value V) bool {
		fn(key, value)
		return false
	})
}

// ForEachUntil executes fn for every key/value pair until fn returns true.
func (m *NavigableMap[K, V]) ForEachUntil(fn func(K, V) bool) {
	if m == nil || fn == nil {
		return
	}
	for idx, k := range m.keys {
		if i, ok := m.index[k]; ok && i == idx {
			if fn(k, m.data[k]) {
				return
			}
		}
	}
}

// Size returns the size of the map.
func (m *NavigableMap[K, V]) Size() int {
	return len(m.data)
}

// IsEmpty returns true if the map is empty.
func (m *NavigableMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Clear clears the map.
func (m *NavigableMap[K, V]) Clear() {
	m.keys = []K{}
	m.data = map[K]V{}
	m.index = map[K]int{}
}

// Sort sorts the map with the specified comparator.
func (m *NavigableMap[K, V]) Sort(comparer func(K, K) int) {
	sort.SliceStable(m.keys, func(i, j int) bool {
		return comparer(m.keys[i], m.keys[j]) < 0
	})
	for i, k := range m.keys {
		if _, ok := m.index[k]; ok {
			m.index[k] = i
		}
	}
}
