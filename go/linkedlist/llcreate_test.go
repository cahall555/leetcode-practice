package linkedlist

import (
	"testing"
)


func TestCreateLinkedListRec(t *testing.T) {
	//Test 1
	test1 := CreateLinkedListRec([]string{"a", "b", "c"}, 0)
	if test1.Value != "a" {
		t.Error("For test 1, expected a but got ", test1.Value)
	}
	if test1.Next.Value != "b" {
		t.Error("For test 1, expected b but got ", test1.Next.Value)
	}
	if test1.Next.Next.Value != "c" {
		t.Error("For test 1, expected c but got ", test1.Next.Next.Value)
	}
	if test1.Next.Next.Next != nil {
		t.Error("For test 1, expected nil but got ", test1.Next.Next.Next)
	}

	//Test 2
	test2 := CreateLinkedListRec([]string{"d"}, 0)

	if test2.Value != "d" {
		t.Error("For test 2, expected d but got ", test2.Value)
	}
	if test2.Next != nil {
		t.Error("For test 2, expected nil but got ", test2.Next)
	}

	//Test 3
	test3 := CreateLinkedListRec([]string{}, 0)

	if test3 != nil {
		t.Error("For test 3, expected nil but got ", test3)
	}
}


func TestCreateLinkedListIt(t *testing.T) {
	//Test 1
	test1 := CreateLinkedListIt([]string{"a", "b", "c"})
	if test1.Value != "a" {
		t.Error("For test 1, expected a but got ", test1.Value)
	}
	if test1.Next.Value != "b" {
		t.Error("For test 1, expected b but got ", test1.Next.Value)
	}
	if test1.Next.Next.Value != "c" {
		t.Error("For test 1, expected c but got ", test1.Next.Next.Value)
	}
	if test1.Next.Next.Next != nil {
		t.Error("For test 1, expected nil but got ", test1.Next.Next.Next)
	}

	//Test 2
	test2 := CreateLinkedListIt([]string{"d"})

	if test2.Value != "d" {
		t.Error("For test 2, expected d but got ", test2.Value)
	}
	if test2.Next != nil {
		t.Error("For test 2, expected nil but got ", test2.Next)
	}

	//Test 3
	test3 := CreateLinkedListIt([]string{})

	if test3 != nil {
		t.Error("For test 3, expected nil but got ", test3)
	}
}
