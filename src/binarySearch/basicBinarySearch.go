package binarysearch

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
