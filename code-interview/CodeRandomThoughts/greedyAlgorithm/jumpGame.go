package greedyalgorithm

// 55. Jump Game

// You are given an integer array nums. You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		// 如果当前位置已经跳不到了，直接失败
		if i > maxReach {
			return false
		}

		// 更新最远可达位置
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}

	return true
}
