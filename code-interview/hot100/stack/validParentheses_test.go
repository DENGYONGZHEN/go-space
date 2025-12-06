package stack

import "testing"

func TestIsValid(t *testing.T) {

	testCases := []struct {
		s    string
		want bool
	}{
		{
			s:    "()",
			want: true,
		},
		{
			s:    "()[]{}",
			want: true,
		},
		{
			s:    "(]",
			want: false,
		},
		{
			s:    "([])",
			want: true,
		},
		{
			s:    "({[)",
			want: false,
		},
	}

	for _, tc := range testCases {
		res := isValid(tc.s)
		if res != tc.want {
			t.Errorf("isValid(%v),should get %v,but %v", tc.s, tc.want, res)
		}
	}
}
