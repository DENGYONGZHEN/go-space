package stackandqueue

// 1047. Remove All Adjacent Duplicates In String

// You are given a string s consisting of lowercase English letters.
// A duplicate removal consists of choosing two adjacent and equal letters and removing them.
// We repeatedly make duplicate removals on s until we no longer can.
// Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

func removeDuplicates(s string) string {
	result := make([]rune, len(s))
	pointer := -1

	for _, val := range s {
		if pointer == -1 {
			pointer++
			result[pointer] = val
		} else {
			switch true {
			case val == result[pointer]:
				pointer--
			default:
				pointer++
				result[pointer] = val
			}
		}

	}

	return string(result[:pointer+1])

}
