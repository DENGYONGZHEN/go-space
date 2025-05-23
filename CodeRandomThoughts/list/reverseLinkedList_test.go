package list

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		{
			nums: []int{},
			want: []int{},
		},
	}

	for _, ts := range testCases {
		result := reverseList(makeLinkedList(ts.nums))
		resultCollected := collectListNodeValue(result)
		if !reflect.DeepEqual(ts.want, resultCollected) {
			t.Fatalf("result: %v, want: %v", resultCollected, ts.want)
		}
	}

}

func makeLinkedList(input []int) *ListNode {

	if len(input) <= 0 {
		return nil
	}

	head := &ListNode{Val: input[0]}
	current := head

	for i := 1; i < len(input); i++ {
		current.Next = &ListNode{Val: input[i]}
		current = current.Next
	}
	return head
}

func collectListNodeValue(head *ListNode) []int {

	if head == nil {
		return []int{}
	}
	current := head
	result := make([]int, 0)
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
