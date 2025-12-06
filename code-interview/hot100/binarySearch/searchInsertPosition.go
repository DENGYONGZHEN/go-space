package binarysearch

import "fmt"

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found. If not,
// return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2

// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1

// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

func searchInsert(nums []int, target int) int {

	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	left, right := 0, len(nums)-1
	//这个问题关键就是middle的最后的位置，只要比较middle指向的值与target的大小
	//就可以得到答案
	middle := -1
	for left <= right {
		middle = (left + right) / 2
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	fmt.Printf("middle is %v", middle)
	if nums[middle] >= target {
		return middle
	} else {
		return middle + 1
	}

}
