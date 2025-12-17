package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = "2"
// Output: ["a","b","c"]

func TestLetterCombinations(t *testing.T) {

	testCases := []struct {
		digits string
		want   []string
	}{
		{
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		got := letterCombinations(tc.digits)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("TestLetterCombinations want %v, got %v", tc.want, got)
		}
	}
}
