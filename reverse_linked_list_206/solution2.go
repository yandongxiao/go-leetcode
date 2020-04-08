package reverse_linked_list_206

import . "github.com/austingebauer/go-leetcode/structures"

func reverseList(head *ListNode) *ListNode {
	node := head
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}
