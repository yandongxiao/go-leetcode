package maximum_depth_of_binary_tree_104

import . "github.com/austingebauer/go-leetcode/structures"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	v1 := maxDepth(root.Left)
	v2 := maxDepth(root.Right)
	if v1 > v2 {
		return v1 + 1
	}
	return v2 + 1
}
