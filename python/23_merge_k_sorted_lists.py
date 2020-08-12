# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        temp = []
        for i in range(0, len(lists)):
            tempNode = lists[i]
            while tempNode != None:
                temp.append(tempNode.val)
                tempNode = tempNode.next
        temp = sorted(temp)
        if len(temp) == 0:
            return None
        tempNode = ListNode(temp[0])
        head = tempNode
        for i in range(1, len(temp)):
            tempNode.next = ListNode(temp[i])
            tempNode = tempNode.next
        return head