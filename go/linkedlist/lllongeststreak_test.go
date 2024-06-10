package linkedlist

import (
	"testing"
)

func TestLongestStreak(t *testing.T) {
	//Test 1
	test1 := &IntNode{Value: 5}
	test1.Next = &IntNode{Value: 5}
	test1.Next.Next = &IntNode{Value: 7}
	test1.Next.Next.Next = &IntNode{Value: 7}
	test1.Next.Next.Next.Next = &IntNode{Value: 7}
	test1.Next.Next.Next.Next.Next = &IntNode{Value: 6}

	result1 := LongestStreak(test1)
	expected1 := 3
	if result1 != expected1 {
		t.Errorf("Expected %d but got %d", expected1, result1)
	}

	//Test 2
	test2 := &IntNode{Value: 3}
	test2.Next = &IntNode{Value: 3}
	test2.Next.Next = &IntNode{Value: 3}
	test2.Next.Next.Next = &IntNode{Value: 3}
	test2.Next.Next.Next.Next = &IntNode{Value:9}
	test2.Next.Next.Next.Next.Next = &IntNode{Value: 9}

	result2 := LongestStreak(test2)
	expected2 := 4
	if result2 != expected2 {
		t.Errorf("Expected %d but got %d", expected2, result2)
	}

	//Test 3
	test3 := &IntNode{Value: 9}
	test3.Next = &IntNode{Value: 9}
	test3.Next.Next = &IntNode{Value: 1}
	test3.Next.Next.Next = &IntNode{Value: 9}
	test3.Next.Next.Next.Next = &IntNode{Value: 9}
	test3.Next.Next.Next.Next.Next = &IntNode{Value: 9}

	result3 := LongestStreak(test3)
	expected3 := 3
	if result3 != expected3 {
		t.Errorf("Expected %d but got %d", expected3, result3)
	}

	//Test 4
	test4 := &IntNode{Value: 5}
	test4.Next = &IntNode{Value: 5}

	result4 := LongestStreak(test4)
	expected4 := 2
	if result4 != expected4 {
		t.Errorf("Expected %d but got %d", expected4, result4)
	}

	//Test 5
	test5 := &IntNode{Value: 4}

	result5 := LongestStreak(test5)
	expected5 := 1
	if result5 != expected5 {
		t.Errorf("Expected %d but got %d", expected5, result5)
	}

	//Test 6
	result6 := LongestStreak(nil)
	expected6 := 0
	if result6 != expected6 {
		t.Errorf("Expected %d but got %d", expected6, result6)
	}

}
