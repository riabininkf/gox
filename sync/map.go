package sync

import "sync"

// Map provides a generic, concurrent-safe map abstraction using sync.Map internally.
type Map[K comparable, V any] struct {
	m sync.Map
}

// NewMap creates a new sync.Map with generic key and value types.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

// Store implements [(*sync.Map).Store]
func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Load implements [(*sync.Map).Load]
func (m *Map[K, V]) Load(key K) (V, bool) {
	var (
		ok  bool
		val any
	)
	if val, ok = m.m.Load(key); val != nil {
		return val.(V), ok
	}

	var zero V
	return zero, ok
}

// Delete implements [(*sync.Map).Delete]
func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// LoadOrStore implements [(*sync.Map).LoadOrStore]
func (m *Map[K, V]) LoadOrStore(key K, val V) (V, bool) {
	var (
		loaded bool
		actual any
	)
	if actual, loaded = m.m.LoadOrStore(key, val); actual != nil {
		return actual.(V), loaded
	}

	var zero V
	return zero, loaded
}

// LoadAndDelete implements [(*sync.Map).LoadAndDelete]
func (m *Map[K, V]) LoadAndDelete(key K) (V, bool) {
	var (
		val    any
		loaded bool
	)
	if val, loaded = m.m.LoadAndDelete(key); val != nil {
		return val.(V), loaded
	}

	var zero V
	return zero, loaded
}

// Range implements [(*sync.Map).Range]
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// CompareAndSwap implements [(*sync.Map).CompareAndSwap]
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

// Swap implements [(*sync.Map).Swap]
func (m *Map[K, V]) Swap(key K, value V) (V, bool) {
	prev, loaded := m.m.Swap(key, value)

	if prev != nil {
		return prev.(V), loaded
	}

	var zero V
	return zero, loaded
}

// CompareAndDelete implements [(*sync.Map).CompareAndDelete]
func (m *Map[K, V]) CompareAndDelete(key K, old V) bool {
	return m.m.CompareAndDelete(key, old)
}
