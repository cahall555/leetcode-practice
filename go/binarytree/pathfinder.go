package binarytree


// PathFinder returns a path from the root to the target node in a binary tree
// Time complexity: O(n)
// Space complexity: O(n)
func PathFinder(root *BTNode, target string) []string {
	path := []string{}
	pathFinderHelper(root, target, &path)
	return path
}

func pathFinderHelper(root *BTNode, target string, path *[]string) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root.Value)
	if root.Value == target {
		return true
	}
	if pathFinderHelper(root.Left, target, path) || pathFinderHelper(root.Right, target, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}

// root := &binarytree.BTNode{Value: "a"}                
// root.Left = &binarytree.BTNode{Value: "b"}	    
// root.Right = &binarytree.BTNode{Value: "c"}
// root.Left.Left = &binarytree.BTNode{Value: "d"}
// root.Left.Right = &binarytree.BTNode{Value: "e"}
// root.Right.Left = &binarytree.BTNode{Value: "f"}

//           a
//         /   \
//        b     c
//       / \   /
//      d   e f

// fmt.Println(binarytree.PathFinder(root, "d")) // [a b d]

//Iteration 1:
// *path = append([]string{}, "a")
// root.Value == "a" == "d" false
// pathFinderHelper(root.Left, "d", &path) || pathFinderHelper(root.Right, "d", &path) (false || false) == false
     //pathFinderHelper(root.Left, "d", &path)
     //*path = append([]string{"a"}, "b")
     //root.Value == "b" == "d" false

     //pathFinderHelper(root.Right, "d", &path)
     //*path = append([]string{"a"}, "c")
     //root.Value == "c" == "d" false


