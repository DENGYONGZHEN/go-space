package binarysearch

// Given an array of integers nums sorted in non-decreasing order,
// find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]

// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]

func searchRange(nums []int, target int) []int {

	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			//就是正常的二分查找，然后再找到一个的时候，分别向左右两边扩散，
			//找到开始位置和结束位置
			first, last := middle, middle
			for first > 0 && nums[first-1] == target {
				first--
			}
			for last < len(nums)-1 && nums[last+1] == target {
				last++
			}
			return []int{first, last}

		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return result
}
