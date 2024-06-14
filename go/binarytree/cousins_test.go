package binarytree

import (
	"testing"
)

func TestCousins(t *testing.T) {
	root := &BTIntNode{Value: 1}
	root.Left = &BTIntNode{Value: 2}
	root.Right = &BTIntNode{Value: 3}
	root.Left.Left = &BTIntNode{Value: 4}
	root.Left.Right = &BTIntNode{Value: 5}
	root.Right.Left = &BTIntNode{Value: 6}
	root.Right.Right = &BTIntNode{Value: 7}

	if Cousins(root, 4, 5) {
		t.Errorf("Expected false, got true")
	}

	if !Cousins(root, 4, 6) {
		t.Errorf("Expected true, got false")
	}

	if Cousins(root, 3, 4) {
		t.Errorf("Expected false, got true")
	}

	if Cousins(root, 2, 3) {
		t.Errorf("Expected false, got true")
	}
}
