package greedyalgorithm

import "sort"

// 1005. Maximize Sum Of Array After K Negations

// Given an integer array nums and an integer k, modify the array in the following way:

// choose an index i and replace nums[i] with -nums[i].
// You should apply this process exactly k times. You may choose the same index i multiple times.

// Return the largest possible sum of the array after modifying it in this way.

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	// 1. 优先翻负数
	for i := 0; i < len(nums) && k > 0 && nums[i] < 0; i++ {
		nums[i] = -nums[i]
		k--
	}

	// 2. 重新排序（保证最小绝对值在前）
	sort.Ints(nums)

	// 3. 如果 k 是奇数，翻最小的
	if k%2 == 1 {
		nums[0] = -nums[0]
	}

	// 4. 求和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
