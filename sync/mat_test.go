package sync_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/riabininkf/gox/sync"
)

func TestMap(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		val, ok := m.Load(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)

		assert.False(t, m.CompareAndSwap(1, "a", "b"))

		val, ok = m.LoadAndDelete(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)

		val, ok = m.LoadAndDelete(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)
	})

	t.Run("store", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		m.Store(1, "a")
		val, ok := m.Load(1)
		assert.True(t, ok)
		assert.Equal(t, "a", val)
	})

	t.Run("compare and swap", func(t *testing.T) {
		m := sync.NewMap[int, string]()
		m.Store(1, "a")

		// success comparison
		assert.True(t, m.CompareAndSwap(1, "a", "b"))
		val, ok := m.Load(1)
		assert.True(t, ok)
		assert.Equal(t, "b", val)

		// failed comparison
		assert.False(t, m.CompareAndSwap(1, "c", "d"))
		val, ok = m.Load(1)
		assert.True(t, ok)
		assert.Equal(t, "b", val)
	})

	t.Run("delete", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		m.Store(1, "a")

		val, ok := m.Load(1)
		assert.True(t, ok)
		assert.Equal(t, "a", val)

		m.Delete(1)
		val, ok = m.Load(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)
	})

	t.Run("load or store", func(t *testing.T) {
		m := sync.NewMap[int, error]()

		old, loaded := m.LoadOrStore(1, nil)
		assert.False(t, loaded)
		assert.Nil(t, old)

		m.Delete(1)

		old, loaded = m.Load(1)
		assert.False(t, loaded)
		assert.Nil(t, old)

		m.Store(1, assert.AnError)
		old, loaded = m.LoadOrStore(1, nil)
		assert.True(t, loaded)
		assert.Equal(t, assert.AnError, old)
	})

	t.Run("load and delete", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		m.Store(1, "a")
		val, ok := m.LoadAndDelete(1)
		assert.True(t, ok)
		assert.Equal(t, "a", val)

		val, ok = m.LoadAndDelete(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)
	})

	t.Run("range", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		testCases := map[int]string{
			1: "1",
			2: "2",
			3: "3",
		}

		for key, val := range testCases {
			m.Store(key, val)
		}

		m.Range(func(k int, v string) bool {
			assert.Equal(t, testCases[k], v)
			delete(testCases, k)
			return true
		})

		assert.Empty(t, testCases)
	})

	t.Run("swap", func(t *testing.T) {
		m := sync.NewMap[int, error]()

		m.Store(1, nil)

		old, ok := m.Swap(1, assert.AnError)
		assert.True(t, ok)
		assert.Nil(t, old)

		m.Delete(1)

		var loaded bool
		old, loaded = m.Load(1)
		assert.False(t, loaded)
		assert.Nil(t, old)

		m.Store(1, assert.AnError)

		old, ok = m.Swap(1, nil)
		assert.True(t, ok)
		assert.Equal(t, assert.AnError, old)
	})

	t.Run("compare and delete", func(t *testing.T) {
		m := sync.NewMap[int, string]()

		m.Store(1, "a")
		m.CompareAndDelete(1, "b")

		val, ok := m.Load(1)
		assert.True(t, ok)
		assert.Equal(t, "a", val)

		m.CompareAndDelete(1, "a")
		val, ok = m.Load(1)
		assert.False(t, ok)
		assert.Equal(t, "", val)
	})
}
