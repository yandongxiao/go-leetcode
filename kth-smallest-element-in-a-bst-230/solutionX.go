package main

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root
	counter := 0
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		last := len(stack) - 1
		node = stack[last]
		stack = stack[:last]
		counter++
		if counter == k {
			return node.Val
		}
		node = node.Right
	}
	return -1
}
