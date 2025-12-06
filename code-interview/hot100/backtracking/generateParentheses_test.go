package backtracking

import (
	"testing"
)

// 辅助函数：检查两个字符串切片是否相等
func equalstr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 测试 n = 0
func TestGenerateParenthesesZero(t *testing.T) {
	n := 0
	expected := []string{""}
	res := generateParenthesis(n)
	if !equalstr(res, expected) {
		t.Errorf("TestGenerateParenthesesZero failed. Expected %v, got %v", expected, res)
	}
}

// 测试 n = 1
func TestGenerateParenthesesOne(t *testing.T) {
	n := 1
	expected := []string{"()"}
	res := generateParenthesis(n)
	if !equalstr(res, expected) {
		t.Errorf("TestGenerateParenthesesOne failed. Expected %v, got %v", expected, res)
	}
}

// 测试 n = 2
func TestGenerateParenthesesTwo(t *testing.T) {
	n := 2
	expected := []string{"(())", "()()"}
	res := generateParenthesis(n)
	if !equalstr(res, expected) {
		t.Errorf("TestGenerateParenthesesTwo failed. Expected %v, got %v", expected, res)
	}
}

// 测试 n = 3
func TestGenerateParenthesesThree(t *testing.T) {
	n := 3
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	res := generateParenthesis(n)
	if !equalstr(res, expected) {
		t.Errorf("TestGenerateParenthesesThree failed. Expected %v, got %v", expected, res)
	}
}
