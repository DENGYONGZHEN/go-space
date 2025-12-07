package string

import (
	"testing"
)

// Example 1:
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"

// Example 2:
// Input: s = "abcd", k = 2
// Output: "bacd"

func TestReverseString2(t *testing.T) {
	test := []struct {
		s    string
		k    int
		want string
	}{
		{s: "abcdefg", k: 2, want: "bacdfeg"},
		{s: "abcd", k: 2, want: "bacd"},
	}

	for _, tt := range test {

		got := reverseStr(tt.s, tt.k)
		if got != tt.want {
			t.Errorf("TestReverseString = %v; want %v", got, tt.want)
		}

	}
}
