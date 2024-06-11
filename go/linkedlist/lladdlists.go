package linkedlist

func AddListsRec(a, b *IntNode, carry int) *IntNode {
	if a ==nil && b == nil && carry == 0 {
		return nil
	}
	value := carry
	if a != nil {
		value += a.Value
	}
	if b != nil {
		value += b.Value
	}
	result := &IntNode{Value: value % 10}
	if a != nil || b != nil {
		var nextA, nextB *IntNode
		if a != nil {
			nextA = a.Next
		}
		if b != nil {
			nextB = b.Next
		}
		var nextCarry int
		if value >= 10 {
			nextCarry = 1
		}
		result.Next = AddListsRec(nextA, nextB, nextCarry)
	}
	return result
}

func AddListsIt(a, b *IntNode) *IntNode {
	carry := 0
	var head, current *IntNode
	for a != nil || b != nil || carry != 0 {
		value := carry
		if a != nil {
			value += a.Value
			a = a.Next
		}
		if b != nil {
			value += b.Value
			b = b.Next
		}
		newNode := &IntNode{Value: value % 10}
		if head == nil {
			head = newNode
		} else {
			current.Next = newNode
		}
		current = newNode
		if value >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}
	return head
}
