package twopointer

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
	}

	for _, tc := range testCases {
		moveZeroes(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.want) {
			t.Errorf("moveZeroes(%v),want %v,but get %v", tc.nums, tc.want, tc.nums)
		}
	}
}
