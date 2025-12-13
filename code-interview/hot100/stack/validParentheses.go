package stack

func isValid(s string) bool {

	stack := make([]rune, len(s))
	pointer := -1
	for _, val := range s {

		if pointer == -1 {
			pointer++
			stack[pointer] = val
		} else {
			switch true {
			case stack[pointer] == '[' && val == ']':
				pointer--
			case stack[pointer] == '{' && val == '}':
				pointer--
			case stack[pointer] == '(' && val == ')':
				pointer--
			default:
				pointer++
				stack[pointer] = val

			}
		}
	}
	return pointer == -1
}
