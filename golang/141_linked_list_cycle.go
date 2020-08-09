package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fastPtr, slowPtr := head, head
	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next

		if fastPtr == slowPtr {
			return true
		}
	}
	return false
}
