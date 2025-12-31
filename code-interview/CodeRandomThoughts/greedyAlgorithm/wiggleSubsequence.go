package greedyalgorithm

// 376. Wiggle Subsequence

// A wiggle sequence is a sequence where the differences between successive numbers strictly alternate between positive and negative.
// The first difference (if one exists) may be either positive or negative.
// A sequence with one element and a sequence with two non-equal elements are trivially wiggle sequences.

// For example, [1, 7, 4, 9, 2, 5] is a wiggle sequence because the differences (6, -3, 5, -7, 3) alternate between positive and negative.
// In contrast, [1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences. The first is not because its first two differences are positive,
// and the second is not because its last difference is zero.
// A subsequence is obtained by deleting some elements (possibly zero) from the original sequence,
// leaving the remaining elements in their original order.

// Given an integer array nums, return the length of the longest wiggle subsequence of nums.

// 相邻差值 正负交替
// 第一个差值可以是正或负
// 子序列（不要求连续）
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	//到当前位置为止，最后一步是上升 / 下降的最长摆动子序列长度
	//单个元素本身就是一个合法摆动序列
	up, down := 1, 1

	for i := 1; i < len(nums); i++ {
		//上升，前一步一定是下降。用“之前所有下降状态里最长的那个”，加上当前这一步“上升”
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
		// 相等：什么都不做，跳过即可
	}

	if up > down {
		return up
	}
	return down
}
