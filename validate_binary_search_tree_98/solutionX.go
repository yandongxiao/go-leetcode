package validate_binary_search_tree_98

import . "github.com/austingebauer/go-leetcode/structures"

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root *TreeNode, left, right *int) bool {
	if root == nil {
		return true
	}

	if !helper(root.Left, left, &root.Val) {
		return false
	}

	if left != nil && *left >= root.Val {
		return false
	}

	if right != nil && *right =< root.Val {
		return false
	}

	if !helper(root.Right, &root.Val, right) {
		return false
	}

	return true
}
