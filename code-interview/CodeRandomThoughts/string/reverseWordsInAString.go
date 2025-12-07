package string

import (
	"slices"
	"strings"
)

// 151. Reverse Words in a String
// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// 菜鸡自己想出来的
func reverseWords(s string) string {

	var stack []string

	for i, j := 0, 0; i < len(s); {
		if s[i] == ' ' {
			if i > 0 && s[i-1] != ' ' {
				stack = append(stack, s[j:i])
			}
			i++
			j++
		} else {
			if i > 0 && s[i-1] == ' ' {
				j = i
			}
			i++

		}

		if i == len(s) && i != j && s[i-1] != ' ' {
			stack = append(stack, s[j:i])
		}
	}

	slices.Reverse(stack)
	return strings.Join(stack, " ")

}

// 优化解法
func reverseWords(s string) string {
	var words []string
	n := len(s)

	for i, j := 0, 0; i < n; {
		if s[i] == ' ' {
			i++
			j = i
			continue
		}

		// 找单词结束
		for i < n && s[i] != ' ' {
			i++
		}
		words = append(words, s[j:i])
		j = i
	}

	// 翻转
	slices.Reverse(words)
	return strings.Join(words, " ")
}
