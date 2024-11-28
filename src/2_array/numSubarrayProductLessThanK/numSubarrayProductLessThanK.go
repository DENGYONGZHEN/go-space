package main

//乘积小于K的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {

	product := 1
	left := 0
	count := 0
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		if right >= left {
			count += right - left + 1
		}
	}
	return count
}
