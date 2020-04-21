package subtree_of_another_tree_572

import (
	. "github.com/austingebauer/go-leetcode/structures"
)

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		if t == nil {
			return true
		}
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

	return equal(s.Left, t.Left) && equal(s.Right, t.Right)
}
