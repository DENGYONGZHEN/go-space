package backtracking

//Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent.
// Return the answer in any order.

//A mapping of digits to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.

//Example 1:

//Input: digits = "23"
//Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//Example 2:

//Input: digits = ""
//Output: []
//Example 3:

//Input: digits = "2"
//Output: ["a","b","c"]

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 数字到字母的映射
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	backtracking(&result, digits, "", 0, digitToLetters)
	return result
}

// 回溯函数
func backtracking(result *[]string, digits string, current string, index int, digitToLetters map[byte]string) {
	// 如果当前组合长度等于输入数字长度，添加到结果中
	if len(current) == len(digits) {
		*result = append(*result, current)
		return
	}

	// 获取当前数字对应的字母
	letters := digitToLetters[digits[index]]
	for i := 0; i < len(letters); i++ {
		// 递归生成组合
		backtracking(result, digits, current+string(letters[i]), index+1, digitToLetters)
	}
}
