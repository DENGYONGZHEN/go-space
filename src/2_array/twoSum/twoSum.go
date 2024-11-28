package twoSum

import "sort"

//两数之和

// 双指针 有序数组(递增或递减)
func TwoSum(nums []int, target int) []int {
	sort.Ints(nums)
	var left, right = 0, len(nums) - 1
	for left < right {

		sum := nums[left] + nums[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left, right}
		}

	}
	return nil
}

//TODO
//双指针 无序数组 对n个元素中的每一个元素k分别遍历n次找到target - k 的元素

//TODO
//把元素存入map中，针对每个元素，找个target - k
