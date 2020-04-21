package binary_tree_level_order_traversal_102

import . "github.com/austingebauer/go-leetcode/structures"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		vals := make([]int, 0, size)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			vals = append(vals, node.Val)
		}
		queue = queue[size:]
		result = append(result, vals)
	}
	return result
}
