package backtracking

func generateParenthesis(n int) []string {
	var result []string
	backtracks(&result, "", 0, 0, n)
	return result
}

// 回溯函数
func backtracks(result *[]string, current string, open int, close int, max int) {
	// 如果当前组合长度等于最大长度的两倍，添加到结果中
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// 如果左括号数量小于最大值，可以添加左括号
	if open < max {
		backtracks(result, current+"(", open+1, close, max)
	}

	// 如果右括号数量小于左括号数量，可以添加右括号
	if close < open {
		backtracks(result, current+")", open, close+1, max)
	}
}
