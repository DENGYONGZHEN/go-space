package string

// 541. Reverse String II
// Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.
// If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters,
// then reverse the first k characters and leave the other as original.

func reverseStr(s string, k int) string {
	b := []byte(s)
	n := len(b)

	for i := 0; i < n; i += 2 * k {
		// 判断剩余长度是否 >= k
		if i+k <= n {
			reverse(b[i : i+k])
		} else {
			// 剩下的不足 k 全部反转
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
