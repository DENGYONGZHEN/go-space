package backtracking

// 46. Permutations

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

func permute(nums []int) [][]int {
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
