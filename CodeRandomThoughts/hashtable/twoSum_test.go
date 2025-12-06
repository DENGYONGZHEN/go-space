package hashtable

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

func TestTwoSum(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tc := range tt {
		got := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("twoSum(%v,%v)=%v; want %v instead \n", tc.nums, tc.target, got, tc.want)
		}
	}
}
