package backtracking

//Given an array nums of distinct integers, return all the possible permutations.
//  You can return the answer in any order.

//Example 1:

//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//Example 2:

//Input: nums = [0,1]
//Output: [[0,1],[1,0]]

// Example 3:
// Input: nums = [1]
// Output: [[1]]

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(&res, nums, 0)
	return res
}

func backtrack(res *[][]int, nums []int, start int) {
	if start == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrack(res, nums, start+1)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
