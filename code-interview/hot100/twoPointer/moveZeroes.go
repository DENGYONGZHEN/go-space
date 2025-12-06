package twopointer

// Given an integer array nums,
// move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

func moveZeroes(nums []int) {

	var fast, slow int

	for fast < len(nums) && slow < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}
		nums[fast], nums[slow] = nums[slow], nums[fast]
		slow++
		fast++
	}

}
