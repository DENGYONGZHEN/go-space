package backtracking

import (
	"testing"
)

// 辅助函数：检查两个字符串切片是否相等
func equals(a, b []string) bool {
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

// 测试空输入
func TestLetterCombinationsEmpty(t *testing.T) {
	digits := ""
	expected := []string{}
	res := letterCombinations(digits)
	if !equals(res, expected) {
		t.Errorf("TestLetterCombinationsEmpty failed. Expected %v, got %v", expected, res)
	}
}

// 测试单数字
func TestLetterCombinationsSingle(t *testing.T) {
	digits := "2"
	expected := []string{"a", "b", "c"}
	res := letterCombinations(digits)
	if !equals(res, expected) {
		t.Errorf("TestLetterCombinationsSingle failed. Expected %v, got %v", expected, res)
	}
}

// 测试两个数字
func TestLetterCombinationsTwo(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	res := letterCombinations(digits)
	if !equals(res, expected) {
		t.Errorf("TestLetterCombinationsTwo failed. Expected %v, got %v", expected, res)
	}
}

// 测试多个数字
func TestLetterCombinationsMultiple(t *testing.T) {
	digits := "79"
	expected := []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"}
	res := letterCombinations(digits)
	if !equals(res, expected) {
		t.Errorf("TestLetterCombinationsMultiple failed. Expected %v, got %v", expected, res)
	}
}
