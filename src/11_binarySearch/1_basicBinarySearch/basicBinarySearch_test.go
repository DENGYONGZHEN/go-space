package main

import "testing"

func TestBasicBinarySearch(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 9},
		{nums: []int{2, 1, 6, 7, -5}, target: 6, want: 6},
	}

	for _, val := range tt {
		got := BasicBinarySearch(val.nums, val.target)
		if got != val.want {
			t.Errorf("BasicBinarySearch(%v,%v)=%v; want %v instead \n", val.nums, val.target, got, val.want)
		}
	}
}
