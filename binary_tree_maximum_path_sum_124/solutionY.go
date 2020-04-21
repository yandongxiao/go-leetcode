package binary_tree_maximum_path_sum_124

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

var maxValue int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue = root.Val
	helper(root)
	return maxValue
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(helper(root.Left), 0)
	right := max(helper(root.Right), 0)
	v := left + right + root.Val
	if maxValue < v {
		maxValue = v
	}
	return root.Val + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
