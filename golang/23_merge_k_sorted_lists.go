package main

import "sort"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var temp []int
	for i := 0; i < len(lists); i++ {
		t := lists[i]
		for t != nil {
			temp = append(temp, t.Val)
			t = t.Next
		}
	}
	if len(temp) == 0 {
		return nil
	}
	sort.Ints(temp)
	tempNode := &ListNode{
		Val:  temp[0],
		Next: nil,
	}
	head := tempNode
	for i := 1; i < len(temp); i++ {
		tempNode.Next = &ListNode{
			Val:  temp[i],
			Next: nil,
		}
		tempNode = tempNode.Next
	}
	return head
}
