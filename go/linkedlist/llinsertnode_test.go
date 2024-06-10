package linkedlist

import (
	"testing"
)

func TestInsertNodeRec(t *testing.T) {
	//Test 1
	test1 := &Node{Value: "a"}
	test1.Next = &Node{Value: "b"}
	test1.Next.Next = &Node{Value: "c"}
	test1.Next.Next.Next = &Node{Value: "d"}

	result1 := InsertNodeRec(test1, "x", 2, 0)
	expected1 := "a b x c d "

	if PrintLinkedList(result1) != expected1 {
		t.Errorf("Expected %s but got %s", expected1, PrintLinkedList(result1))
	}

	//Test 2
	test2 := &Node{Value: "a"}
	test2.Next = &Node{Value: "b"}
	test2.Next.Next = &Node{Value: "c"}
	test2.Next.Next.Next = &Node{Value: "d"}

	result2 := InsertNodeRec(test2, "y", 3, 0)
	expected2 := "a b c y d "

	if PrintLinkedList(result2) != expected2 {
		t.Errorf("Expected %s but got %s", expected2, PrintLinkedList(result2))
	}	

	//Test 3
	test3 := &Node{Value: "a"}
	test3.Next = &Node{Value: "b"}
	test3.Next.Next = &Node{Value: "c"}
	test3.Next.Next.Next = &Node{Value: "d"}

	result3 := InsertNodeRec(test3, "z", 4, 0)
	expected3 := "a b c d z "

	if PrintLinkedList(result3) != expected3 {
		t.Errorf("Expected %s but got %s", expected3, PrintLinkedList(result3))
	}

}

func TestInsertNodeIt(t *testing.T) {
	//Test 1
	test1 := &Node{Value: "a"}
	test1.Next = &Node{Value: "b"}
	test1.Next.Next = &Node{Value: "c"}
	test1.Next.Next.Next = &Node{Value: "d"}

	result1 := InsertNodeIt(test1, "x", 2)
	expected1 := "a b x c d "

	if PrintLinkedList(result1) != expected1 {
		t.Errorf("Expected %s but got %s", expected1, PrintLinkedList(result1))
	}

	//Test 2
	test2 := &Node{Value: "a"}
	test2.Next = &Node{Value: "b"}
	test2.Next.Next = &Node{Value: "c"}
	test2.Next.Next.Next = &Node{Value: "d"}

	result2 := InsertNodeIt(test2, "y", 3)
	expected2 := "a b c y d "

	if PrintLinkedList(result2) != expected2 {
		t.Errorf("Expected %s but got %s", expected2, PrintLinkedList(result2))
	}

	//Test 3
	test3 := &Node{Value: "a"}
	test3.Next = &Node{Value: "b"}
	test3.Next.Next = &Node{Value: "c"}
	test3.Next.Next.Next = &Node{Value: "d"}

	result3 := InsertNodeIt(test3, "z", 4)
	expected3 := "a b c d z "
		
	if PrintLinkedList(result3) != expected3 {
		t.Errorf("Expected %s but got %s", expected3, PrintLinkedList(result3))
	}
}

