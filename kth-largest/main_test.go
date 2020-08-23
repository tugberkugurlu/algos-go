package main

import (
	"fmt"
	"math/rand"
	"testing"
)

var testdata = []struct {
	in  []int
	k   int
	f   func(nums []int, k int) int
	out int
}{
	{[]int{3, 2, 1, 5, 6, 4}, 2, FindLargestKthElement, 5},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, FindLargestKthElement, 4},
	{[]int{3, 2, 1, 5, 6, 4}, 2, FindLargestKthElementWithSort, 5},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, FindLargestKthElementWithSort, 4},
}

func TestLogic(t *testing.T) {
	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			out := tt.f(tt.in, tt.k)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}

var nums []int
func TestMain(m *testing.M) {
	maxVal := 10000000
	nums = make([]int, maxVal)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(maxVal)
	}
	m.Run()
}

func BenchmarkFindLargestKthElementK500(b *testing.B) {
	k := 500
	for n := 0; n < b.N; n++ {
		nums2 := make([]int, len(nums))
		for i, v := range nums {
			nums2[i] = v
		}
		FindLargestKthElement(nums2, k)
	}
}

func BenchmarkFindLargestKthElementWithSortK500(b *testing.B) {
	k := 500
	for n := 0; n < b.N; n++ {
		nums2 := make([]int, len(nums))
		for i, v := range nums {
			nums2[i] = v
		}
		FindLargestKthElementWithSort(nums2, k)
	}
}

