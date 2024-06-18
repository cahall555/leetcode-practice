package graph

import (
	"testing"
)

func TestHasPathRec(t *testing.T) {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {"f"},
		"f": {},
	}

	if !HasPathRec(&graph, "a", "f") {
		t.Error("Expected true")
	}

	if HasPathRec(&graph, "a", "g") {
		t.Error("Expected false")
	}
}


func TestHasPathIt(t *testing.T) {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {"f"},
		"f": {},
	}

	if !HasPathIt(&graph, "a", "f") {
		t.Error("Expected true")
	}

	if HasPathIt(&graph, "a", "g") {
		t.Error("Expected false")
	}
}
