package binarysearch

import "math"

// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//确定nums1是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		partition1 := (left + right) / 2
		partition2 := (m+n+1)/2 - partition1

		// 计算左右边界
		maxLeft1 := math.MinInt
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		minRight1 := math.MaxInt
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := math.MinInt
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight2 := math.MaxInt
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 1 {
				return float64(max(maxLeft1, maxLeft2))
			}
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
		} else if maxLeft1 > minRight2 {
			right = partition1 - 1
		} else {
			left = partition1 + 1
		}
	}
	return float64(0)
}
