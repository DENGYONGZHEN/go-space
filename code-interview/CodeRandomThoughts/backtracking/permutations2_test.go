package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [1,1,2]
// Output: [[1,1,2],[1,2,1],[2,1,1]]

// Example 2:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func TestPermuteUnique(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 1, 2}, want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tc := range testCases {
		got := permuteUnique(tc.nums)

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestPermute want %v, got %v", tc.want, got)
		}
	}
}
