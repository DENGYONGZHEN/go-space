package string

import (
	"testing"
)

// Example 1:
// Input: s = "the sky is blue"
// Output: "blue is sky the"

// Example 2:
// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

// Example 3:
// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

func TestReverseWords(t *testing.T) {
	test := []struct {
		s    string
		want string
	}{
		{s: "the sky is blue", want: "blue is sky the"},
		{s: "  hello world  ", want: "world hello"},
		{s: "a good   example", want: "example good a"},
	}

	for _, tt := range test {

		got := reverseWords(tt.s)
		if got != tt.want {
			t.Errorf("TestReverseString = %v; want %v", got, tt.want)
		}

	}
}
