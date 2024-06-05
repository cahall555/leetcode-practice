package dynamic

import (
	"testing"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test 0",
			n:    0,
			want: 0,
		},
		{
			name: "test 1",
			n:    1,
			want: 0,
		},
		{
			name: "test 2",
			n:    2,
			want: 1,
		},
		{	
			name: "test 3",
			n:    5,
			want: 4,
		},
		{
			name: "test 4",
			n:    7,
			want: 13,
		},
			{
			name: "test 5",
			n:    14,
			want: 927,
		},
		{
			name: "test 6",
			n:    20,
			want: 35890,
		},
		{
			name: "test 7",
			n:    37,
			want: 1132436852,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tribonacci(tt.n); got != tt.want {
				t.Errorf("Tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
