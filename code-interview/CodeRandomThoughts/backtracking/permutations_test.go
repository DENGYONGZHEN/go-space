package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// Example 2:
// Input: nums = [0,1]
// Output: [[0,1],[1,0]]

// Example 3:
// Input: nums = [1]
// Output: [[1]]
func TestPermute(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{1}, want: [][]int{{1}},
		},
	}
	for _, tc := range testCases {
		got := permute(tc.nums)

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestPermute want %v, got %v", tc.want, got)
		}
	}
}
