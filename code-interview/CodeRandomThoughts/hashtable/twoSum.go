package hashtable

// 1. Two Sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {

	hashtable := make(map[int]int)
	for index, val := range nums {
		if v, ok := hashtable[target-val]; ok {
			return []int{v, index}
		}
		hashtable[val] = index
	}
	return nil

}

// 其实也可以不用哈希表，用双指针也可以解决
// func twoSum(nums []int, target int) []int {

// 	sort.Ints(nums)
// 	left, right := 0, len(nums)-1
// 	for left < right {
// 		if nums[left]+nums[right] > target {
// 			right--
// 		} else if nums[left]+nums[right] < target {
// 			left++
// 		} else {
// 			return []int{left, right}
// 		}
// 	}
// 	return nil
// }
