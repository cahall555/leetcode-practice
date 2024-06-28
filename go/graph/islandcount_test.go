package graph

import (
	"testing"
)

func TestIslandCount(t *testing.T){
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
	}

	expected := 3
	result := IslandCount(grid)

	if expected != result {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
