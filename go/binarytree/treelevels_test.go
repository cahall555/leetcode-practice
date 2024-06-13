package binarytree

import ( 
	"testing"
)

func TestTreeLevelsRec(t *testing.T) {
	//Test 1
	root1 := &BTNode{Value: "a"}
	root1.Left = &BTNode{Value: "b"}
	root1.Right = &BTNode{Value: "c"}
	root1.Left.Left = &BTNode{Value: "d"}
	root1.Left.Right = &BTNode{Value: "e"}
	root1.Right.Left = &BTNode{Value: "f"}


	expected1 := [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}}
	result1 := TreeLevelsRec(root1)

	if (len(result1) != len(expected1)) {
		t.Errorf("Expected %v but got %v", expected1, result1)
	}

	for i := 0; i < len(expected1); i++ {
		if (len(result1[i]) != len(expected1[i])) {
			t.Errorf("Expected %v but got %v", expected1, result1)
		}
		for j := 0; j < len(expected1[i]); j++ {
			if (expected1[i][j] != result1[i][j]) {
				t.Errorf("Expected %v but got %v", expected1, result1)
			}
		}
	}

	//Test 2
	root2 := &BTNode{}

	expected2 := [][]string{}
	result2 := TreeLevelsRec(root2)

	if (len(result2) != len(expected2)) {
		t.Errorf("Expected %v but got %v", expected2, result2)
	}

	//Test 3
	root3 := &BTNode{Value: "a"}

	expected3 := [][]string{{"a"}}
	result3 := TreeLevelsRec(root3)

	if (len(result3) != len(expected3)) {
		t.Errorf("Expected %v but got %v", expected3, result3)
	}

}
