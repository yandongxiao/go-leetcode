package binary_tree_maximum_path_sum_124

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := root.Val
	helper(root, &maxVal)
	return maxVal
}

func helper(root *TreeNode, maxVal *int) int {
	if root == nil {
		return 0
	}

	v1 := max(helper(root.Left, maxVal), 0)
	v2 := max(helper(root.Right, maxVal), 0)
	v := v1 + v2 + root.Val
	if *maxVal < v {
		*maxVal = v
	}
	return root.Val + max(v1, v2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
