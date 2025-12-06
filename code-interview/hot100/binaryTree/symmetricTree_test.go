package binarytree

import (
	"testing"
)

func TestIsSymmetricTree(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			nums: []int{1, 2, 2, 0, 3, 0, 3},
			want: false,
		},
	}

	for _, tc := range testCases {
		root := &TreeNode{
			Val: tc.nums[0],
		}
		makeSymmetricTree(tc.nums, root, 0)
		val := isSymmetric(root)
		if val != tc.want {
			t.Errorf("isSymmetric(%v) should %v ,but get %v", tc.nums, tc.want, val)
		}
	}
}

func makeSymmetricTree(nums []int, node *TreeNode, level int) {

	valIndex := level*2 + 1
	if valIndex < len(nums) {
		node.Left = &TreeNode{
			Val: nums[valIndex],
		}
		makeSymmetricTree(nums, node.Left, valIndex)
	}
	valIndex = level*2 + 2
	if valIndex < len(nums) {
		node.Right = &TreeNode{
			Val: nums[valIndex],
		}
		makeSymmetricTree(nums, node.Right, valIndex)
	}

}
