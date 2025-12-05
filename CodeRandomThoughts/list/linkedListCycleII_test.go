package list

import (
	"slices"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	testCases := []struct {
		nums []int
		pos  int
		want int
	}{
		{nums: []int{3, 2, 0, -4},
			pos:  1,
			want: 1,
		},
		{
			nums: []int{1, 2},
			pos:  0,
			want: 0,
		},
		{
			nums: []int{1},
			pos:  -1,
			want: -1,
		},
	}

	for _, tc := range testCases {
		head := makeLinkedListWithCycle(tc.nums, tc.pos)
		result := detectCycle(head)
		var resultValue int
		if result == nil {
			resultValue = -1
		} else {
			resultValue = slices.Index(tc.nums, result.Val)
		}
		if resultValue != tc.want {
			t.Fatalf("want %v,get %v", tc.want, result)
		}
	}
}
