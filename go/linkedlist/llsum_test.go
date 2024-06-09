package linkedlist

import (
	"testing"
)

func TestLinkedListSum(t *testing.T) {
	//Test 1
	one := IntNode{Value: 2}
	two := IntNode{Value: 8}
	three := IntNode{Value: 3}
	four := IntNode{Value: -1}
	five := IntNode{Value: 7}

	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five

	result := LinkedListSum(&one)
	expected := 19
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	//Test 2
	x := IntNode{Value: 38}
	y := IntNode{Value: 4}

	x.Next = &y

	result = LinkedListSum(&x)
	expected = 42
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	
	//Test 3
	z := IntNode{Value: 100}

	result = LinkedListSum(&z)
	expected = 100
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	
	//Test 4
	result = LinkedListSum(nil)
	expected = 0
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}


