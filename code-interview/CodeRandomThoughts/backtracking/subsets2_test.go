package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsets2(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			nums: []int{4, 4, 4, 1, 4},
			want: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		},
	}
	for _, tc := range testCases {
		got := subsetsWithDup(tc.nums)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("TestSubsets want %v, got %v", tc.want, got)
		}
	}
}
