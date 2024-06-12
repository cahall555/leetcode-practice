package binarytree

import (
	"testing"
)

func TestDepthFirstRec(t *testing.T) {
	root := &BTNode{Value: "a"}
	b := &BTNode{Value: "b"}
	c := &BTNode{Value: "c"}
	d := &BTNode{Value: "d"}
	e := &BTNode{Value: "e"}
	f := &BTNode{Value: "f"}
	g := &BTNode{Value: "g"}

	root.Left = b
	root.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g

//     a
//    / \
//   b   c
//  / \   \
// d   e   f
//    /
//   g

	expected := []string{"a", "b", "d", "e", "g", "c", "f"}
	result := DepthFirstRec(root)

	if len(expected) != len(result) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	}

}

func TestDepthFirstIt(t *testing.T) {
	root := &BTNode{Value: "a"}
	b := &BTNode{Value: "b"}
	c := &BTNode{Value: "c"}
	d := &BTNode{Value: "d"}
	e := &BTNode{Value: "e"}
	f := &BTNode{Value: "f"}
	g := &BTNode{Value: "g"}

	root.Left = b
	root.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g

//     a
//    / \
//   b   c
//  / \   \
// d   e   f
//    /
//   g

result := DepthFirstIt(root)
expected := []string{"a", "b", "d", "e", "g", "c", "f"}

if len(expected) != len(result) {
	t.Errorf("Expected %v but got %v", expected, result)
}

for i := 0; i < len(expected); i++ {
	if expected[i] != result[i] {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
}
