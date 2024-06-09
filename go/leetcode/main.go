package main

import (
	"fmt"
	"example.com/dynamic"
	"example.com/recursive"
	"example.com/linkedlist"
)

func main() {
	fmt.Println(dynamic.Tribonacci(14))

	fmt.Println(recursive.SumNumbers([]int{1, 2, 3, 4, 5}))

	fmt.Println(recursive.Factorial(5))

	fmt.Println(recursive.SumOfLengths([]string{"Luna", "Maci", "Teddy"}))

	fmt.Println(recursive.ReverseString("Hello"))

	fmt.Println(recursive.Palindrome("racecar"))

	fmt.Println(recursive.Fibonacci(6))

	a := linkedlist.Node{Value: "a"}
	b := linkedlist.Node{Value: "b"}
	c := linkedlist.Node{Value: "c"}
	d := linkedlist.Node{Value: "d"}

	a.Next = &b
	b.Next = &c
	c.Next = &d

	fmt.Println(linkedlist.LinkedListValues(&a))

	fmt.Println(linkedlist.LinkedListFind(&a, "c"))

	one := linkedlist.IntNode{Value: 2}
	two := linkedlist.IntNode{Value: 8}
	three := linkedlist.IntNode{Value: 3}
	four := linkedlist.IntNode{Value: -1}
	five := linkedlist.IntNode{Value: 7}

	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five

	fmt.Println(linkedlist.LinkedListSum(&one))
	
	linkedlist.PrintLinkedList(linkedlist.LinkedListReverse(&a))
	println("")

	e := linkedlist.Node{Value: "e"}
	f := linkedlist.Node{Value: "f"}
	g := linkedlist.Node{Value: "g"}
	h := linkedlist.Node{Value: "h"}

	e.Next = &f
	f.Next = &g
	g.Next = &h
	linkedlist.PrintLinkedList(linkedlist.LinkedListItReverse(&e))
	
	println("")
	head1 := &linkedlist.Node{Value: "1"}
	head1.Next = &linkedlist.Node{Value: "3"}
	head1.Next.Next = &linkedlist.Node{Value: "5"}

	head2 := &linkedlist.Node{Value: "2"}
	head2.Next = &linkedlist.Node{Value: "4"}
	head2.Next.Next = &linkedlist.Node{Value: "6"}
	
	newHead := linkedlist.LinkedListRecZipper(head1, head2)
	linkedlist.PrintLinkedList(newHead)


//	itNewHead := linkedlist.LinkedListItZipper(head1, head2)
//	linkedlist.PrintLinkedList(itNewHead)
}


