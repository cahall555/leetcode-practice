package recursive

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test 0",
			n:    3,
			want: 6,
		},
		{
			name: "test 1",
			n:    6,
			want: 720,
		},
		{
			name: "test 2",
			n:     18,
			want: 6402373705728000,
		},
		{
			name: "test 3",
			n:     1,
			want: 1,
		},
		{
			name: "test 4",
			n:     13,
			want: 6227020800,
		},
		{
			name: "test 5",
			n:     0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
