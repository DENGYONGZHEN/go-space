package binarysearch

import (
	"reflect"
	"testing"
)

func TestFindFirstAndLastPosition(t *testing.T) {

	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		val := searchRange(tc.nums, tc.target)
		if !reflect.DeepEqual(val, tc.want) {
			t.Errorf("searchRange(%v, %v), want %v but %v", tc.nums, tc.target, tc.want, val)
		}
	}

}
