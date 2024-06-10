package linkedlist

// LongestStreak returns the length of the longest streak of consecutive equal nodes in the linked list.
// time complexity is O(n)
// space complexity is O(1)
func LongestStreak(head *IntNode) int {
	if head == nil {
		return 0
	}
	current := head
	prev := head.Value
	count := 1
	longest := 1
	for current.Next != nil {
		current = current.Next
		if prev == current.Value {
			count++
		} else {
			count = 1
		}
		if count > longest {
			longest = count
		}
		
		prev = current.Value
	}
	return longest
}
