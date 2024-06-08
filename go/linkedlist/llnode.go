package linkedlist

type Node struct {
	Value string
	Next  *Node
}

type IntNode struct {
	Value int
	Next  *IntNode
}

type LinkedList struct {
	Head *Node
}
