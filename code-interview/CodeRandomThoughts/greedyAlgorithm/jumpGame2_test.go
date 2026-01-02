package greedyalgorithm

import "testing"

// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:

// Input: nums = [2,3,0,1,4]
// Output: 2
func TestJump(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4}, want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4}, want: 2,
		},
	}
	for _, tc := range testCases {
		got := jump(tc.nums)
		if tc.want != got {
			t.Fatalf("TestJump want %v, got %v", tc.want, got)
		}
	}
}
