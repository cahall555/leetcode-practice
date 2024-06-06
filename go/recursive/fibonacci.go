package recursive

// O(2^n) time complexity and O(n) space complexity

//                     5       Time complexity can be visualized by thinking of a binary tree
//                   /   \
//                 4       3
//               /   \    /  \
//             3      2  2    1  space complexity can be visualized by thinking of the call stack, which is seen by looking at the far left side of the tree
//           /   \   / \	
//         2     1  1  0		
//       /  \
//     1     0


func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
