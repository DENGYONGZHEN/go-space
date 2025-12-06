package binarysearch

import "testing"

func TestSearchRotatedSortedArray(t *testing.T) {

	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
	}

	for _, tc := range testCases {
		val := search(tc.nums, tc.target)
		if val != tc.want {
			t.Errorf("search(%v,%v),want %v,but get %v", tc.nums, tc.target, tc.want, val)
		}
	}
}
