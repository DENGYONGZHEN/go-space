package dynamicprogramming

// 509. Fibonacci Number

// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
// such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).

// Dynamic Programming 是一种通过保存子问题结果，来避免重复计算，从而高效求解最优解的方法。
// 1. 最优子结构（Optimal Substructure）。 大问题的最优解，可以由小问题的最优解组合而成
// 2. 重叠子问题（Overlapping Subproblems）。同一个子问题会被反复计算

func fib(n int) int {
	if n < 2 {
		return n
	}

	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}

// func fib(n int) int {

// if n < 2 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }
