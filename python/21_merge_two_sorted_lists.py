# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        head1 = l1
        head2 = l2
        head_merged, tail_merged = None, None
        while head1 != None and head2 != None:
            if head1.val < head2.val:
                temp = ListNode(head1.val)
                if tail_merged != None:
                    tail_merged.next = temp
                    tail_merged = temp
                else:
                    head_merged = temp
                    tail_merged = temp
                head1 = head1.next
            else:
                temp = ListNode(head2.val)
                if tail_merged != None:
                    tail_merged.next = temp
                    tail_merged = temp
                else:
                    head_merged = temp
                    tail_merged = temp
                head2 = head2.next
        while head1:
            temp = ListNode(head1.val)
            if tail_merged != None:
                tail_merged.next = temp
                tail_merged = temp
            else:
                head_merged = temp
                tail_merged = temp
            head1 = head1.next
        while head2:
            temp = ListNode(head2.val)
            if tail_merged != None:
                tail_merged.next = temp
                tail_merged = temp
            else:
                head_merged = temp
                tail_merged = temp
            head2 = head2.next
        return head_merged