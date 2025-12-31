package greedyalgorithm

import "testing"

// Example 1:

// Input: nums = [1,7,4,9,2,5]
// Output: 6
// Explanation: The entire sequence is a wiggle sequence with differences (6, -3, 5, -7, 3).

// Example 2:
// Input: nums = [1,17,5,10,13,15,10,5,16,8]
// Output: 7
// Explanation: There are several subsequences that achieve this length.
// One is [1, 17, 10, 13, 10, 16, 8] with differences (16, -7, 3, -3, 6, -8).

// Example 3:
// Input: nums = [1,2,3,4,5,6,7,8,9]
// Output: 2

func TestWiggleMaxLength(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 7, 4, 9, 2, 5}, want: 6,
		},
		{
			nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, want: 7,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 2,
		},
	}
	for _, tc := range testCases {

		got := wiggleMaxLength(tc.nums)

		if tc.want != got {
			t.Fatalf("TestWiggleMaxLength want %v, got %v", tc.want, got)
		}

	}
}
