package dynamicprogramming

import (
	"testing"
)

// Example 1:
// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

// Example 2:
// Input: n = 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

// Example 3:
// Input: n = 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
func TestFib(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{
			input: 2, output: 1,
		},
		{
			input: 3, output: 2,
		},
		{
			input: 4, output: 3,
		},
	}
	for _, tc := range testCases {
		got := fib(tc.input)

		if tc.output != got {
			t.Fatalf("TestFib want %v,got %v", tc.output, got)
		}
	}
}
