package greedyalgorithm

import "testing"

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{2, 3, 1, 1, 4}, want: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4}, want: false,
		},
	}
	for _, tc := range testCases {
		got := canJump(tc.nums)
		if got != tc.want {
			t.Fatalf("TestCanJump want %v,got %v", tc.want, got)
		}
	}
}
