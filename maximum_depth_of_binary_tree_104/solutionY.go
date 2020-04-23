package maximum_depth_of_binary_tree_104

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	backup := root
	stack := []*TreeNode{}
	visited := map[*TreeNode]struct{}{}
	result := map[*TreeNode]int{}

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			last := len(stack) - 1
			root = stack[last]
			stack = stack[:last]

			if _, ok := visited[root]; !ok {
				visited[root] = struct{}{}
				stack = append(stack, root)
				root = root.Right
			} else {
				result[root] = max(result[root.Left], result[root.Right]) + 1
				root = nil
			}
		}
	}

	return result[backup]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
