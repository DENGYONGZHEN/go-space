package stackandqueue

import "testing"

// Example 1:
// Input: s = "abbaca"
// Output: "ca"
// Explanation:
// For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.
// The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

// Example 2:
// Input: s = "azxxzy"
// Output: "ay"

func TestRemoveAllAdjacentDuplicatesInString(t *testing.T) {

	testCase := []struct {
		input string
		want  string
	}{
		{
			input: "abbaca",
			want:  "ca",
		},
		{
			input: "azxxzy",
			want:  "ay",
		},
	}

	for _, tc := range testCase {
		got := removeDuplicates(tc.input)
		if got != tc.want {
			t.Fatalf("TestRemoveAllAdjacentDuplicatesInString want %v,got %v", tc.want, got)
		}
	}
}
