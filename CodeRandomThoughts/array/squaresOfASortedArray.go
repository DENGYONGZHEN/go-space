package array

//Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.

//Example 1:

// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:

// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

func sortedSquares(nums []int) []int {

	//新数组的下标
	resultIndex := len(nums) - 1
	result := make([]int, len(nums))

	for left, right := 0, len(nums)-1; left <= right; {
		temLeftV := nums[left] * nums[left]
		temRightV := nums[right] * nums[right]

		//最大值肯定出现在数组的两边，用双指针从两边向内遍历，比较两指针指向的值的平方
		if temLeftV < temRightV {
			result[resultIndex] = temRightV
			right--
		} else {
			result[resultIndex] = temLeftV
			left++
		}

		resultIndex--
	}
	return result
}
