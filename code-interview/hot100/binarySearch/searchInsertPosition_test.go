package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			nums:   []int{3, 5, 7, 9, 10},
			target: 8,
			want:   3,
		},
	}

	for _, tc := range testCases {
		val := searchInsert(tc.nums, tc.target)
		if val != tc.want {
			t.Errorf("searchInsert(%v,%v),want %v, but %v", tc.nums, tc.target, tc.want, val)
		}
	}
}
