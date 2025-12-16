package stackandqueue

import (
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

func TestMaxSlidingWindow(t *testing.T) {

	testCase := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			input: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:     3,
			want:  []int{3, 3, 5, 5, 6, 7},
		},
		{
			input: []int{1},
			k:     1,
			want:  []int{1},
		},
	}

	for _, tc := range testCase {
		got := maxSlidingWindow(tc.input, tc.k)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("TestMaxSlidingWindow want %v, instead %v", tc.want, got)
		}
	}
}
