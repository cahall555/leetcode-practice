package recursive

import (
	"testing"
)

func TestSumOfLengths(t *testing.T) {
	tests := []struct {
		name    string
		strings []string
		want    int
	}{
		{
			name:    "test 0",
			strings: []string{"goat", "cat", "purple"},
			want:    13,
		},
		{
			name:    "test 1",
			strings: []string{"bike", "at", "pencils", "phone"},
			want:    18,
		},
		{
			name:    "test 2",
			strings: []string{},
			want:    0,
		},
		{
			name:    "test 3",
			strings: []string{"", " ", "  ", "   ", "    ", "     "},
			want:    15,
		},
		{
			name:    "test 4",
			strings: []string{"0", "313", "1234567890"},
			want:    14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfLengths(tt.strings); got != tt.want {
				t.Errorf("SumOfLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
