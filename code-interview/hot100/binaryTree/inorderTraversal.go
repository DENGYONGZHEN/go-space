package binarytree

// Input: root = [1,null,2,3]
// Output: [1,3,2]

// Example 2:
// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
// Output: [4,2,6,5,7,1,3,9,8]

// Example 3:
// Input: root = []
// Output: []

// Example 4:
// Input: root = [1]
// Output: [1]

// func inorderTraversal(root *TreeNode) []int {

// 	result := []int{}
// 	if root == nil {
// 		return result
// 	}
// 	result = traverse(root, result)
// 	return result

// }

// func traverse(node *TreeNode, result []int) []int {

// 	if node.Left != nil {
// 		result = traverse(node.Left, result)
// 	}
// 	result = append(result, node.Val)
// 	if node.Right != nil {
// 		result = traverse(node.Right, result)
// 	}
// 	return result
// }

func inorderTraversal(root *TreeNode) []int {

	//注意 使用 var re1 []int  和 re2 := []int{}， re1不等于re2，re1
	// re1 == nil → true
	// re2 == nil → false
	// len(re1) = 0
	// len(re2) = 0
	result := []int{}
	travers(root, &result)
	return result
}

func travers(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	travers(root.Left, result)
	*result = append(*result, root.Val)
	travers(root.Right, result)
}
