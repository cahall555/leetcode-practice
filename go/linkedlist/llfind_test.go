package linkedlist

import (
	"testing"
)

func TestLinkedListFind(t *testing.T) {
	//Test 1
	a := Node{Value: "a"}
	b := Node{Value: "b"}
	c := Node{Value: "c"}
	d := Node{Value: "d"}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	result := LinkedListFind(&a, "c")
	expected := true
        if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	//Test 2
	result = LinkedListFind(&a, "d")
	expected = true
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	//Test 3
	result = LinkedListFind(&a, "z")
	expected = false
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
