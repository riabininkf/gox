package container

import "container/heap"

// NewHeap creates a new generic heap.
func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		heap: &lessHeap[T]{less: less},
	}
}

// Heap is a generic implementation of the heap.
// It's not safe for concurrent use, and such safety should be provided by the caller.
type Heap[T any] struct {
	heap *lessHeap[T]
}

// Top returns the top element of the heap without removing it, or false if the heap is empty.
func (h *Heap[T]) Top() (T, bool) {
	if h.heap.Len() == 0 {
		var zero T
		return zero, false
	}

	return h.heap.elements[0], true
}

// Push adds an element to the heap.
func (h *Heap[T]) Push(x T) {
	heap.Push(h.heap, x)
}

// Pop removes and returns the top element. Returns (zero, false) if the heap is empty.
func (h *Heap[T]) Pop() (T, bool) {
	if h.heap.Len() == 0 {
		var zero T
		return zero, false
	}

	return heap.Pop(h.heap).(T), true
}

// Len returns the number of elements in the heap.
func (h *Heap[T]) Len() int {
	return h.heap.Len()
}
