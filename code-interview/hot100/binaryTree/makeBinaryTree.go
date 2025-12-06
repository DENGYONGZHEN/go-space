package binarytree

import "fmt"

// 从数组构建二叉树（最常用方法）
func makeBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if index < len(nums) && nums[index] != 0 {
			node.Left = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Left)
		}
		index++

		// 处理右子节点
		if index < len(nums) && nums[index] != 0 {
			node.Right = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// 打印树结构（辅助函数）
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}
}
