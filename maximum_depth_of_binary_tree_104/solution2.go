package maximum_depth_of_binary_tree_104

func maxDepth(root *TreeNode) int {
	max := 0
	helper(root, &max)
	return *max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	lmax := helper(root.Left, max)
	rmax := helper(root.Right, max)
	depth := lmax + 1
	if rmax > lmax {
		depth = rmax + 1
	}

	if *max < depth {
		*max = depth
	}

	return depth
}
