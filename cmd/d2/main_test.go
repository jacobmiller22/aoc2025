package main

import "testing"

func TestRepeatingSequence(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "simple non repeating",
			s:    "1000",
			want: false,
		},
		{
			name: "simple all repeating",
			s:    "111",
			want: true,
		},
		{
			name: "bigger repeating",
			s:    "1010",
			want: true,
		},
		{
			name: "panic",
			s:    "2121212118",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := repeatingSequence(tc.s)
			if tc.want != got {
				t.Errorf("Want: %v; Got: %v", tc.want, got)
			}
		})
	}
}
