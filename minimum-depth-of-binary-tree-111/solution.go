package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)
	if ld == 0 {
		return rd + 1
	} else if rd == 0 {
		return ld + 1
	} else if ld > rd {
		return rd + 1
	}
	return ld + 1
}
