package list

//24. Swap Nodes in Pairs
//Given a linked list, swap every two adjacent nodes and return its head.
//  You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// example
// 1 -> 2 -> 3 -> 4
// 2 -> 1 -> 4 -> 3

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	mid, right := head, head.Next
	if right == nil {
		return mid
	}
	dummyHead := &ListNode{}
	left := dummyHead
	for mid != nil && right != nil {
		temp := right.Next
		right.Next = nil
		mid.Next = nil
		left.Next = right
		right.Next = mid
		left = mid
		mid = temp
		if mid != nil {
			right = mid.Next
			if right == nil {
				left.Next = mid
			}
		}

	}
	return dummyHead.Next

}
