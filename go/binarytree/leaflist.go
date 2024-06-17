package binarytree

// LeafListRec returns a list of leaf nodes in a binary tree recursively
// Time complexity: O(n)
// Space complexity: O(n)
func LeafListRec(root *BTNode) []string {
	leaves := []string{}
	leafListHelper(root, &leaves)
	return leaves
}

func leafListHelper(root *BTNode, leaves *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Value)
	}
	leafListHelper(root.Left, leaves)
	leafListHelper(root.Right, leaves)
}


// LeafListIt returns a list of leaf nodes in a binary tree iteratively
// Time complexity: O(n)
// Space complexity: O(n)
func LeafListIt(root *BTNode) []string {
	leaves := []string{}
	if root == nil {
		return leaves
	}
	stack := []*BTNode{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left == nil && current.Right == nil {
			leaves = append(leaves, current.Value)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return leaves
}
