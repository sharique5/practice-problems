package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	list := []*ListNode{}
	tmp := head

	for tmp != nil {
		list = append(list, tmp)
		tmp = tmp.Next
	}

	h := head
	for i := len(list) - 1; i > (len(list)-1)/2; i-- {
		next := h.Next
		h.Next = list[i]
		h.Next.Next = next
		h = next
	}
	h.Next = nil
}
