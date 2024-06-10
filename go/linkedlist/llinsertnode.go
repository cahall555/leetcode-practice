package linkedlist

//InsertNodeRec inserts a new node with the target value at the given index in the linked list recursively.
//Time complexity is O(n)
//Space complexity is O(n)

func InsertNodeRec(head *Node, target string, index int, i int) *Node {
	if head == nil {
		return nil
	}
	if index == 0 {
		newHead := &Node{Value: target, Next: head}
		return newHead
	}
	
	if i == index - 1 {
		newNode := &Node{Value: target, Next: head.Next}
		head.Next = newNode
	}
	
	InsertNodeRec(head.Next, target, index, i + 1)
	return head
}

//InsertNodeIt inserts a new node with the target value at the given index in the linked list iteratively.
//Time complexity is O(n)
//Space complexity is O(1)

func InsertNodeIt(head *Node, target string, index int) *Node {
	if head == nil {
		return nil
	}
	if index == 0 {
		newHead := &Node{Value: target, Next: head}
		return newHead
	}
	current := head
	i := 0
	for current != nil {
		if i == index - 1 {
			newNode := &Node{Value: target, Next: current.Next}
			current.Next = newNode
		}
		i++
		current = current.Next
	}

	return head
}
