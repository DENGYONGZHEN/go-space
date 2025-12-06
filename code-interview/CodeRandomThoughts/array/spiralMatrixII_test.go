package array

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		num  int
		want [][]int
	}{
		{
			num: 3,
			want: [][]int{
				{1, 2, 3}, {8, 9, 4}, {7, 6, 5},
			},
		},
		{
			num: 1,
			want: [][]int{
				{1},
			},
		},
		{
			num: 4,
			want: [][]int{
				{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7},
			},
		},
	}

	for _, tc := range testCases {
		val := generateMatrix(tc.num)
		if !reflect.DeepEqual(val, tc.want) {
			t.Errorf("generateMatrix(%v),want %v but get %v", tc.num, tc.want, val)
		}
	}
}
