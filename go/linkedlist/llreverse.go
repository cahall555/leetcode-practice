package linkedlist

//Recursive version of linked list reverse
//Time complexity: O(n)
//Space complexity: O(n)
func LinkedListReverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	reverse := LinkedListReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	
	return reverse
}
//Iterative version of linked list reverse
//Time complexity: O(n)
//Space complexity: O(1)
func LinkedListItReverse(head *Node) *Node {
	var prev, current *Node = nil, head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

