package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Example 2:
// Input: nums = [0]
// Output: [[],[0]]

func TestSubsets(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}
	for _, tc := range testCases {
		got := subsets(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestSubsets want %v, got %v", tc.want, got)
		}
	}
}
