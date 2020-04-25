package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var last *ListNode
	var prev *ListNode
	var newHead *ListNode
	node := head
	counter := 1
	for node != nil {
		if counter%2 == 1 {
			prev = node
			node = node.Next
			counter++
		} else {
			temp := node.Next
			node.Next = prev
			if last == nil {
				last = prev
				newHead = node
			} else {
				last.Next = node
				last = prev
			}
			last.Next = nil
			node = temp
			counter++
		}
	}

	if prev != last {
		last.Next = prev
	}
	return newHead
}
