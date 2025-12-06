package hashtable

// 242. Valid Anagram
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// m := make(map[rune]int, len(s))

	// for i, r := range s {
	// 	m[r]++
	// 	m[rune(t[i])]--
	// }

	// for _, value := range m {
	// 	if value != 0 {
	// 		return false
	// 	}
	// }

	// return true

	//还可以使用数组实现,总共26个字母，以a为基准，从0索引开始map相应的字母
	hashtable := make([]int, 26)
	for _, r := range s {
		hashtable[(r-'a')]++
	}
	for _, u := range t {
		hashtable[(u-'a')]--
	}
	for _, f := range hashtable {
		if f != 0 {
			return false
		}
	}
	return true
}
