package binary_tree_maximum_path_sum_124

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := math.MinInt32
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	lsum := Max(helper(root.Left, max), 0)
	rsum := Max(helper(root.Right, max), 0)
	pathSum := root.Val + lsum + rsum
	if pathSum > *max {
		*max = pathSum
	}

	return root.Val + Max(lsum, rsum)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
