package binarytree

//Using stack methodology to traverse binary tree.



//DepthFirstRec is a recursive function that traverses a binary tree in depth-first order.
//Time complexity is O(n^2)
//Space complexity is O(n)
func DepthFirstRec(root *BTNode) []string {
	result := []string{}
	if root == nil {
		return result
	}
	result = append(result, root.Value)
	result = append(result, DepthFirstRec(root.Left)...)
	result = append(result, DepthFirstRec(root.Right)...)
	return result
}

//DepthFirstIt is an iterative function that traverses a binary tree in depth-first order.
//Time complexity is O(n)
//Space complexity is O(n)
func DepthFirstIt(root *BTNode) []string {
	result := []string{}
	if root == nil {
		return result
	}
	stack := []*BTNode{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Value)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return result
}

