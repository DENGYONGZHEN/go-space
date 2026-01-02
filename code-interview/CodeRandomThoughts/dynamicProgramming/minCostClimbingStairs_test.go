package dynamicprogramming

import "testing"

// Example 1:
// Input: cost = [10,15,20]
// Output: 15
// Explanation: You will start at index 1.
// - Pay 15 and climb two steps to reach the top.
// The total cost is 15.

// Example 2:
// Input: cost = [1,100,1,1,1,100,1,1,100,1]
// Output: 6
// Explanation: You will start at index 0.
// - Pay 1 and climb two steps to reach index 2.
// - Pay 1 and climb two steps to reach index 4.
// - Pay 1 and climb two steps to reach index 6.
// - Pay 1 and climb one step to reach index 7.
// - Pay 1 and climb two steps to reach index 9.
// - Pay 1 and climb one step to reach the top.
// The total cost is 6.
func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		cost []int
		want int
	}{
		{
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}

	for _, tc := range testCases {
		got := minCostClimbingStairs(tc.cost)

		if got != tc.want {
			t.Fatalf("TestMinCostClimbingStairs want %v, got %v", tc.want, got)
		}
	}

}
