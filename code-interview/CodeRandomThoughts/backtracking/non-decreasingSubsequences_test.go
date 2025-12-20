package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [4,6,7,7]
// Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

// Example 2:
// Input: nums = [4,4,3,2,1]
// Output: [[4,4]]

func TestFindSubsequences(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{4, 6, 7, 7}, want: [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
		},
		{
			nums: []int{4, 4, 3, 2, 1}, want: [][]int{{4, 4}},
		},
	}
	for _, tc := range testCases {
		got := findSubsequences(tc.nums)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("TestFindSubsequences want %v, got %v", tc.want, got)
		}
	}
}
