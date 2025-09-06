package container

import "container/ring"

// Ring provides a generic wrapper for ring.Ring.
type Ring[T any] struct {
	r *ring.Ring
}

// NewRing creates a new ring of n generic elements.
func NewRing[T any](n int) *Ring[T] {
	return &Ring[T]{r: ring.New(n)}
}

// Len implements [(*ring.Ring).Len]
func (r *Ring[T]) Len() int {
	return r.r.Len()
}

// Value implements [(*ring.Ring).Value]
func (r *Ring[T]) Value() (T, bool) {
	if r == nil || r.r == nil || r.r.Value == nil {
		var zero T
		return zero, false
	}

	return r.r.Value.(T), true
}

// Set sets the value of the ring element at the current position.
func (r *Ring[T]) Set(v T) {
	r.r.Value = v
}

// Next implements [(*ring.Ring).Next]
func (r *Ring[T]) Next() *Ring[T] {
	r.r = r.r.Next()
	return r
}

// Prev implements [(*ring.Ring).Prev]
func (r *Ring[T]) Prev() *Ring[T] {
	r.r = r.r.Prev()
	return r
}

// Do implements [(*ring.Ring).Do]
func (r *Ring[T]) Do(f func(T)) {
	r.r.Do(func(v any) {
		val, ok := v.(T)
		if ok {
			f(val)
		}
	})
}

// Link implements [(*ring.Ring).Link]
func (r *Ring[T]) Link(other *Ring[T]) *Ring[T] {
	return &Ring[T]{r: r.r.Link(other.r)}
}

// Unlink implements [(*ring.Ring).Unlink]
func (r *Ring[T]) Unlink(n int) *Ring[T] {
	return &Ring[T]{r: r.r.Unlink(n)}
}
