package recursive

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test 0",
			n:   0,
			want: 0,
		},
		{
			name: "test 1",
			n:   1,
			want: 1,
		},
		{
			name: "test 2",
			n:   2,
			want: 1,
		},
		{
			name: "test 3",
			n:   3,
			want: 2,
		},
		{
			name: "test 4",
			n:   4,
			want: 3,
		},
		{
			name: "test 5",
			n:   5,
			want: 5,
		},
		{
			name: "test 6",
			n:   8,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
