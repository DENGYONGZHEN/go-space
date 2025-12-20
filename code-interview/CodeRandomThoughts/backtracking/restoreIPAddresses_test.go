package backtracking

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]

// Example 2:
// Input: s = "0000"
// Output: ["0.0.0.0"]

// Example 3:
// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

func TestRestoreIpAddresses(t *testing.T) {
	testCases := []struct {
		s    string
		want []string
	}{
		{
			s: "25525511135", want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s: "0000", want: []string{"0.0.0.0"},
		},
		{
			s: "101023", want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tC := range testCases {
		got := restoreIpAddresses(tC.s)

		if !reflect.DeepEqual(got, tC.want) {
			t.Fatalf("TestRestoreIpAddresses want %v, got %v", tC.want, got)
		}
	}
}
