package graph

import (
	"testing"
)

func TestClosestCarrot(t *testing.T) {
	closestCarrotOne := [][]string{
		{"O", "O", "X", "X", "X"},
		{"O", "X", "X", "X", "C"},
		{"O", "X", "O", "X", "X"},
		{"O", "O", "O", "O", "O"},
		{"O", "X", "X", "X", "X"},
		{"O", "O", "O", "O", "O"},
		{"O", "O", "C", "O", "O"},
		{"O", "O", "O", "O", "O"},
	}

	actualClosestCarrotOne := ClosestCarrot(closestCarrotOne, 3, 4)
	expectedClosestCarrotOne := 9

	if actualClosestCarrotOne != expectedClosestCarrotOne {
		t.Errorf("ClosestCarrot was incorrect, got: %d, want: %d.", actualClosestCarrotOne, expectedClosestCarrotOne)
	}
}
