package list

// 19. Remove Nth Node From End of List
//Given the head of a linked list, remove the nth node from the end of the list and return its head.

// 思路，设置两个指针，第二个指针要快于第一个指针 N 步，当第二个指针到达终点时，说明第一个指针也到了倒数第N的位置
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	//虚拟头节点
	slow := &ListNode{
		Next: head,
	}
	fast := head
	//快节点先走n步
	for range n {
		if fast != nil {
			fast = fast.Next
		}
	}
	//同时移动快慢指针，直到快节点为nil时
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//如果慢节点未被移动，表明条件不满足
	if slow.Next == head {
		return slow.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}
