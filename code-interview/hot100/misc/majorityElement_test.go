package misc

import "testing"

func TestMajorityElement(t *testing.T) {

	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, tc := range testCases {
		val := majorityElement(tc.nums)
		if val != tc.want {
			t.Errorf("majorityElement(%v),should get %v,but get %v", tc.nums, tc.want, val)
		}
	}
}
