package linkedlist

//LinkedListRecZipper zips two linked lists recursively
//Time complexity: O(min(n, m)
//Space complexity: O(min(n, m))
func LinkedListRecZipper(head1, head2 *Node) *Node {
	if head1 == nil && head2 == nil {
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	next1 := head1.Next
	next2 := head2.Next

	head1.Next = head2
	head2.Next = LinkedListRecZipper(next1, next2)
	return head1
}

//LinkedListItZipper zips two linked lists iteratively
//Time complexity: O(min(n, m)
//Space complexity: O(1)
func LinkedListItZipper(head1, head2 *Node) *Node {
	head := head1
	tail := head
	current1 := head1.Next
	current2 := head2
	count := 0
	for current1 != nil && current2 != nil {
		if count % 2 == 0 {
			tail.Next = current2
			current2 = current2.Next
		} else {
			tail.Next = current1
			current1 = current1.Next
		}
		tail = tail.Next
		count++
	}
	if current1 != nil {
		tail.Next = current1
	}
	if current2 != nil {
		tail.Next = current2
	}
	return head
}

