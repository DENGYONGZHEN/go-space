package greedyalgorithm

// 45. Jump Game II

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at index 0.
// Each element nums[i] represents the maximum length of a forward jump from index i.
// In other words, if you are at index i, you can jump to any index (i + j) where:

// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach index n - 1. The test cases are generated such that you can reach index n - 1.

func jump(nums []int) int {
	//长度为小于等于1时，不用跳，返回0
	if len(nums) <= 1 {
		return 0
	}

	step := 0
	maxReach := 0
	curEnd := 0

	for i := 0; i < len(nums)-1; i++ {
		//获取最远到达
		maxReach = max(maxReach, i+nums[i])

		//在一个段内只计算一步。到达这一段的最后才会更新下一段的终点
		if i == curEnd {
			step++
			curEnd = maxReach
		}
	}

	return step
}
