package backtracking

// 491. Non-decreasing Subsequences

// Given an integer array nums, return all the different possible non-decreasing subsequences of the given array with at least two elements.
// You may return the answer in any order.

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	k := len(nums)

	var backtrack func(start int)
	backtrack = func(start int) {

		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		// ⭐ 同一层去重
		used := map[int]bool{}

		for i := start; i < k; i++ {

			// ⭐ 非递减剪枝（选之前）
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}

			// ⭐ 491 的正确去重方式
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true

			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
