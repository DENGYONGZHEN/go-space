package hashtable

import "testing"

func TestIsAnagram(t *testing.T) {
	tt := []struct {
		s    string
		t    string
		want bool
	}{
		{s: "anagram", t: "nagaram", want: true},
		{s: "rat", t: "car", want: false},
	}

	for _, tc := range tt {
		got := isAnagram(tc.s, tc.t)
		if got != tc.want {
			t.Errorf("isAnagram(%v,%v)=%v; want %v instead \n", tc.s, tc.t, got, tc.want)
		}
	}
}
