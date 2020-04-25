package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var newHead *ListNode
	var last *ListNode
	var prev *ListNode
	node := head
	counter := 1
	for node != nil {
		if counter%2 == 1 {
			prev = node
			node = node.Next
			counter++
		} else {
			if newHead == nil {
				newHead = node
				last = prev
			} else {
				last.Next = node
				last = prev
			}
			last.Next = nil
			temp := node.Next
			node.Next = prev
			node = temp
			counter++
		}
	}
	if prev != last {
		last.Next = prev
	}
	return newHead
}
