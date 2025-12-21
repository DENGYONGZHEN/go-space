package backtracking

import "sort"

// 47. Permutations II

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// 终止条件：用完所有数
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// ⭐ 同一层去重
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
