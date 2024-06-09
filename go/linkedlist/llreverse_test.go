package linkedlist

import (
	"testing"
)

func TestLinkedListReverse(t *testing.T) {

	a := Node{Value: "a"}
	b := Node{Value: "b"}
	c := Node{Value: "c"}
	d := Node{Value: "d"}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	reversed := LinkedListReverse(&a)
	if reversed.Value != "d" {
		t.Errorf("Expected d, got %s", reversed.Value)
	}

	head := &Node{Value: "1", Next: &Node{Value: "2", Next: &Node{Value: "3", Next: nil}}}
	reversed = LinkedListReverse(head)
	if reversed.Value != "3" {
		t.Errorf("Expected 3, got %s", reversed.Value)
	}
	if reversed.Next.Value != "2" {
		t.Errorf("Expected 2, got %s", reversed.Next.Value)
	}
	if reversed.Next.Next.Value != "1" {
		t.Errorf("Expected 1, got %s", reversed.Next.Next.Value)
	}
}

func TestLinkedListItReverse(t *testing.T) {
	a := Node{Value: "a"}
	b := Node{Value: "b"}
	c := Node{Value: "c"}
	d := Node{Value: "d"}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	reversed := LinkedListItReverse(&a)
	if reversed.Value != "d" {
		t.Errorf("Expected d, got %s", reversed.Value)
	}

	head := &Node{Value: "1", Next: &Node{Value: "2", Next: &Node{Value: "3", Next: nil}}}
	reversed = LinkedListItReverse(head)
	if reversed.Value != "3" {
		t.Errorf("Expected 3, got %s", reversed.Value)
	}
	if reversed.Next.Value != "2" {
		t.Errorf("Expected 2, got %s", reversed.Next.Value)
	}
	if reversed.Next.Next.Value != "1" {
		t.Errorf("Expected 1, got %s", reversed.Next.Next.Value)
	}
}
