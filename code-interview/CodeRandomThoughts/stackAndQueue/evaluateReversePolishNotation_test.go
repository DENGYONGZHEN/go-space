package stackandqueue

import "testing"

// Example 1:
// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

// Example 2:
// Input: tokens = ["4","13","5","/","+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

// Example 3:
// Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// Output: 22
// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22
func TestEvalRPN(t *testing.T) {

	testcases := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"2", "1", "+", "3", "*"},
			want:  9,
		},
		{
			input: []string{"4", "13", "5", "/", "+"},
			want:  6,
		},
		{
			input: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:  22,
		},
	}

	for _, tc := range testcases {
		got := evalRPN(tc.input)
		if got != tc.want {
			t.Fatalf("TestEvalRPN want %v, got %v", tc.want, got)
		}
	}
}
