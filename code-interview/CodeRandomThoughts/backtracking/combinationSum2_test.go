package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [[1,1,6],[1,2,5],[1,7],[2,6]]

// Example 2:
// Input: candidates = [2,5,2,1,2], target = 5
// Output:
// [[1,2,2],[5]]

func TestCombinationSum2(t *testing.T) {
	testcases := []struct {
		candidate []int
		target    int
		want      [][]int
	}{
		{
			candidate: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			candidate: []int{2, 5, 2, 1, 2}, target: 5, want: [][]int{{1, 2, 2}, {5}},
		},
	}

	for _, tc := range testcases {
		got := combinationSum2(tc.candidate, tc.target)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("TestCombinationSum want %v, got %v", tc.want, got)
		}
	}
}
