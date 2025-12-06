package linkedlist

import "testing"

func TestHasCycle(t *testing.T) {

	testCases := []struct {
		list1 []int
		pos   int
		want  bool
	}{
		{
			list1: []int{3, 2, 0, -4},
			pos:   1,
			want:  true,
		},
		{
			list1: []int{1, 2},
			pos:   0,
			want:  true,
		},
		{
			list1: []int{1},
			pos:   -1,
			want:  false,
		},
	}

	for _, tc := range testCases {

		head := makeList(tc.list1, tc.pos)
		val := hasCycle(head)
		if val != tc.want {
			t.Errorf("hasCycle(%v),want %v,but %v", tc.list1, tc.want, val)
		}
	}

}

func makeList(list []int, pos int) *ListNode {

	dummyhead := &ListNode{}
	current := dummyhead
	count, node := 0, &ListNode{}
	for _, val := range list {
		current.Next = &ListNode{
			Val: val,
		}
		if pos == count {
			node = current.Next
		}
		current = current.Next
		count++
	}

	if pos >= 0 {
		current.Next = node
	}

	return dummyhead.Next

}
