# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if head == None:
            return head
        temp = head.next
        doubleEnded = collections.deque([])
        while temp:
            doubleEnded.append(temp)
            temp = temp.next
        popRight = True
        temp = head
        while len(doubleEnded) > 0:
            if popRight:
                temp2 = doubleEnded.pop()
                temp.next = temp2
                temp = temp2
                popRight = False
            else:
                temp2 = doubleEnded.popleft()
                temp.next = temp2
                temp = temp2
                popRight = True
        temp.next = None