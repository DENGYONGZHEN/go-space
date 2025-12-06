package backtracking

import (
	"fmt"
	"testing"
)

// 辅助函数：检查两个二维切片是否包含相同的元素（顺序无关）
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 将每个排列转换为字符串并统计次数
	mapA := make(map[string]int)
	for _, arr := range a {
		mapA[fmt.Sprint(arr)]++
	}

	// 检查b中的每个排列是否在mapA中存在且次数匹配
	for _, arr := range b {
		key := fmt.Sprint(arr)
		if mapA[key] == 0 {
			return false
		}
		mapA[key]--
	}

	return true
}

// 测试空输入
func TestPermuteEmpty(t *testing.T) {
	nums := []int{}
	expected := [][]int{{}}
	res := permute(nums)
	if !equal(res, expected) {
		t.Errorf("TestPermuteEmpty failed. Expected %v, got %v", expected, res)
	}
}

// 测试单元素
func TestPermuteSingle(t *testing.T) {
	nums := []int{1}
	expected := [][]int{{1}}
	res := permute(nums)
	if !equal(res, expected) {
		t.Errorf("TestPermuteSingle failed. Expected %v, got %v", expected, res)
	}
}

// 测试两个元素
func TestPermuteTwo(t *testing.T) {
	nums := []int{1, 2}
	expected := [][]int{
		{1, 2},
		{2, 1},
	}
	res := permute(nums)
	if !equal(res, expected) {
		t.Errorf("TestPermuteTwo failed. Expected %v, got %v", expected, res)
	}
}

// 测试三个元素
func TestPermuteThree(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	res := permute(nums)
	if !equal(res, expected) {
		t.Errorf("TestPermuteThree failed. Expected %v, got %v", expected, res)
	}
}
