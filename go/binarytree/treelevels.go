package binarytree


// Breadth first recursive traversal of a binary tree
// Time complexity: O(n)
// Space complexity: O(n)
func TreeLevelsRec(root *BTNode) [][]string {
	levels := [][]string{}
	treeLevelsRecHelper(root, 0, &levels)
	return levels
}

func treeLevelsRecHelper(root *BTNode, level int, levels *[][]string) {
	if root == nil || root.Value == "" {
		return 
	}

	if len(*levels) == level {
		*levels = append(*levels, []string{root.Value})
	} else {
		(*levels)[level] = append((*levels)[level], root.Value)
	}

	treeLevelsRecHelper(root.Left, level+1, levels)
	treeLevelsRecHelper(root.Right, level+1, levels)
}


//         a
//       /   \
//     b       c
//    / \     / \
//   d   e   f   g
//  / \
// h   i

// TreeLevelsRec(a) => [[a] [b c] [d e f g] [h i]]

// itteraion 1: treeLevelsRecHelper(a, 0, levels)
// level := 0
// len(levels) = 0
// len(levels) == level => true
// levels = append(levels, [a]) => [[a]]
// treeLevelsRecHelper(b, 1, levels)
// treeLevelsRecHelper(c, 1, levels)

// itteraion 2: treeLevelsRecHelper(b, 1, levels)
// level := 1
// len(levels) = 1
// len(levels) == level => true
// levels[level] = append(levels[level], b) => [[a] [b]]
// treeLevelsRecHelper(d, 2, levels)
// treeLevelsRecHelper(e, 2, levels)

// treeLevelsRecHelper(c, 1, levels) (parallel to call to treeLevelsRecHelper(b, 1, levels))
// level := 1
// len(levels) = 2
// len(levels) == level => false
// levels[level] = append(levels[level], c) => [[a] [b c]]
// treeLevelsRecHelper(f, 2, levels)
// treeLevelsRecHelper(g, 2, levels)

// itteraion 3: treeLevelsRecHelper(d, 2, levels)
// level := 2
// len(levels) = 2
// len(levels) == level => true
// levels[level] = append(levels[level], d) => [[a] [b c] [d]]
// treeLevelsRecHelper(h, 3, levels)
// treeLevelsRecHelper(i, 3, levels)

//...
