package binarytree

// BreadthFirst traverses the binary tree in breadth-first order
//Time complexity: O(n)
//Space complexity: O(n)

func BreadthFirst(root *BTNode) []string {
	var result []string
	if root == nil {
		return result
	}

	queue := []*BTNode{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return result
}

//           a
//         /   \
//       b       c
//      / \     / \
//     d   e   f   g
//    / \
//   h   i


// BreadthFirst(a) => [a b c d e f g h i]

// itteraion 1:
// current := queue[0] => a
// queue = queue[1:] => []
// result = append(result, a) => [a]
// queue = append(queue, b) => [b]
// queue = append(queue, c) => [b c]

// itteraion 2:
// current := queue[0] => b
// queue = queue[1:] => [c]
// result = append(result, b) => [a b]
// queue = append(queue, d) => [c d]
// queue = append(queue, e) => [c d e]

// itteraion 3:
// current := queue[0] => c
// queue = queue[1:] => [d e]
// result = append(result, c) => [a b c]
// queue = append(queue, f) => [d e f]
// queue = append(queue, g) => [d e f g]
//...
