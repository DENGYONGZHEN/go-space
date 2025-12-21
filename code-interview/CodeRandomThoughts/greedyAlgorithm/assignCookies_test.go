package greedyalgorithm

import (
	"testing"
)

// Example 1:
// Input: g = [1,2,3], s = [1,1]
// Output: 1
// Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
// And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
// You need to output 1.

// Example 2:
// Input: g = [1,2], s = [1,2,3]
// Output: 2
// Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2.
// You have 3 cookies and their sizes are big enough to gratify all of the children,
// You need to output 2.

func TestFindContentChildren(t *testing.T) {
	testCases := []struct {
		g    []int
		s    []int
		want int
	}{
		{
			g: []int{1, 2, 3}, s: []int{1, 1}, want: 1,
		},
		{
			g: []int{1, 2}, s: []int{1, 2, 3}, want: 2,
		},
	}
	for _, tc := range testCases {

		got := findContentChildren(tc.g, tc.s)

		if tc.want != got {
			t.Fatalf("TestFindContentChildren want %v,got %v", tc.want, got)
		}
	}
}
