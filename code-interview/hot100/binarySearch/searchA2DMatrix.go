package binarysearch

//You are given an m x n integer matrix matrix with the following two properties:
//Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.
// You must write a solution in O(log(m * n)) time complexity.
// Example 1:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

// Example 2:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		middle := (left + right) / 2

		//注意是除以内数组的长度得到row
		row, column := middle/n, middle%n
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false
}
