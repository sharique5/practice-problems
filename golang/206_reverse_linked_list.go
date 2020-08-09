package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n, c := head.Next, head
	for n != nil {
		c.Next = n.Next
		n.Next = head
		head = n
		n = c.Next
	}
	return head
}
