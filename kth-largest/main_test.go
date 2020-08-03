package main

import (
	"fmt"
	"testing"
)

var testdata = []struct {
	in  []int
	k   int
	out int
}{
	{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
}

func TestLogic(t *testing.T) {
	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			out := FindLargestKthElement(tt.in, tt.k)
			if out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}
