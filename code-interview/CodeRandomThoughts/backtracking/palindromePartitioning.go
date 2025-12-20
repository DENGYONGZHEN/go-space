package backtracking

// 131. Palindrome Partitioning

// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.

func partition(s string) [][]string {

	result := [][]string{}
	path := []string{}
	k := len(s)
	var backtrack func(start int, path []string)

	backtrack = func(start int, path []string) {
		if start == k {
			allPalindrome := true
			for _, s := range path {
				if palindrome := checkIsPalindrome(s); !palindrome {
					allPalindrome = false
					break
				}
			}
			if allPalindrome {
				temp := make([]string, len(path))
				copy(temp, path)
				result = append(result, temp)
			}
			return
		}

		for i := start; i < k; i++ {
			path = append(path, s[start:i+1])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, path)
	return result

}

func checkIsPalindrome(s string) bool {

	if len(s) == 0 {
		return false
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
