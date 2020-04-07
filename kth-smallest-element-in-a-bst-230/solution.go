package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	counter := 0
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			last := len(stack) - 1
			root = stack[last]
			counter++
			if counter == k {
				return root.Val
			}
			stack = stack[:last]
			root = root.Right
		}

	}
	return -1
}
