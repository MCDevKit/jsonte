package utils

import "sort"

// NavigableMap is a map that can be navigated in order of keys. It is not thread-safe.
type NavigableMap[K comparable, V any] struct {
	keys []K
	data map[K]V
}

// NewNavigableMap creates a new NavigableMap.
func NewNavigableMap[K comparable, V any]() NavigableMap[K, V] {
	return NavigableMap[K, V]{
		data: map[K]V{},
		keys: []K{},
	}
}

// ToNavigableMap creates a new NavigableMap with given keys and values.
func ToNavigableMap(entries ...interface{}) NavigableMap[string, interface{}] {
	n := NavigableMap[string, interface{}]{
		data: map[string]interface{}{},
		keys: []string{},
	}
	for i := 0; i < len(entries); i += 2 {
		n.Put(entries[i].(string), entries[i+1])
	}
	return n
}

// MapToNavigableMap creates a new NavigableMap from a golang map.
func MapToNavigableMap(entries map[string]interface{}) NavigableMap[string, interface{}] {
	n := NavigableMap[string, interface{}]{
		data: map[string]interface{}{},
		keys: []string{},
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

// Put puts the value associated with the key.
func (m *NavigableMap[K, V]) Put(key K, value V) {
	m.data[key] = value
	if !m.ContainsKey(key) {
		m.keys = append(m.keys, key)
	}
}

// Remove removes the value associated with the key.
func (m *NavigableMap[K, V]) Remove(key K) {
	delete(m.data, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

// ContainsKey returns true if the key exists.
func (m *NavigableMap[K, V]) ContainsKey(key K) bool {
	for _, k := range m.keys {
		if k == key {
			return true
		}
	}
	return false
}

// ContainsMatchingKey returns true if the matching key exists.
func (m *NavigableMap[K, V]) ContainsMatchingKey(matchFunc func(K) bool) bool {
	for _, k := range m.keys {
		if matchFunc(k) {
			return true
		}
	}
	return false
}

// Keys returns the keys in order.
func (m *NavigableMap[K, V]) Keys() []K {
	return m.keys
}

// Values returns the values in order of keys.
func (m *NavigableMap[K, V]) Values() []V {
	var values []V
	for _, k := range m.keys {
		values = append(values, m.data[k])
	}
	return values
}

// Size returns the size of the map.
func (m *NavigableMap[K, V]) Size() int {
	return len(m.keys)
}

// IsEmpty returns true if the map is empty.
func (m *NavigableMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Clear clears the map.
func (m *NavigableMap[K, V]) Clear() {
	m.keys = []K{}
	m.data = map[K]V{}
}

// Sort sorts the map with the specified comparator.
func (m *NavigableMap[K, V]) Sort(comparer func(K, K) int) {
	sort.SliceStable(m.keys, func(i, j int) bool {
		return comparer(m.keys[i], m.keys[j]) < 0
	})
}
