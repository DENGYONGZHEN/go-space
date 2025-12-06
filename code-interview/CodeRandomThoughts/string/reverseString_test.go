package string

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

func TestReverseString(t *testing.T) {
	test := []struct {
		s    []byte
		want []byte
	}{
		{s: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, tt := range test {

		reverseString(tt.s)
		if !reflect.DeepEqual(tt.s, tt.want) {
			t.Errorf("TestReverseString = %v; want %v", tt.s, tt.want)
		}

	}
}
