package recursive

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name string
		numbers []int
		want int
	}{
		{
			name: "test 0",
			numbers: []int{5, 2, 9, 10},
			want: 26,
		},
		{
			name: "test 1",
			numbers: []int{1, -1, 1, -1, 1, -1, 1},
			want: 1,
		},
		{
			name: "test 2",
			numbers: []int{},
			want: 0,
		},
		{
			name: "test 3",
			numbers: []int{1000, 0, 0, 0, 0, 0, 1},
			want: 1001,
		},
		{
			name: "test 4",
			numbers: []int{700, 70, 7},
			want: 777,
		},
		{
			name: "test 5",
			numbers: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
			want: -55,
		},
		{
			name: "test 6",
			numbers: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "test 7",
			numbers: []int{123456789, 12345678, 1234567, 123456, 12345, 1234, 123, 12, 1, 0},
			want: 137174205,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumbers(tt.numbers); got != tt.want {
				t.Errorf("SumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
