package container_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riabininkf/gox/container"
)

func TestHeap(t *testing.T) {
	heap := container.NewHeap(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 0, heap.Len())

	top, ok := heap.Top()
	assert.False(t, ok)
	assert.Equal(t, 0, top)

	val, ok := heap.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	count := 5
	for i := 1; i < count; i++ {
		heap.Push(i)

		top, ok = heap.Top()
		assert.True(t, ok)
		assert.Equal(t, 1, top)
		assert.Equal(t, i, heap.Len())
	}

	for i := 1; i < count; i++ {
		val, ok = heap.Pop()

		assert.True(t, ok)
		assert.Equal(t, i, val)
		assert.Equal(t, count-i-1, heap.Len())
	}

	top, ok = heap.Top()
	assert.False(t, ok)
	assert.Equal(t, 0, top)
	assert.Equal(t, 0, heap.Len())

	val, ok = heap.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}
