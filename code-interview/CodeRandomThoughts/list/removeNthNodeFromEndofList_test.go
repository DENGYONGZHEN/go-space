package list

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		nums []int
		n    int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			nums: []int{1},
			n:    1,
			want: []int{},
		},
		{
			nums: []int{1, 2},
			n:    1,
			want: []int{1},
		},
	}

	for _, tc := range testCases {
		head := makeLinkedList(tc.nums)
		result := collectListNodeValue(removeNthFromEnd(head, tc.n))
		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("want %v,get %v", tc.want, result)
		}
	}
}
