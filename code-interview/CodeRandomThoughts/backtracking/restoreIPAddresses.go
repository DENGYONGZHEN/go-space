package backtracking

import (
	"strconv"
	"strings"
)

// 93. Restore IP Addresses

// A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
// Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s.
// You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.

func restoreIpAddresses(s string) []string {

	result := []string{}
	path := []string{}
	k := len(s)

	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {

		if start == k || len(path) == 4 {
			if start == k && len(path) == 4 {
				allPartCorrect := true
				//判定每个部分是不是都合规
				for _, part := range path {
					if len(part) > 1 && part[0] == '0' {
						allPartCorrect = false
						break
					}
					if val, _ := strconv.Atoi(part); val > 255 {
						allPartCorrect = false
						break
					}
				}
				if allPartCorrect {
					result = append(result, strings.Join(path, "."))
				}
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
