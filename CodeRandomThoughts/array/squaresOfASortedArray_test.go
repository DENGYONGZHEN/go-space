package array

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {

	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
	}

	for _, tc := range testCases {
		result := sortedSquares(tc.nums)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("TestSortedSquares(%v); want %v instead %v \n", tc.nums, tc.want, result)
		}
	}

}
