package list

// 142. Linked List Cycle II
// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed).
// It is -1 if there is no cycle. Note that pos is not passed as a parameter.
// Do not modify the linked list.

// 思路： 设置快慢指针，快指针总是走两步，慢指针走一步。如果有环，快慢指针一定会在环内某个节点上相遇，因为每次移动，相当于快指针移动一步追慢指针
// 假设从头节点到环的入口为a,环的入口到快慢指针相遇的地方为b，相遇的地方再到入口为从，然后 慢指针所走的为 a + b ，快指针为 a + b + n * (c + b) ,n为快指针走的环的圈数
// 所以 2 *(a + b ) = a + b + n * (c + b) ,所以 a  = (n-1)(b+c) + c,当n为1时，a = c 所以把慢指针放在头指针继续走，一定会在入口遇到走一步的快指针
// 即使n不为1，也表示从头节点到环的入口距离会等于转n圈后再走c的距离
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	//先判断是否有环
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	//如果快指针为nil或者下一个节点为nil，说明没环
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 有环的话，就将慢指针移到头节点
	slow = head
	for slow != fast {
		//快慢指针同时移动一步向前走，相遇时，就是环的入口
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
