package hashtable

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	tt := []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
		want  int
	}{
		{nums1: []int{-1, -1}, nums2: []int{-1, 1}, nums3: []int{-1, 1}, nums4: []int{1, -1}, want: 6},
		{nums1: []int{1, 2}, nums2: []int{-2, -1}, nums3: []int{-1, 2}, nums4: []int{0, 2}, want: 2},
		{nums1: []int{0}, nums2: []int{0}, nums3: []int{0}, nums4: []int{0}, want: 1},
	}

	for _, tc := range tt {
		got := fourSumCount(tc.nums1, tc.nums2, tc.nums3, tc.nums4)
		if tc.want != got {
			t.Errorf("fourSumCount(%v,%v,%v,%v)=%v; want %v instead \n", tc.nums1, tc.nums2, tc.nums3, tc.nums4, tc.want, got)
		}
	}
}
