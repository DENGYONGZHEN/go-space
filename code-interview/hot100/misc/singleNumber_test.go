package misc

import "testing"

func TestSingleNumber(t *testing.T) {

	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			nums: []int{1},
			want: 1,
		},
	}

	for _, tc := range testCases {
		val := singleNumber(tc.nums)
		if val != tc.want {
			t.Errorf("singleNumber(%v),want %v,but get %v", tc.nums, tc.want, val)
		}
	}
}
