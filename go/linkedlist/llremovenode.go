package linkedlist

func RemoveNodeRec(head *Node, target string) *Node {
	if head == nil {
		return nil
	}
	if head.Value == target {
		return head.Next
	}
	head.Next = RemoveNodeRec(head.Next, target)
	return head
}

//target c
//a -> b -> c -> d -> e
//current = a
//prev = a

func RemoveNodeIt(head *Node, target string) *Node {
	if head == nil {
		return nil
	}
	if head.Value == target {
		return head.Next
	}
	current := head
	prev := head

	for current != nil {
		if current.Value == target {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
	}
	return head
}
