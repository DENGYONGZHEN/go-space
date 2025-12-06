package hashtable

import "slices"

// 349. Intersection of Two Arrays

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.

func intersection(nums1 []int, nums2 []int) []int {

	hashtable := make(map[int]bool)
	result := make([]int, 0)

	for _, i := range nums1 {
		hashtable[i] = true
	}
	for _, j := range nums2 {
		if hashtable[j] && slices.Index(result, j) == -1 {
			result = append(result, j)
		}
	}

	return result

}
