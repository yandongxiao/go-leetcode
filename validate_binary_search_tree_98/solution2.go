package validate_binary_search_tree_98

func isValidBST(root *TreeNode) bool {
	return valid(root, nil, nil)
}

func valid(root *TreeNode, leftMax *int, rightMin *int) bool {
	if root == nil {
		return true
	}

	if !valid(root.Left, leftMax, &root.Val) {
		return false
	}

	if leftMax != nil && *leftMax >= root.Val {
		return false
	}

	if rightMin != nil && *rightMin <= root.Val {
		return false
	}

	if !valid(root.Right, &root.Val, rightMin) {
		return false
	}
	return true
}
