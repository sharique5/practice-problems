package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{
		Val:  -1,
		Next: head,
	}

	fastPtr := temp
	slowPtr := temp

	for i := 0; i < n+1; i++ {
		fastPtr = fastPtr.Next
	}

	for fastPtr != nil {
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}
	slowPtr.Next = slowPtr.Next.Next
	return temp.Next
}
