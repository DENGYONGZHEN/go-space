package greedyalgorithm

import "testing"

// Example 1:
// Input: nums = [4,2,3], k = 1
// Output: 5
// Explanation: Choose index 1 and nums becomes [4,-2,3].

// Example 2:
// Input: nums = [3,-1,0,2], k = 3
// Output: 6
// Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].

// Example 3:
// Input: nums = [2,-3,-1,5,-4], k = 2
// Output: 13
// Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
func TestLargestSumAfterKNegations(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{4, 2, 3}, k: 1, want: 5,
		},
		{
			nums: []int{3, -1, 0, 2}, k: 3, want: 6,
		},
		{
			nums: []int{2, -3, -1, 5, -4}, k: 2, want: 13,
		},
	}
	for _, tc := range testCases {
		got := largestSumAfterKNegations(tc.nums, tc.k)
		if tc.want != got {
			t.Fatalf("TestJump want %v, got %v", tc.want, got)
		}
	}
}
