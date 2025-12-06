package binarysearch

import "testing"

func TestFindMin(t *testing.T) {

	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, tc := range testCases {
		val := findMin(tc.nums)
		if val != tc.want {
			t.Errorf("findMin(%v),want %d, but get %d", tc.nums, tc.want, val)
		}
	}
}
