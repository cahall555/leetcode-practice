package binarytree


func Cousins(root *BTIntNode, x, y int) bool {
	xParent, xLevel, _ := isCousinHelper(root, x, 0)
	yParent, yLevel, _ := isCousinHelper(root, y, 0)

	if xParent != yParent && xLevel == yLevel {
		return true
	}

	return false
}

func isCousinHelper(root *BTIntNode, target int, level int) (*BTIntNode, int, bool) {
	if root == nil {
		return nil, -1, false 
	}
	if (root.Left != nil && root.Left.Value == target) || (root.Right != nil && root.Right.Value == target) {
		return root, level + 1, true
	}

	lParent, lLevel, lFound := isCousinHelper(root.Left, target, level+1)

	if lFound {
		return lParent, lLevel, true
	}

	rParent, rLevel, rFound := isCousinHelper(root.Right, target, level+1)

	if rFound {
		return rParent, rLevel, true
	}

	return nil, -1, false

}
