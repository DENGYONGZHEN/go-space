package dynamicprogramming

// 70. Climbing Stairs

// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climbStairs(n int) int {

	if n < 4 {
		return n
	}

	pre1, pre2 := 2, 3

	for i := 4; i <= n; i++ {
		current := pre1 + pre2
		pre1 = pre2
		pre2 = current
	}
	return pre2
}
