package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {

	testCases := []struct {
		matrix [][]int
		target int
		exist  bool
	}{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			exist:  true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			exist:  false,
		},
		{
			matrix: [][]int{{1, 1}},
			target: 2,
			exist:  false,
		},
		{
			matrix: [][]int{{1, 3}},
			target: 3,
			exist:  true,
		},
	}

	for _, tc := range testCases {
		val := searchMatrix(tc.matrix, tc.target)
		if val != tc.exist {
			t.Errorf("search(%v,%v),exist %v,but get %v", tc.matrix, tc.target, tc.exist, val)
		}
	}
}
