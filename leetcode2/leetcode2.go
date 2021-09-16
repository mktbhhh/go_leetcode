package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			p = head
		} else {
			p.Next = &ListNode{Val: sum}
			p = p.Next
		}
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head
}

func main() {
	l1 := ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	l := addTwoNumbers(&l1, &l2)

	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
}
