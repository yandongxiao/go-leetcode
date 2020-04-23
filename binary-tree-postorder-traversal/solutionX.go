package main

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	visited := map[*TreeNode]struct{}{}

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = apped(stack, root)
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
				result = append(result, root.Val)
				root = nil
			}
		}
	}
	return result
}
