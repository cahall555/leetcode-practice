package graph

import (
	"testing"
)

func TestBestBridge(t *testing.T) {
	grid := [][]string{
		{"W", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "L"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
	}
	expected := 2
	if result := BestBridge(grid); result != expected {
		t.Errorf("BestBridge(%v) = %v, want %v", grid, result, expected)
	}

	gridTwo := [][]string{
		{"W", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "L"},
		{"L", "L", "L", "W", "L"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "W", "W", "W"},
	}
	expectedTwo := 1
	if result := BestBridge(gridTwo); result != expectedTwo {
		t.Errorf("BestBridge(%v) = %v, want %v", gridTwo, result, expectedTwo)
	}
}
