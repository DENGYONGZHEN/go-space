package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	testCases := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
	}

	for _, tc := range testCases {

		dummyhead1 := &ListNode{}
		current1 := dummyhead1
		for _, v := range tc.list1 {
			current1.Next = &ListNode{
				Val: v,
			}
			current1 = current1.Next
		}

		dummyhead2 := &ListNode{}
		current2 := dummyhead2
		for _, v := range tc.list2 {
			current2.Next = &ListNode{
				Val: v,
			}
			current2 = current2.Next
		}
		val := mergeTwoLists(dummyhead1.Next, dummyhead2.Next)

		result := []int{}
		for val != nil {
			result = append(result, val.Val)
			val = val.Next
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("mergeTwoLists(%v, %v),should %v,but get %v", tc.list1, tc.list2, tc.want, result)
		}
	}

}
