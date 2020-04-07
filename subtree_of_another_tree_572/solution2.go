package subtree_of_another_tree_572

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return equal(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equal(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	if !equal(s.Left, t.Left) {
		return false
	}

	if !equal(s.Right, t.Right) {
		return false
	}

	return true
}
