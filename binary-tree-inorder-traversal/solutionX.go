package main

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			last := len(stack) - 1
			root = stack[last]
			stack = stack[:last]
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
}
