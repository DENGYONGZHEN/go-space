package backtracking

import "sort"

// 90. Subsets II

// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:
// Input: nums = [1,2,2]
// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

// Example 2:
// Input: nums = [0]
// Output: [[],[0]]

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums) // ⭐️ 必须排序
	result := [][]int{}
	path := []int{}
	k := len(nums)

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		if start == k {
			return
		}

		for i := start; i < k; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return result
}
