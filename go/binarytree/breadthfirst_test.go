package binarytree

import ( 
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	//Test 1
	root1 := &BTNode{Value: "a"}
	root1.Left = &BTNode{Value: "b"}
	root1.Right = &BTNode{Value: "c"}
	root1.Left.Left = &BTNode{Value: "d"}
	root1.Left.Right = &BTNode{Value: "e"}
	root1.Right.Left = &BTNode{Value: "f"}

	expected1 := []string{"a", "b", "c", "d", "e", "f"}
	result1 := BreadthFirst(root1)
	if (len(result1) != len(expected1)) {
		t.Errorf("Expected %v but got %v", expected1, result1)
	}

	for i := 0; i < len(expected1); i++ {
		if (expected1[i] != result1[i]) {
			t.Errorf("Expected %v but got %v", expected1, result1)
		}
	}

	//Test 2
	root2 := &BTNode{Value: "a"}
	root2.Left = &BTNode{Value: "b"}
	root2.Right = &BTNode{Value: "c"}
	root2.Left.Left = &BTNode{Value: "d"}
	root2.Left.Right = &BTNode{Value: "e"}
	root2.Right.Left = &BTNode{Value: "f"}
	root2.Right.Right = &BTNode{Value: "g"}
	root2.Left.Left.Left = &BTNode{Value: "h"}

	expected2 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	result2 := BreadthFirst(root2)
	if (len(result2) != len(expected2)) {
		t.Errorf("Expected %v but got %v", expected2, result2)
	}

	for i := 0; i < len(expected2); i++ {
		if (expected2[i] != result2[i]) {
			t.Errorf("Expected %v but got %v", expected2, result2)
		}
	}
}
