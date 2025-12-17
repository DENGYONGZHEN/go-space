package backtracking

// 216. Combination Sum III

// Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

// Only numbers 1 through 9 are used.
// Each number is used at most once.
// Return a list of all possible valid combinations.
// The list must not contain the same combination twice, and the combinations may be returned in any order.

func combinationSum3(k int, n int) [][]int {

	result := [][]int{}
	elem := []int{}
	sum := 0

	var backtrack func(start int, sum int)
	backtrack = func(start int, sum int) {

		if len(elem) == k {
			if sum == n {
				temp := make([]int, k)
				copy(temp, elem)
				result = append(result, temp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			sum += i
			elem = append(elem, i)
			backtrack(i+1, sum)
			elem = elem[:len(elem)-1]
			sum -= i
		}
	}

	backtrack(1, sum)
	return result
}
