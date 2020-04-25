package reverse_linked_list_206

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	node := head
	for node != nil {
		temp := node.Next
		node.next = prev
		prev = node
		node = temp
	}
	return prev
}
