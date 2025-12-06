package binarytree

import "testing"

func TestMaxDepth(t *testing.T) {

	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 9, 20, 0, 0, 15, 7},
			want: 3,
		},

		{
			nums: []int{1, 0, 2},
			want: 2,
		},
	}

	for _, tc := range testCases {
		root := makeBinaryTree(tc.nums)
		result := maxDepth(root)

		if result != tc.want {
			t.Errorf("maxDepth(%v),should get %v ,but get %v", tc.nums, tc.want, result)
		}
	}
}
