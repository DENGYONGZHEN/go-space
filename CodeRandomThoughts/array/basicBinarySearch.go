package array

//Given an array of integers nums which is sorted in ascending order,
// and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.

//You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

func search(nums []int, target int) int {

	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	//1.边界的处理规则是左闭右闭 [1,1]
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			//已知 nums[middle] > target，所以区间的右边界应该是middle-1
			right = middle - 1
			continue
		} else {
			//类似的，已知 nums[middle] < target，所以区间的右边界应该是middle+1
			left = middle + 1
		}
	}

	//2.边界的处理规则是左闭右开 [1,1）
	// left, right := 0, len(nums)

	// for left < right {
	// 	middle := (left + right) / 2
	// 	if nums[middle] == target {
	// 		return middle
	// 	} else if nums[middle] > target {
	// 		right = middle
	// 	} else {
	// 		left = middle + 1 //左闭右开
	// 	}
	// }

	return -1
}
