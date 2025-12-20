package backtracking

// 78. Subsets
// Given an integer array nums of unique elements, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Example 2:
// Input: nums = [0]
// Output: [[],[0]]

func subsets(nums []int) [][]int {

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
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return result

}
