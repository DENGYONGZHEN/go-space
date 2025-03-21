package array

// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

// Example 1:
// Input: n = 3
// Output: [[1,2,3],[8,9,4],[7,6,5]]

// Example 2:
// Input: n = 1
// Output: [[1]]

// 解题思路：就是一个正方形，环绕旋转
func generateMatrix(n int) [][]int {

	//创建对应的二维slice
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	rowV, columnV := 0, 0
	offset := 1 //控制循环的边界
	count := 1

	// 边界取左闭右开
	for range n / 2 {

		//向右遍历
		for columnV < n-offset {
			result[rowV][columnV] = count
			count++
			columnV++
		}
		//向下遍历
		for rowV < n-offset {
			result[rowV][columnV] = count
			count++
			rowV++
		}
		//向左遍历
		for columnV > offset-1 {
			result[rowV][columnV] = count
			count++
			columnV--
		}

		//向上遍历
		for rowV > offset-1 {
			result[rowV][columnV] = count
			count++
			rowV--
		}
		//每一轮遍历完，更新边界
		rowV++
		columnV++
		offset += 1

	}

	//奇数时，赋值最后一个值
	if n%2 == 1 {
		result[n/2][n/2] = n * n
	}
	return result
}
