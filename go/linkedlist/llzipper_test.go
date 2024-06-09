package linkedlist

import (
	"testing"
)

func TestLinkedListRecZipper(t *testing.T) {
	head1 := &Node{Value: "1"}
	head1.Next = &Node{Value: "2"}
	head1.Next.Next = &Node{Value: "3"}
	head1.Next.Next.Next = &Node{Value: "4"}

	head2 := &Node{Value: "a"}
	head2.Next = &Node{Value: "b"}
	head2.Next.Next = &Node{Value: "c"}
	
	newHead := LinkedListRecZipper(head1, head2)
	if newHead.Value != "1" {
		t.Errorf("Expected 1, got %s", newHead.Value)
	}
	if newHead.Next.Value != "a" {
		t.Errorf("Expected a, got %s", newHead.Next.Value)
	}
	if newHead.Next.Next.Value != "2" {
		t.Errorf("Expected 2, got %s", newHead.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Value != "b" {
		t.Errorf("Expected b, got %s", newHead.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Value != "3" {
		t.Errorf("Expected 3, got %s", newHead.Next.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Next.Value != "c" {
		t.Errorf("Expected c, got %s", newHead.Next.Next.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Next.Next.Value != "4" {
		t.Errorf("Expected 4, got %s", newHead.Next.Next.Next.Next.Next.Next.Value)
	}
}

func TestLinkedListItZipper(t *testing.T) {
	head1 := &Node{Value: "1"}
	head1.Next = &Node{Value: "2"}
	head1.Next.Next = &Node{Value: "3"}
	head1.Next.Next.Next = &Node{Value: "4"}

	head2 := &Node{Value: "a"}
	head2.Next = &Node{Value: "b"}
	head2.Next.Next = &Node{Value: "c"}
	
	newHead := LinkedListItZipper(head1, head2)
	if newHead.Value != "1" {
		t.Errorf("Expected 1, got %s", newHead.Value)
	}
	if newHead.Next.Value != "a" {
		t.Errorf("Expected a, got %s", newHead.Next.Value)
	}
	if newHead.Next.Next.Value != "2" {
		t.Errorf("Expected 2, got %s", newHead.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Value != "b" {
		t.Errorf("Expected b, got %s", newHead.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Value != "3" {
		t.Errorf("Expected 3, got %s", newHead.Next.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Next.Value != "c" {
		t.Errorf("Expected c, got %s", newHead.Next.Next.Next.Next.Next.Value)
	}
	if newHead.Next.Next.Next.Next.Next.Next.Value != "4" {
		t.Errorf("Expected 4, got %s", newHead.Next.Next.Next.Next.Next.Next.Value)
	}
}
