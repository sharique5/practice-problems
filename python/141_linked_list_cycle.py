# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head == None or head.next == None:
            return False
        fast_ptr = head
        slow_ptr = head
        while fast_ptr != None and fast_ptr.next != None:
            fast_ptr = fast_ptr.next.next
            slow_ptr = slow_ptr.next
      
            if fast_ptr == slow_ptr:
                return True
        return False