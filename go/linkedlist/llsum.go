package linkedlist

func LinkedListSum(head *IntNode) int {
	if head == nil {
		return 0
	}
	return head.Value + LinkedListSum(head.Next)
}
