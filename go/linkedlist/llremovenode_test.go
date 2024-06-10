package linkedlist

import (
	"testing"
)

func TestRemoveNodeRec(t *testing.T) {
	//Test 1
	test1 := &Node{Value: "a"}
	test1.Next = &Node{Value: "b"}
	test1.Next.Next = &Node{Value: "c"}
	test1.Next.Next.Next = &Node{Value: "d"}
	test1.Next.Next.Next.Next = &Node{Value: "e"}

	result1 := RemoveNodeRec(test1, "c")
	expected1 := "a b d e "
	if PrintLinkedList(result1) != expected1 {
		t.Errorf("Expected %s but got %s", expected1, PrintLinkedList(result1))
	}
	//Test 2
	test2 := &Node{Value: "x"}
	test2.Next = &Node{Value: "y"}
	test2.Next.Next = &Node{Value: "z"}

	result2 := RemoveNodeRec(test2, "z")
	expected2 := "x y "
	if PrintLinkedList(result2) != expected2 {
		t.Errorf("Expected %s but got %s", expected2, PrintLinkedList(result2))
	}

	//Test 3
	test3 := &Node{Value: "q"}
	test3.Next = &Node{Value: "r"}
	test3.Next.Next = &Node{Value: "s"}

	result3 := RemoveNodeRec(test3, "q")
	expected3 := "r s "
	if PrintLinkedList(result3) != expected3 {
		t.Errorf("Expected %s but got %s", expected3, PrintLinkedList(result3))
	}
	
	//Test 4
	test4 := &Node{Value: "h"}
	test4.Next = &Node{Value: "i"}
	test4.Next.Next = &Node{Value: "j"}
	test4.Next.Next.Next = &Node{Value: "i"}

	result4 := RemoveNodeRec(test4, "i")
	expected4 := "h j i "
	if PrintLinkedList(result4) != expected4 {
		t.Errorf("Expected %s but got %s", expected4, PrintLinkedList(result4))
	}

	//Test 5
	test5 := &Node{Value: "t"}

	result5 := RemoveNodeRec(test5, "t")
	expected5 := ""
	if PrintLinkedList(result5) != expected5 {
		t.Errorf("Expected %s but got %s", expected5, PrintLinkedList(result5))
	}
}

func TestRemoveNodeIt(t *testing.T) {
	//Test 1
	test1 := &Node{Value: "a"}
	test1.Next = &Node{Value: "b"}	
	test1.Next.Next = &Node{Value: "c"}
	test1.Next.Next.Next = &Node{Value: "d"}
	test1.Next.Next.Next.Next = &Node{Value: "e"}

	result1 := RemoveNodeIt(test1, "c")
	expected1 := "a b d e "
	if PrintLinkedList(result1) != expected1 {
		t.Errorf("Expected %s but got %s", expected1, PrintLinkedList(result1))
	}
	//Test 2
	test2 := &Node{Value: "x"}
	test2.Next = &Node{Value: "y"}
	test2.Next.Next = &Node{Value: "z"}

	result2 := RemoveNodeIt(test2, "z")
	expected2 := "x y "
	if PrintLinkedList(result2) != expected2 {
		t.Errorf("Expected %s but got %s", expected2, PrintLinkedList(result2))
	}

	//Test 3
	test3 := &Node{Value: "q"}
	test3.Next = &Node{Value: "r"}
	test3.Next.Next = &Node{Value: "s"}

	result3 := RemoveNodeIt(test3, "q")
	expected3 := "r s "
	if PrintLinkedList(result3) != expected3 {
		t.Errorf("Expected %s but got %s", expected3, PrintLinkedList(result3))
	}

	//Test 4
	test4 := &Node{Value: "h"}
	test4.Next = &Node{Value: "i"}
	test4.Next.Next = &Node{Value: "j"}
	test4.Next.Next.Next = &Node{Value: "i"}

	result4 := RemoveNodeIt(test4, "i")
	expected4 := "h j i "
	if PrintLinkedList(result4) != expected4 {
		t.Errorf("Expected %s but got %s", expected4, PrintLinkedList(result4))
	}

	//Test 5
	test5 := &Node{Value: "t"}

	result5 := RemoveNodeIt(test5, "t")
	expected5 := ""
	if PrintLinkedList(result5) != expected5 {
		t.Errorf("Expected %s but got %s", expected5, PrintLinkedList(result5))
	}
}
