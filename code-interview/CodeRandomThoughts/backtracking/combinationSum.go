package backtracking

// 39. Combination Sum

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers sum to target.
// You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

func combinationSum(candidates []int, target int) [][]int {

	result := [][]int{}
	path := []int{}
	sum := 0
	var backtrack func(start, sum int)

	backtrack = func(start, sum int) {
		if sum >= target {
			if sum == target {
				tem := make([]int, len(path))
				copy(tem, path)
				result = append(result, tem)
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtrack(i, sum)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtrack(0, sum)
	return result
}
