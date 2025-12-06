package list

import (
	"reflect"
	"testing"
)

func TestSwapNodesInPairs(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1, 2, 3},
			want: []int{2, 1, 3},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
	}

	for _, tc := range testCases {
		head := makeLinkedList(tc.nums)
		result := collectListNodeValue(swapPairs(head))
		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("want %v,get %v", tc.want, result)
		}
	}
}
