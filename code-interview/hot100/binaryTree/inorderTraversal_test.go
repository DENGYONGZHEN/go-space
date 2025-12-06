package binarytree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 0, 2, 3},
			want: []int{1, 3, 2},
		},

		{
			nums: []int{1, 2, 3, 4, 5, 0, 8, 0, 0, 6, 7, 9},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
	}

	for _, tc := range testCases {
		root := makeBinaryTree(tc.nums)
		result := inorderTraversal(root)

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("inorderTraversal(%v),should get %v,but get %v", tc.nums, tc.want, result)
		}
	}
}
