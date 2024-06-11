package linkedlist

import (
	"testing"
	"reflect"
)

func TestAddListsRec(t *testing.T) {
	//Test 1
	testa1 := &IntNode{Value: 1}
	testa1.Next = &IntNode{Value: 2}
	testa1.Next.Next = &IntNode{Value: 6}

	testb1 := &IntNode{Value: 4}
	testb1.Next = &IntNode{Value: 5}
	testb1.Next.Next = &IntNode{Value: 3}

	result1 := AddListsRec(testa1, testb1, 0)
	expected1 := []int{5, 7, 9}
	if !reflect.DeepEqual(PrintIntLinkedList(result1), expected1) {
		t.Errorf("For test 1, expected %v but got %v", expected1, PrintIntLinkedList(result1))
	}

	//Test 2
	testa2 := &IntNode{Value: 1}
	testa2.Next = &IntNode{Value: 4}
	testa2.Next.Next = &IntNode{Value: 5}
	testa2.Next.Next.Next = &IntNode{Value: 7}

	testb2 := &IntNode{Value: 2}
	testb2.Next = &IntNode{Value: 3}

	result2 := AddListsRec(testa2, testb2, 0)
	expected2 := []int{3, 7, 5, 7}
	if !reflect.DeepEqual(PrintIntLinkedList(result2), expected2) {
		t.Errorf("For test 2, expected %v but got %v", expected2, PrintIntLinkedList(result2))
	}
	//Test 3
	testa3 := &IntNode{Value: 9}
	testa3.Next = &IntNode{Value: 8}

	testb3 := &IntNode{Value: 7}
	testb3.Next = &IntNode{Value: 4}

	result3 := AddListsRec(testa3, testb3, 0)
	expected3 := []int{6, 3, 1}
	if !reflect.DeepEqual(PrintIntLinkedList(result3), expected3) {
		t.Errorf("For test 3, expected %v but got %v", expected3, PrintIntLinkedList(result3))
	}

	//Test 4
	testa4 := &IntNode{Value: 9}
	testa4.Next = &IntNode{Value: 9}
	testa4.Next.Next = &IntNode{Value: 9}

	testb4 := &IntNode{Value: 6}

	result4 := AddListsRec(testa4, testb4, 0)
	expected4 := []int{5, 0, 0, 1}
	if !reflect.DeepEqual(PrintIntLinkedList(result4), expected4) {
		t.Errorf("For test 4, expected %v but got %v", expected4, PrintIntLinkedList(result4))
	}

}

func TestAddListsIt(t *testing.T) {
	//Test 1
	testa1 := &IntNode{Value: 1}
	testa1.Next = &IntNode{Value: 2}
	testa1.Next.Next = &IntNode{Value: 6}

	testb1 := &IntNode{Value: 4}
	testb1.Next = &IntNode{Value: 5}
	testb1.Next.Next = &IntNode{Value: 3}

	result1 := AddListsIt(testa1, testb1)
	expected1 := []int{5, 7, 9}
	if !reflect.DeepEqual(PrintIntLinkedList(result1), expected1) {
		t.Errorf("For test 1, expected %v but got %v", expected1, PrintIntLinkedList(result1))
	}

	//Test 2
	testa2 := &IntNode{Value: 1}
	testa2.Next = &IntNode{Value: 4}
	testa2.Next.Next = &IntNode{Value: 5}
	testa2.Next.Next.Next = &IntNode{Value: 7}

	testb2 := &IntNode{Value: 2}
	testb2.Next = &IntNode{Value: 3}

	result2 := AddListsIt(testa2, testb2)
	expected2 := []int{3, 7, 5, 7}
	if !reflect.DeepEqual(PrintIntLinkedList(result2), expected2) {
		t.Errorf("For test 2, expected %v but got %v", expected2, PrintIntLinkedList(result2))
	}

	//Test 3
	testa3 := &IntNode{Value: 9}
	testa3.Next = &IntNode{Value: 8}

	testb3 := &IntNode{Value: 7}
	testb3.Next = &IntNode{Value: 4}

	result3 := AddListsIt(testa3, testb3)
	expected3 := []int{6, 3, 1}
	if !reflect.DeepEqual(PrintIntLinkedList(result3), expected3) {
		t.Errorf("For test 3, expected %v but got %v", expected3, PrintIntLinkedList(result3))
	}

	//Test 4
	testa4 := &IntNode{Value: 9}
	testa4.Next = &IntNode{Value: 9}
	testa4.Next.Next = &IntNode{Value: 9}

	testb4 := &IntNode{Value: 6}

	result4 := AddListsIt(testa4, testb4)
	expected4 := []int{5, 0, 0, 1}
	if !reflect.DeepEqual(PrintIntLinkedList(result4), expected4) {
		t.Errorf("For test 4, expected %v but got %v", expected4, PrintIntLinkedList(result4))
	}
}


