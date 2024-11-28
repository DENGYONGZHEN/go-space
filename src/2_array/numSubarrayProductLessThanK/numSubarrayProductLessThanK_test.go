package main

import (
	"os"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {

	tt := []struct {
		nums   []int
		want   int
		target int
	}{
		{nums: []int{10, 5, 2, 6}, want: 8, target: 100},
		{nums: []int{1, 2, 3}, want: 0, target: 0},
	}
	for _, val := range tt {

		got := numSubarrayProductLessThanK(val.nums, val.target)
		if got != val.want {
			t.Errorf("Expected numSubarrayProductLessThanK(%v,%d)=%d; got %d instead\n", val.nums, val.target, val.want, got)
			os.Exit(1)
		}
	}
}
