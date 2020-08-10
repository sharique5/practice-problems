# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        temp = ListNode(-1)
        temp.next = head
        
        fast_ptr = temp
        slow_ptr = temp
        
        for i in range(n + 1):
            fast_ptr = fast_ptr.next
        
        while fast_ptr:
            fast_ptr = fast_ptr.next
            slow_ptr = slow_ptr.next
        
        slow_ptr.next = slow_ptr.next.next
        return temp.next