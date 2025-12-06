package array

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

// Example 2:
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]

func TestFourSum(t *testing.T) {

	testcases := []struct {
		nums   []int
		target int
		want   [][]int
	}{

		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},

		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}}},
	}

	for _, tc := range testcases {
		got := fourSum(tc.nums, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("fourSum(%v,%v) should got: %v, instead: %v", tc.nums, tc.target, tc.want, got)
		}
	}

}
