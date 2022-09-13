package utils

import "sort"

type NavigableMap[K comparable, V any] struct {
	keys []K
	data map[K]V
}

func NewNavigableMap[K comparable, V any]() NavigableMap[K, V] {
	return NavigableMap[K, V]{
		data: map[K]V{},
		keys: []K{},
	}
}

func (m *NavigableMap[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *NavigableMap[K, V]) Put(key K, value V) {
	m.data[key] = value
	if !m.ContainsKey(key) {
		m.keys = append(m.keys, key)
	}
}

func (m *NavigableMap[K, V]) Remove(key K) {
	delete(m.data, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *NavigableMap[K, V]) ContainsKey(key K) bool {
	for _, k := range m.keys {
		if k == key {
			return true
		}
	}
	return false
}

func (m *NavigableMap[K, V]) Keys() []K {
	return m.keys
}

func (m *NavigableMap[K, V]) Values() []V {
	var values []V
	for _, k := range m.keys {
		values = append(values, m.data[k])
	}
	return values
}

func (m *NavigableMap[K, V]) Size() int {
	return len(m.keys)
}

func (m *NavigableMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *NavigableMap[K, V]) Clear() {
	m.keys = []K{}
	m.data = map[K]V{}
}

func (m *NavigableMap[K, V]) Sort(comparer func(K, K) int) {
	sort.SliceStable(m.keys, func(i, j int) bool {
		return comparer(m.keys[i], m.keys[j]) < 0
	})
}
