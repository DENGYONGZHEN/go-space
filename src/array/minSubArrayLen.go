package array

//209. Minimum Size Subarray Sum

//输入一个正整数组成的数组，和一个正整数k，输出数组中和大于等于k的连续子数组的最短长度

func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := n + 1 // 初始化为比数组长度大的值
	sum := 0           // 当前窗口的总和
	left := 0          // 滑动窗口的左边界

	for right := 0; right < n; right++ {
		// 增加当前右边界的值
		sum += nums[right]

		// 当窗口内的和满足条件时，尝试收缩左边界
		for sum >= target {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	// 如果没有找到符合条件的子数组，返回 0
	if minLength == n+1 {
		return 0
	}
	return minLength
}

// 辅助函数：取较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
