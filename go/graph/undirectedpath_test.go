package graph

import (
	"testing"
)

func TestUndirectedPath(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	if UndirectedPath(edges, "i", "l") != true {
		t.Error("Expected true")
	}

	if UndirectedPath(edges, "i", "m") != true {
		t.Error("Expected true")
	}

	if UndirectedPath(edges, "i", "o") != false {
		t.Error("Expected false")
	}
}
