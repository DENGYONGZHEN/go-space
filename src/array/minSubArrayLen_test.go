package array

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tt := []struct {
		want   int
		nums   []int
		target int
	}{{want: 2,
		nums:   []int{5, 1, 4, 3},
		target: 7},
		{want: 1, nums: []int{3, 4, 2, 6}, target: 1},
	}

	for _, val := range tt {

		got := MinSubArrayLen(val.target, val.nums)

		if got != val.want {
			t.Errorf("MinSubArrayLen(%v,%v)=%v; want %v instead \n", val.target, val.nums, got, val.want)
		}

	}
}
