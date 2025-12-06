package array

import (
	"reflect"
	"testing"
)

func TestThreeSumZero(t *testing.T) {
	test := []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{-2, 0, 1, 1, 2}, want: [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for _, tt := range test {

		got := threeSum(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ThreeSum(%v)=%v; want %v", tt.nums, got, tt.want)
		}

	}
}
