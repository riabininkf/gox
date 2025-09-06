package container_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riabininkf/gox/container"
)

func TestRing(t *testing.T) {
	ln := 5

	ring := container.NewRing[int](ln)

	assert.NotPanics(t, func() {
		val, ok := ring.Value()
		assert.Equal(t, 0, val)
		assert.False(t, ok)
	}, "Value() should panic on non-initialized values")

	assert.Equal(t, ln, ring.Len())

	// initialize values
	for i := 1; i <= ln; i++ {
		ring.Set(i)

		val, ok := ring.Value()
		assert.Equal(t, i, val)
		assert.True(t, ok)

		ring.Next()
	}

	// link extra ring
	ring.Link(container.NewRing[int](1))
	assert.Equal(t, ln+1, ring.Len())

	// unlink extra ring
	ring.Unlink(1)
	assert.Equal(t, ln, ring.Len())

	// traverse forward
	for i := 1; i < ln; i++ {
		val, ok := ring.Value()
		assert.Equal(t, i, val)
		assert.True(t, ok)
		ring.Next()
	}

	// traverse backward
	for i := ln; i > 1; i-- {
		val, ok := ring.Value()
		assert.Equal(t, i, val)
		assert.True(t, ok)
		ring.Prev()
	}

	expVal := 1
	ring.Do(func(i int) {
		assert.Equal(t, expVal, i)
		expVal++
	})
}
