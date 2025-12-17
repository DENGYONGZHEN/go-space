package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: k = 3, n = 7
// Output: [[1,2,4]]
// Explanation:
// 1 + 2 + 4 = 7
// There are no other valid combinations.

// Example 2:
// Input: k = 3, n = 9
// Output: [[1,2,6],[1,3,5],[2,3,4]]
// Explanation:
// 1 + 2 + 6 = 9
// 1 + 3 + 5 = 9
// 2 + 3 + 4 = 9
// There are no other valid combinations.

// Example 3:
// Input: k = 4, n = 1
// Output: []
// Explanation: There are no valid combinations.
// Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.

func TestCombinationSum3(t *testing.T) {
	testCases := []struct {
		k    int
		n    int
		want [][]int
	}{
		{
			k: 3, n: 7, want: [][]int{{1, 2, 4}},
		},
		{
			k: 3, n: 9, want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			k: 4, n: 1, want: [][]int{},
		},
	}

	for _, tc := range testCases {
		got := combinationSum3(tc.k, tc.n)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestCombinationSum3 want %v, got %v", tc.want, got)
		}
	}
}
