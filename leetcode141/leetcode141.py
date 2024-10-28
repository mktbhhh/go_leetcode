# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head == None:
            return False

        fast, slow = head, head
        if head.next != None and head.next.next != None:
            fast = fast.next.next
            slow = slow.next
        else:
            return False
        while fast != slow and fast.next and fast.next.next:
            fast = fast.next.next
            slow = slow.next

        if fast == slow and fast != None:
            return True
        else:
            return False
