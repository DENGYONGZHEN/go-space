package string

// 151. Reverse Words in a String
// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// 菜鸡自己想出来的
// func reverseWords(s string) string {

// 	var stack []string

// 	for i, j := 0, 0; i < len(s); {
// 		if s[i] == ' ' {
// 			if i > 0 && s[i-1] != ' ' {
// 				stack = append(stack, s[j:i])
// 			}
// 			i++
// 			j++
// 		} else {
// 			if i > 0 && s[i-1] == ' ' {
// 				j = i
// 			}
// 			i++

// 		}

// 		if i == len(s) && i != j && s[i-1] != ' ' {
// 			stack = append(stack, s[j:i])
// 		}
// 	}

// 	slices.Reverse(stack)
// 	return strings.Join(stack, " ")

// }

// 优化解法
// func reverseWords(s string) string {
// 	var words []string
// 	n := len(s)

// 	for i, j := 0, 0; i < n; {
// 		if s[i] == ' ' {
// 			i++
// 			j = i
// 			continue
// 		}

// 		// 找单词结束
// 		for i < n && s[i] != ' ' {
// 			i++
// 		}
// 		words = append(words, s[j:i])
// 		j = i
// 	}

// 	// 翻转
// 	slices.Reverse(words)
// 	return strings.Join(words, " ")
// }

// 标准双指针解法
func reverseWords(s string) string {
	// 转成 byte slice 可原地修改
	b := []byte(s)

	// 1. 去掉多余空格
	b = trimSpaces(b)

	// 2. 整体反转
	reverseStrings(b, 0, len(b)-1)

	// 3. 逐单词反转
	reverseEachWord(b)

	return string(b)
}

func trimSpaces(b []byte) []byte {
	slow := 0
	fast := 0
	n := len(b)

	// 跳过前导空格
	for fast < n && b[fast] == ' ' {
		fast++
	}

	for fast < n {
		// 如果是空格，且下一个还是空格，则跳过
		if fast > 0 && b[fast] == ' ' && b[fast-1] == ' ' {
			fast++
			continue
		}
		b[slow] = b[fast]
		slow++
		fast++
	}

	// 去掉尾部空格
	if slow > 0 && b[slow-1] == ' ' {
		return b[:slow-1]
	}

	return b[:slow]
}

func reverseStrings(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseEachWord(b []byte) {
	n := len(b)
	start := 0
	end := 0

	for start < n {
		// 找到单词末尾
		for end < n && b[end] != ' ' {
			end++
		}
		// 反转这个单词
		reverseStrings(b, start, end-1)
		// 移动到下一个单词
		end++
		start = end
	}
}
