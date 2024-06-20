package graph

import (
	"testing"
)

func TestConnectedComponentCount(t *testing.T) {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0},
		8: {0},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	if ConnectedComponentCount(&graph) != 2 {
		t.Error("Expected 2")
	}
}

func BenchmarkConnectedComponentCount(b *testing.B) {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0},
		8: {0},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	for i := 0; i < b.N; i++ {
		ConnectedComponentCount(&graph)
	}
}

func TestConnectedComponentCountSet(t *testing.T) {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0},
		8: {0},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	if ConnectedComponentCountSet(&graph) != 2 {
		t.Error("Expected 2")
	}
}

func BenchmarkConnectedComponentCountSet(b *testing.B) {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0},
		8: {0},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	for i := 0; i < b.N; i++ {
		ConnectedComponentCountSet(&graph)
	}
}
