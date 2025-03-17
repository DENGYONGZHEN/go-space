package array

func removeElement(nums []int, val int) int {

	//使用双指针，快慢指针，快指针向前寻找需要留下来的数据
	//满指针定位在新数组的最后的位置
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow

}
