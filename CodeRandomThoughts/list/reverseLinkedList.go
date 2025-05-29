package list

//206. Reverse Linked List
//  1 -> 2 -> 3 -> 4 -> 5
//  5 -> 4 -> 3 -> 2 -> 1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法：依次遍历每个节点，插入到dummyHead的后面
// func reverseList(head *ListNode) *ListNode {

// 	dummyHead := &ListNode{}
// 	if head == nil {
// 		return nil
// 	}

// 	currentNode := head

// 	for currentNode != nil {
// 		next := currentNode.Next
// 		currentNode.Next = dummyHead.Next
// 		dummyHead.Next = currentNode
// 		currentNode = next
// 	}

// 	return dummyHead.Next
// }

// 双指针
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	var pre *ListNode

	currentNode := head

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = pre
		pre = currentNode
		currentNode = next
	}

	return pre
}

// 递归
