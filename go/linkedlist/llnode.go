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

//Print values of reversed linked list
func PrintLinkedList(head *Node) {
	for head != nil {
		print(head.Value, " ")
		head = head.Next
	}
}
