package binarytree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {

	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			nums: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		{
			nums: []int{},
			want: []int{},
		},
	}

	for _, tc := range testCases {
		result := make([]int, len(tc.nums))
		root := makeBinaryTree(tc.nums)
		index := 0
		root = invertTree(root)
		getNodeVal(root, &result, index)

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("invertTree(%v),should get %v,but get %v", tc.nums, tc.want, result)
		}
	}

}

func getNodeVal(root *TreeNode, result *[]int, index int) {

	if root == nil {
		return
	}
	if index < len(*result) {

		(*result)[index] = root.Val
	}

	getNodeVal(root.Left, result, index*2+1)

	getNodeVal(root.Right, result, index*2+2)

}
