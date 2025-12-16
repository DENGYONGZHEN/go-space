package backtracking

// 77. Combinations
// Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
// You may return the answer in any order.

func combine(n int, k int) [][]int {
	var res [][]int
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 剪枝
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			//递归
			backtrack(i + 1)
			//回溯
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return res
}
