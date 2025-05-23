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
func reverseList(head *ListNode) *ListNode {

	dummyHead := &ListNode{}
	if head == nil {
		return nil
	}

	currentNode := head

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = dummyHead.Next
		dummyHead.Next = currentNode
		currentNode = next
	}

	return dummyHead.Next
}
