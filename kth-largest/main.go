package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {
	nums := []int{3,2,20,5,3,1,2,5,6,9,10,4}

	// initialize the heap data structure
	h := &IntHeap{}

	// add all the values to heap, O(n log n)
	for _, val := range nums { // O(n)
		heap.Push(h, val) // O(log n)
	}

	// print all the values from the heap
	// which should be in ascending order
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d,", heap.Pop(h).(int))
	}

	fmt.Println()
	fmt.Println(FindLargestKthElement(nums, 4))
	fmt.Println(FindBestKElements(nums, 4))
	fmt.Println(FindLargestKthElementWithSort(nums, 4))
	fmt.Println(FindBestKElementsWithSort(nums, 4))
}

// O(n log k + n-k log k)
// https://leetcode.com/problems/kth-largest-element-in-an-array/
func FindLargestKthElement(nums []int, k int) int {
	h := &IntHeap{}
	for _, val := range nums { // O(N)
		heap.Push(h, val) // O(log K)
		if h.Len() > k {
			heap.Pop(h) // O(log K)
		}
	}

	return heap.Pop(h).(int) // O(log K)
}

// O(n log k + n-k log k + k log k)
// simplified: // O(n log k + k log k)
func FindBestKElements(nums []int, k int) []int {
	h := &IntHeap{}
	for _, val := range nums { // O(N)
		heap.Push(h, val) // O(log K)
		if h.Len() > k {
			heap.Pop(h) // O(log K)
		}
	}

	return func() []int {
		result := make([]int, h.Len())
		initialLen := h.Len()
		for i := initialLen; i > 0; i-- {
			result[i-1] = heap.Pop(h).(int)
		}
		return result
	}()
}

// O (n log n)
func FindLargestKthElementWithSort(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { // O (n log n)
		return nums[i] < nums[j]
	})
	return nums[len(nums)-k] // O(1)
}

// O (n log n + k)
func FindBestKElementsWithSort(nums []int, k int) []int {
	sort.Slice(nums, func(i, j int) bool { // O (n log n)
		return nums[i] > nums[j]
	})

	return func() []int {
		result := make([]int, k)
		for i := 0; i < k; i++ {
			result[i] = nums[i]
		}
		return result
	}()
}