package binarysearch

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.00000,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
	}

	for _, tc := range testCases {
		val := findMedianSortedArrays(tc.nums1, tc.nums2)
		if val != tc.want {
			t.Errorf("findMin(%v,%v),want %v, but get %v", tc.nums1, tc.nums2, tc.want, val)
		}
	}
}
