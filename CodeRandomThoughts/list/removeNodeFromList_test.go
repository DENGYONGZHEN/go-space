package list

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {

	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{1, 2, 6, 3, 4, 5, 6},
			target: 6,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			nums:   []int{},
			target: 1,
			want:   []int{},
		},
		{
			nums:   []int{7, 7, 7, 7},
			target: 7,
			want:   []int{},
		},
	}

	for _, tc := range testCases {

		if len(tc.nums) == 0 {
			val := removeElements(nil, tc.target)
			if val != nil {
				t.Errorf(" removeElements(nil,%d),should get %v, but get %v", tc.target, tc.want, val)
			}
		}
		head := &ListNode{}
		current := head

		for _, v := range tc.nums {
			current.Next = &ListNode{
				Val: v,
			}
			current = current.Next
		}
		val := removeElements(head.Next, tc.target)
		current.Next = val
		result := []int{}

		for current.Next != nil {
			result = append(result, current.Next.Val)
			current = current.Next
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf(" removeElements(nil,%d),should get %v, but get %v", tc.target, tc.want, result)
		}

	}
}
