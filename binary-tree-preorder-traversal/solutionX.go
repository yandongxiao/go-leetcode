package main

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			last := len(stack) - 1
			root = stack[last]
			stack = stack[:last]
			root = root.Right
		}
	}
	return result
}
