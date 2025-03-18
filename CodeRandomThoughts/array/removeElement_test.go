package array

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
	}

	for _, tt := range tests {
		got := removeElement(tt.nums, tt.target)
		if !reflect.DeepEqual(tt.nums[:got], tt.want) {
			t.Errorf("removeElement(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}

}
