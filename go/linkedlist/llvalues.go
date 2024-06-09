package linkedlist //Article with linked lists in Go: https://edwinsiby.medium.com/understanding-linked-lists-in-go-a-comprehensive-guide-for-all-skill-levels-1e3d91a24d08


func LinkedListValues(head *Node) []string {
	var values []string
	fillValues(head, &values)
	return values
} 

func fillValues(head *Node, values *[]string) {
	if head == nil {
		return
	}
	*values = append(*values, head.Value)
	fillValues(head.Next, values)
}

