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

//Print values of string linked list
func PrintLinkedList(head *Node) string {
	result := ""
	current := head
	for current != nil {
		result += current.Value + " "
		current = current.Next
	}
	return result
}

func PrintIntLinkedList(head *IntNode) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
