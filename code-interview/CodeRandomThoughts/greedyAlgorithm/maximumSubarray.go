package greedyalgorithm

// 53. Maximum Subarray
// Given an integer array nums, find the subarray with the largest sum, and return its sum.

// 以当前位置结尾的最大子数组和，Kadane 算法
// 固定「结尾」，不固定「起点」

func maxSubArray(nums []int) int {
	cur := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		//一旦高度跌到负数，就从当前位置重新开始爬
		//每一步都问自己——前面的和，还值不值得接？
		cur = max(nums[i], cur+nums[i])
		maxSum = max(maxSum, cur) //只在当前值大于最大值时改变并记录
	}

	return maxSum
}
