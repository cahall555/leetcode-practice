package recursive

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "test 0",
			s:    "pop",
			want: true,
		},
		{
			name: "test 1",
			s:    "kayak",
			want: true,
		},
		{
			name: "test 2",
			s:    "pops",
			want: false,
		},
		{
			name: "test 3",
			s:    "boot",
			want: false,
		},
		{
			name: "test 4",
			s:    "rotator",
			want: true,
		},
		{
			name: "test 5",
			s:    "abcbca",
			want: false,
		},
		{
			name: "test 6",
			s:    "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.s); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
