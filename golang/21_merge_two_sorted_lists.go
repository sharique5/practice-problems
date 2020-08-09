/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h3 := &ListNode{
		Next: nil,
		Val:  -999999,
	}
	t3 := &ListNode{
		Next: nil,
		Val:  -999999,
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp := &ListNode{
				Next: nil,
				Val:  l1.Val,
			}
			if t3.Val != -999999 {
				t3.Next = temp
				t3 = temp
			} else {
				h3 = temp
				t3 = temp
			}
			l1 = l1.Next
		} else {
			temp := &ListNode{
				Next: nil,
				Val:  l2.Val,
			}
			if t3.Val != -999999 {
				t3.Next = temp
				t3 = temp
			} else {
				h3 = temp
				t3 = temp
			}
			l2 = l2.Next
		}
	}
	for l1 != nil {
		temp := &ListNode{
			Next: nil,
			Val:  l1.Val,
		}
		if t3.Val != -999999 {
			t3.Next = temp
			t3 = temp
		} else {
			h3 = temp
			t3 = temp
		}
		l1 = l1.Next
	}
	for l2 != nil {
		temp := &ListNode{
			Next: nil,
			Val:  l2.Val,
		}
		if t3.Val != -999999 {
			t3.Next = temp
			t3 = temp
		} else {
			h3 = temp
			t3 = temp
		}
		l2 = l2.Next
	}
	if h3.Val == -999999 {
		return nil
	}
	return h3
}