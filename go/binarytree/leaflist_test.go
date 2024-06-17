package binarytree

import (
	"testing"
)

func TestLeafListRec(t *testing.T) {
	root1 := &BTNode{Value: "a"}
	root1.Left = &BTNode{Value: "b"}
	root1.Right = &BTNode{Value: "c"}
	root1.Left.Left = &BTNode{Value: "d"}
	root1.Left.Right = &BTNode{Value: "e"}
	root1.Right.Left = &BTNode{Value: "f"}

	result1 := LeafListRec(root1)
	expected1 := []string{"d", "e", "f"}

	if len(result1) != len(expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	for i := range result1 {
		if result1[i] != expected1[i] {
			t.Errorf("Expected %v, but got %v", expected1, result1)
		}
	}
}

func TestLeafListIt(t *testing.T) {

	root1 := &BTNode{Value: "a"}
	root1.Left = &BTNode{Value: "b"}
	root1.Right = &BTNode{Value: "c"}
	root1.Left.Left = &BTNode{Value: "d"}
	root1.Left.Right = &BTNode{Value: "e"}
	root1.Right.Left = &BTNode{Value: "f"}

	result1 := LeafListIt(root1)
	expected1 := []string{"d", "e", "f"}

	if len(result1) != len(expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	for i := range result1 {
		if result1[i] != expected1[i] {
			t.Errorf("Expected %v, but got %v", expected1, result1)
		}
	}
}
