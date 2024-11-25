package threeSumZero

import (
	"sort"
)

//在一个数组中找到三数之和的值为0的所有组合
//1.数组排序
//2.给定i，i>=0 并且i <len (nums) -2 的数，使用两个指针，left为i+，right为len(nums)-1，进行匹配
//3.注意要跳过重复的数
//  1.在匹配到3数之和为0后，left向右跳过相同的数
//  2.内循环结束后，i向右跳过相同的数

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				skipLeft := nums[left]
				for left < right && skipLeft == nums[left] {
					left++
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
		skipI := nums[i]
		for i < len(nums) && skipI == nums[i] {
			i++
		}
	}
	return result
}
