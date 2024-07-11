package dynamic

import (
	"testing"
)

func TestArrayStepper(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 4, 2, 0, 0, 1}, true},
		{[]int{1, 1, 0, 1}, false},
		{[]int{2, 4, 2, 0, 0, 1}, true},
		{[]int{1, 1, 1, 0, 1}, false},
	}
	for _, test := range tests {
		actual := ArrayStepper(test.nums, 0, make(map[int]int))
		if actual != test.expected {
			t.Errorf("ArrayStepper(%v): expected %v, actual %v", test.nums, test.expected, actual)
		}
	}
}
