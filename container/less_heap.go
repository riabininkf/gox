package container

// lessHeap is an implementation of heap.Interface with a custom less function.
type lessHeap[T any] struct {
	elements []T
	less     func(a, b T) bool
}

// Len implements [(*sort.Interface).Len]
func (h *lessHeap[T]) Len() int {
	return len(h.elements)
}

// Less implements [(*sort.Interface).Less]
func (h *lessHeap[T]) Less(i, j int) bool {
	return h.less(h.elements[i], h.elements[j])
}

// Swap implements [(*sort.Interface).Swap]
func (h *lessHeap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// Push implements [(*heap.Interface).Push]
func (h *lessHeap[T]) Push(x any) {
	h.elements = append(h.elements, x.(T))
}

// Pop implements [(*heap.Interface).Pop]
func (h *lessHeap[T]) Pop() any {
	old := h.elements
	n := len(old)
	v := old[n-1]
	h.elements = old[:n-1]
	return v
}
