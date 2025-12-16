package stackandqueue

// 239. Sliding Window Maximum
// You are given an array of integers nums,
// there is a sliding window of size k which is moving from the very left of the array to the very right.
// You can only see the k numbers in the window. Each time the sliding window moves right by one position.
// Return the max sliding window.

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 {
		return nil
	}

	//记录nums中元素的索引
	deque := []int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		if len(deque) > 0 && deque[0] <= i-k {
			//当队列的头位置的元素被划出窗口外，就移除
			deque = deque[1:]
		}

		// 2. 维护单调递减
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// 4. 记录结果
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
