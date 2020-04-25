package main

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    } else if head.Next == nil {
        return head
    }

    prev := head
    node := head.Next
    newHead := head.Next
    for node != nil {
        temp := node.Next
        node.Next = prev
        node = temp

        if node == nil || node.Next == nil {
            prev.Next = node
            break
        }
        prev.Next = node.Next
        prev = node
        node = node.Next
    }

    return newHead
}
