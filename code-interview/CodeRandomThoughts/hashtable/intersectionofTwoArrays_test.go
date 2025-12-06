package hashtable

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

func TestIntersection(t *testing.T) {
	tt := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{9, 4}},
	}

	for _, tc := range tt {
		got := intersection(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("intersection(%v,%v)=%v; want %v instead \n", tc.nums1, tc.nums2, tc.want, got)
		}
	}
}
