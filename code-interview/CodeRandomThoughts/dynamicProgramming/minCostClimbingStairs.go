package dynamicprogramming

// 746. Min Cost Climbing Stairs

// You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the floor.

// 1.起始站位  刚开始的站位可以选择是 0 或 1
// 2.付费规则  踏到哪个台阶，付哪个台阶的钱
// 3.顶点（top）顶点是不付钱的
func minCostClimbingStairs(cost []int) int {
	prev2, prev1 := 0, 0

	for i := 2; i <= len(cost); i++ {
		cur := min(
			prev1+cost[i-1],
			prev2+cost[i-2],
		)
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
