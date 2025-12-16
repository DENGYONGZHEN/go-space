package stackandqueue

// 347. Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

type Pair struct {
	num  int
	freq int
}

// ---------- 最小堆 ----------
// type MinHeap []Pair

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq } // 频率小的在前
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(Pair))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// 最小堆的解法
// func topKFrequent(nums []int, k int) []int {
// 	freq := make(map[int]int)
// 	for _, num := range nums {
// 		freq[num]++
// 	}

// 	h := &MinHeap{}
// 	heap.Init(h)

// 	for num, count := range freq {
// 		if h.Len() < k {
// 			heap.Push(h, Pair{num, count})
// 		} else if count > (*h)[0].freq {
// 			heap.Pop(h)
// 			heap.Push(h, Pair{num, count})
// 		}
// 	}

// 	result := make([]int, 0, k)
// 	for h.Len() > 0 {
// 		result = append(result, heap.Pop(h).(Pair).num)
// 	}
// 	return result
// }

//桶排序的解法

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// buckets[i] 存出现 i 次的数字
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}
