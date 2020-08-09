# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        p = head
        c = head
        if head == None or head.next == None:
            return head
        n = head.next
        while n:
            c.next = n.next
            n.next = p
            p = n
            n = c.next
        return p