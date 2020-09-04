package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	var newHead *ListNode
	var last *ListNode
	var begin *ListNode
	var end *ListNode

	node := head
	counter := 0
	for node != nil {
		temp := node.Next
		flag := counter % k
		if flag == 0 {
			begin = node
		} else if flag == k-1 {
			end = node
			end.Next = nil
			reverse(begin)
			if newHead == nil {
				newHead = end
				last = begin
			} else {
				last.Next = end
				last = begin
			}
			last.Next = nil
		}
		node = temp
		counter++
	}

	if last != begin {
		last.Next = begin
	}
	return newHead
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
