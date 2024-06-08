package linkedlist

func LinkedListFind(head *Node, target string) bool {
	if head == nil {
		return false
	}
	if head.Value == target {
		return true
	}
	return LinkedListFind(head.Next, target)
}
