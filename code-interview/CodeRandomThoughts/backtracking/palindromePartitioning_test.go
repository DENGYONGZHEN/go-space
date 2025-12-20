package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]

// Example 2:
// Input: s = "a"
// Output: [["a"]]

func TestPartition(t *testing.T) {

	testCases := []struct {
		s    string
		want [][]string
	}{
		{
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:    "a",
			want: [][]string{{"a"}},
		},
	}

	for _, tc := range testCases {
		got := partition(tc.s)

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestPartition want %v, got %v", tc.want, got)
		}
	}
}
