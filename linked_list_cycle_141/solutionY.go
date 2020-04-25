package linked_list_cycle_141

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next
	for fast != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next
		if fast == nil {
			return false
		} else if fast == slow {
			return true
		}
		fast = fast.Next
		slow = slow.Next
	}
	return false
}
