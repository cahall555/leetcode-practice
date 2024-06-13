package main

import (
	"fmt"
	"example.com/dynamic"
	"example.com/recursive"
	"example.com/linkedlist"
	"example.com/binarytree"
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
	
	println(linkedlist.PrintLinkedList(linkedlist.LinkedListReverse(&a)))

	e := linkedlist.Node{Value: "e"}
	f := linkedlist.Node{Value: "f"}
	g := linkedlist.Node{Value: "g"}
	h := linkedlist.Node{Value: "h"}

	e.Next = &f
	f.Next = &g
	g.Next = &h
	println(linkedlist.PrintLinkedList(linkedlist.LinkedListItReverse(&e)))
	
	head1 := &linkedlist.Node{Value: "1"}
	head1.Next = &linkedlist.Node{Value: "3"}
	head1.Next.Next = &linkedlist.Node{Value: "5"}

	head2 := &linkedlist.Node{Value: "2"}
	head2.Next = &linkedlist.Node{Value: "4"}
	head2.Next.Next = &linkedlist.Node{Value: "6"}
	
	newHead := linkedlist.LinkedListRecZipper(head1, head2)
	println(linkedlist.PrintLinkedList(newHead))


//	itNewHead := linkedlist.LinkedListItZipper(head1, head2)
//	println(linkedlist.PrintLinkedList(itNewHead))

	lshead := &linkedlist.IntNode{Value: 1}
	lshead.Next = &linkedlist.IntNode{Value: 1}
	lshead.Next.Next = &linkedlist.IntNode{Value: 2}
	lshead.Next.Next.Next = &linkedlist.IntNode{Value: 3}
	lshead.Next.Next.Next.Next = &linkedlist.IntNode{Value: 3}
	lshead.Next.Next.Next.Next.Next = &linkedlist.IntNode{Value: 3}

	println(linkedlist.LongestStreak(lshead))

	remove := &linkedlist.Node{Value: "a"}
	remove.Next = &linkedlist.Node{Value: "b"}
	remove.Next.Next = &linkedlist.Node{Value: "c"}
	remove.Next.Next.Next = &linkedlist.Node{Value: "d"}
	remove.Next.Next.Next.Next = &linkedlist.Node{Value: "c"}
	remove.Next.Next.Next.Next.Next = &linkedlist.Node{Value: "e"}

	println(linkedlist.PrintLinkedList(linkedlist.RemoveNodeRec(remove, "c")))
//	println(linkedlist.PrintLinkedList(linkedlist.RemoveNodeIt(remove, "c")))

        insert := &linkedlist.Node{Value: "a"}
	insert.Next = &linkedlist.Node{Value: "b"}
	insert.Next.Next = &linkedlist.Node{Value: "d"}
	insert.Next.Next.Next = &linkedlist.Node{Value: "e"}

//	println(linkedlist.PrintLinkedList(linkedlist.InsertNodeRec(insert, "c", 2, 0)))
	println(linkedlist.PrintLinkedList(linkedlist.InsertNodeIt(insert, "c", 2)))

	values := []string{"h", "i", "j", "k", "l"}
//	println(linkedlist.PrintLinkedList(linkedlist.CreateLinkedListRec(values, 0)))
	println(linkedlist.PrintLinkedList(linkedlist.CreateLinkedListIt(values)))

	addA := &linkedlist.IntNode{Value: 7}
	addA.Next = &linkedlist.IntNode{Value: 1}
	addA.Next.Next = &linkedlist.IntNode{Value: 6}

	addB := &linkedlist.IntNode{Value: 5}
	addB.Next = &linkedlist.IntNode{Value: 9}
	addB.Next.Next = &linkedlist.IntNode{Value: 2}
	addB.Next.Next.Next = &linkedlist.IntNode{Value: 3}

//	fmt.Println(linkedlist.PrintIntLinkedList(linkedlist.AddListsRec(addA, addB, 0)))
	fmt.Println(linkedlist.PrintIntLinkedList(linkedlist.AddListsIt(addA, addB)))

	root := &binarytree.BTNode{Value: "a"}
	root.Left = &binarytree.BTNode{Value: "b"}
	root.Right = &binarytree.BTNode{Value: "c"}
	root.Left.Left = &binarytree.BTNode{Value: "d"}
	root.Left.Right = &binarytree.BTNode{Value: "e"}
	root.Right.Left = &binarytree.BTNode{Value: "f"}
	
//	fmt.Println(binarytree.DepthFirstRec(root))
	fmt.Println(binarytree.DepthFirstIt(root))

	fmt.Println(binarytree.BreadthFirst(root))

	fmt.Println(binarytree.TreeLevelsRec(root))
}


