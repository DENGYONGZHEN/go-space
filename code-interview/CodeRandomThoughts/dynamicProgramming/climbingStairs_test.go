package dynamicprogramming

import "testing"

// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps

// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{
			input:  2,
			output: 2,
		},
		{
			input:  3,
			output: 3,
		},
	}
	for _, tc := range testCases {
		got := climbStairs(tc.input)

		if got != tc.output {
			t.Fatalf("TestClimbStairs want %v, got %v", tc.output, got)
		}
	}
}
