package binarytree

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 3

// Example 2:
// Input: root = [1,null,2]
// Output: 2

// func maxDepth(root *TreeNode) int {
// if root == nil{
//     return 0
// }
// 	max, depth := 1, 1

// 	traverseTree(root, &max, &depth)
// 	return max
// }

// func traverseTree(node *TreeNode, max *int, depth *int) {
// 	if node.Left != nil {
// 		*depth = *depth + 1
// 		if *depth > *max {
// 			*max = *depth
// 		}
// 		traverseTree(node.Left, max, depth)
// 		*depth = *depth - 1
// 	}
// 	if node.Right != nil {
// 		*depth = *depth + 1
// 		if *depth > *max {
// 			*max = *depth
// 		}
// 		traverseTree(node.Right, max, depth)
// 		*depth = *depth - 1
// 	}
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
