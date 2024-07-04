package dynamic

import (
	"testing"
)

func TestCountPaths(t *testing.T) {
	grid := [][]string{
		{"O", "O", "O"},
		{"O", "O", "X"},
		{"O", "O", "O"},
	}

	memo := make(map[string]int)
	result := CountPaths(grid, 0, 0, memo)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}
