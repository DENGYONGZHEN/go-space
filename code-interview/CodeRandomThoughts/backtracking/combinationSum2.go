package backtracking

import "sort"

// 40. Combination Sum II

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Note: The solution set must not contain duplicate combinations.

func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	result := [][]int{}

	path := []int{}
	sum := 0

	var backtrack func(start, sum int)

	backtrack = func(start, sum int) {

		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				result = append(result, temp)
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			//关键去重，每层都要去重
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			backtrack(i+1, sum)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	backtrack(0, sum)
	return result
}
