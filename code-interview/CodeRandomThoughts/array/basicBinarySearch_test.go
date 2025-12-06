package array

import "testing"

func TestBasicBinarySearch(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{nums: []int{-5, 1, 2, 6, 7}, target: 6, want: 3},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
	}

	for _, val := range tt {
		got := search(val.nums, val.target)
		if got != val.want {
			t.Errorf("BasicBinarySearch(%v,%v)=%v; want %v instead \n", val.nums, val.target, got, val.want)
		}
	}
}
