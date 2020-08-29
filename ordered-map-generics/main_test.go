package main

import (
	"fmt"
	"github.com/deliveroo/assert-go"
	"math/rand"
	"testing"
)

func TestNewLinkedListBasedOrderedMap(t *testing.T) {
	t.Run("delete works as expected when the value is not at the edges", func(t *testing.T) {
		m := NewLinkedListBasedOrderedMap()
		m.Set(3, "val3")
		m.Set(1, "val1")
		m.Set(2, "val2")
		m.Set(4, "val4")

		m.Delete(2)

		it := m.Iterator()
		var keys []int
		var vals []string
		for {
			i, key, val := it()
			if i == nil {
				break
			}
			keys = append(keys, *key)
			vals = append(vals, val.(string))
		}

		assert.Equal(t, keys, []int{3,1,4})
		assert.Equal(t, vals, []string{"val3","val1","val4"})
	})

	t.Run("set after delete works as expected", func(t *testing.T) {
		m := NewLinkedListBasedOrderedMap()
		m.Set(3, "val3")
		m.Set(1, "val1")
		m.Set(2, "val2")
		m.Set(4, "val4")
		m.Delete(2)
		m.Set(7, "val4")
		m.Set(5, "val5")
		m.Delete(7)

		it := m.Iterator()
		var keys []int
		var vals []string
		for {
			i, key, val := it()
			if i == nil {
				break
			}
			keys = append(keys, *key)
			vals = append(vals, val.(string))
		}

		assert.Equal(t, keys, []int{3,1,4,5})
		assert.Equal(t, vals, []string{"val3","val1","val4","val5"})
	})

	t.Run("delete works as expected when the value is at the front", func(t *testing.T) {
		m := NewLinkedListBasedOrderedMap()
		m.Set(3, "val3")
		m.Set(1, "val1")
		m.Set(2, "val2")
		m.Set(4, "val4")

		m.Delete(3)

		it := m.Iterator()
		var keys []int
		var vals []string
		for {
			i, key, val := it()
			if i == nil {
				break
			}
			keys = append(keys, *key)
			vals = append(vals, val.(string))
		}

		assert.Equal(t, keys, []int{1,2,4})
		assert.Equal(t, vals, []string{"val1","val2","val4"})
	})

	t.Run("delete works as expected when the value is at the back", func(t *testing.T) {
		m := NewLinkedListBasedOrderedMap()
		m.Set(3, "val3")
		m.Set(1, "val1")
		m.Set(2, "val2")
		m.Set(4, "val4")

		m.Delete(4)

		it := m.Iterator()
		var keys []int
		var vals []string
		for {
			i, key, val := it()
			if i == nil {
				break
			}
			keys = append(keys, *key)
			vals = append(vals, val.(string))
		}

		assert.Equal(t, keys, []int{3,1,2,})
		assert.Equal(t, vals, []string{"val3", "val1","val2"})
	})

	t.Run("delete works as expected when there is only one value", func(t *testing.T) {
		m := NewLinkedListBasedOrderedMap()
		m.Set(3, "val3")

		m.Delete(3)

		it := m.Iterator()
		var keys []int
		var vals []string
		for {
			i, key, val := it()
			if i == nil {
				break
			}
			keys = append(keys, *key)
			vals = append(vals, val.(string))
		}

		assert.Equal(t, len(keys), 0)
		assert.Equal(t, len(vals), 0)
	})
}

func BenchmarkOrderedMapLinkedListBasedDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		seedCount := 1000000
		m := NewLinkedListBasedOrderedMap()
		for i := 1; i <= seedCount; i++ {
			m.Set(i, fmt.Sprintf("string%d", i))
		}
		for i := 0; i < 100000; i++ {
			m.Delete(rand.Intn(seedCount-1) + 1)
		}
	}
}

func BenchmarkOrderedMapSliceBasedDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		seedCount := 1000000
		m := NewOrderedMap()
		for i := 1; i <= seedCount; i++ {
			m.Set(i, fmt.Sprintf("string%d", i))
		}
		for i := 0; i < 100000; i++ {
			m.Delete(rand.Intn(seedCount-1) + 1)
		}
	}
}
