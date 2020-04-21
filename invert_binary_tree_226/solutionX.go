package invert_binary_tree_226

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return root
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}
