package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	lt := lowestCommonAncestor(root.Left, p, q)
	rt := lowestCommonAncestor(root.Right, p, q)
	if lt == nil {
		return rt
	} else if rt == nil {
		return lt
	}
	return root
}
