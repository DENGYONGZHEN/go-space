package main

import "sort"

func BasicBinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	return binarySearch(nums, target)
}

func binarySearch(nums []int, target int) int {

	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	mid := len(nums) / 2

	if nums[mid] == target {
		return nums[mid]
	} else if nums[mid] > target {
		return binarySearch(nums[:mid], target)
	} else {
		return binarySearch(nums[mid+1:], target)
	}
}
