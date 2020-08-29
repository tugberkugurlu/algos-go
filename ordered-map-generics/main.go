package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := NewLinkedListBasedOrderedMap()
	seedCount := 1000000
	for i := 1; i <= seedCount; i++ {
		m.Set(i, fmt.Sprintf("string%d", i))
	}

	func() {
		for i := 0; i < 100000; i++ {
			m.Delete(rand.Intn(seedCount-1) + 1)
		}
	}()

	iterator := m.Iterator()
	var count int
	for {
		i, _, _ := iterator()
		if i == nil {
			break
		}
		count++
	}
	fmt.Printf("count: %d\n", count)
}
