package recursive

import (
	"testing"
)

func TestReverseString (t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 0",
			s:   "hello",
			want: "olleh",
		},
		{
			name: "test 1",
			s:   "abcdefg",
			want: "gfedcba",
		},
		{
			name: "test 2",
			s:   "stopwatch",
			want: "hctawpots",
		},
		{
			name: "test 3",
			s:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
