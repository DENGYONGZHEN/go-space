package backtracking

import "strings"

// 17. Letter Combinations of a Phone Number

// Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// 2(a,b,c)
// 3(d,e,f)
// 4(g,h,i)
// 5(j,k,l)
// 6(m,n,o)
// 7(p,q,r,s)
// 8(t,u,v)
// 9(w,x,y,z)

func letterCombinations(digits string) []string {

	mapping := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	mapped := make([][]string, len(digits)+1)
	for i, dig := range digits {
		mapped[i] = mapping[dig]
	}
	mapped[len(mapped)-1] = []string{}

	result := []string{}
	k := len(digits)
	path := []string{}
	var backtrack func(start int, arr []string)

	backtrack = func(start int, arr []string) {

		if len(path) == k {
			result = append(result, strings.Join(path, ""))
			return
		}

		for i := 0; i < len(arr); i++ {
			path = append(path, arr[i])
			backtrack(start+1, mapped[start+1])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, mapped[0])

	return result

}
