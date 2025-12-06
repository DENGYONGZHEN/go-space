package array

import "sort"

// 18. 4Sum
// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.

func fourSum(nums []int, target int) [][]int {

	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		//第一重去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			//第二重去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// 跳过 left 重复值
					leftVal, rightVal := nums[left], nums[right]
					for left < right && nums[left] == leftVal {
						left++
					}
					// 跳过 right 重复值
					for left < right && nums[right] == rightVal {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}

		}
	}
	return result
}
