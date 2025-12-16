package stackandqueue

import (
	"log"
	"strconv"
)

// 150. Evaluate Reverse Polish Notation
// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
// Evaluate the expression. Return an integer that represents the value of the expression.

// Note that:
// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// The division between two integers always truncates toward zero.
// There will not be any division by zero.
// The input represents a valid arithmetic expression in a reverse polish notation.
// The answer and all the intermediate calculations can be represented in a 32-bit integer.

// 逆波兰表达式：是类似于二叉树的后序遍历。详细可搜索。另外关于二叉树遍历，二叉树章节有详细

// 思路：遇到数字就入栈，遇到操作符，就从栈中取两个数进行相关操作，操作结果再次入栈
func evalRPN(tokens []string) int {

	stack := make([]int, len(tokens))
	pointer := -1
	for _, s := range tokens {

		if pointer == -1 {
			pointer++
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			stack[pointer] = v
		} else {

			switch true {
			case s == "+":
				stack[pointer-1] = stack[pointer-1] + stack[pointer]
				pointer--
			case s == "-":
				stack[pointer-1] = stack[pointer-1] - stack[pointer]
				pointer--
			case s == "*":
				stack[pointer-1] = stack[pointer-1] * stack[pointer]
				pointer--
			case s == "/":
				stack[pointer-1] = stack[pointer-1] / stack[pointer]
				pointer--
			default:
				pointer++
				v, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				stack[pointer] = v
			}
		}
	}

	return stack[0]
}
