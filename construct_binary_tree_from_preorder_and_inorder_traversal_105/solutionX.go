package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	idx := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			idx = i
		}
	}

	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
