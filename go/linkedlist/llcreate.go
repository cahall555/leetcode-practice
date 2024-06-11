package linkedlist

//CreateLinkedListRec creates a linked list recursively.
//Time complexity is O(n)
//Space complexity is O(n)
func CreateLinkedListRec(values []string, i int) *Node {
	if len(values) == 0 || i == len(values) {
		return nil
	}
	head := &Node{Value: values[i]}
	head.Next = CreateLinkedListRec(values, i + 1)
	return head
}

//CreateLinkedListIt creates a linked list iteratively.
//Time complexity is O(n)
//Space complexity is O(n)
func CreateLinkedListIt(values []string) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{Value: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		newNode := &Node{Value: values[i]}
		current.Next = newNode
		current = newNode
	}
	return head
}
