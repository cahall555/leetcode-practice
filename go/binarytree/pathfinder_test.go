package binarytree


import (
	"testing"
)

func TestPathFinder(t *testing.T) {
	root := &BTNode{Value: "a"}
	root.Left = &BTNode{Value: "b"}
	root.Right = &BTNode{Value: "c"}
	root.Left.Left = &BTNode{Value: "d"}
	root.Left.Right = &BTNode{Value: "e"}

	result := PathFinder(root, "d")
	expected := []string{"a", "b", "d"}
	if len(result) != len(expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	}

}
