package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
// Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

// Example 2:
// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is 1 choose 1 = 1 total combination.

func TestCombination(t *testing.T) {

	testCase := []struct {
		input  int
		k      int
		output [][]int
	}{
		{
			input: 4, k: 2, output: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			input: 1, k: 1, output: [][]int{{1}},
		},
	}

	for _, tc := range testCase {
		got := combine(tc.input, tc.k)

		if !reflect.DeepEqual(got, tc.output) {
			t.Fatalf("TestCombination want %v, got %v", tc.output, got)
		}
	}

}
