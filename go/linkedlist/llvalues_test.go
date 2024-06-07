package linkedlist

import (
	"testing"
)

func TestLinkedListValues( t *testing.T) {
	//Test 1
	a := Node{Value: "a"}
	b := Node{Value: "b"}
	c := Node{Value: "c"}
	d := Node{Value: "d"}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	result := LinkedListValues(&a)
	expected := []string{"a", "b", "c", "d"}

	if !compareSlices(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	//Test 2
	x := Node{Value: "x"}
	y := Node{Value: "y"}

	x.Next = &y
	result = LinkedListValues(&x)
	expected = []string{"x", "y"}
	if !compareSlices(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	
	//Test 3
	q:= Node{Value: "q"}
	result = LinkedListValues(&q)
	expected = []string{"q"}
	if !compareSlices(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	//Test 4
	result = LinkedListValues(nil)
	expected = []string{}
	if !compareSlices(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

}



func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
