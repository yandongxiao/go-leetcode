package main

func detectCycle(head *ListNode) *ListNode {
	point := meet(head)
	if point == nil {
		return nil
	}

	node := head
	for node != point {
		node = node.Next
		point = point.Next
	}
	return node
}

func meet(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			return fast
		}
	}
	return nil
}
