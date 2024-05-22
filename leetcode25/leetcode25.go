package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curTail := &ListNode{}
	var reversedHead *ListNode
	node := head
	for node != nil {
		var curHead *ListNode
		// var p *ListNode
		formerHead := node
		i := 0
		var tail *ListNode
		for i = 0; i < k && node != nil; i++ {
			if curHead == nil {
				curHead = &ListNode{Val: node.Val}
				tail = curHead
				node = node.Next
				continue
			}
			p := &ListNode{Val: node.Val}
			p.Next = curHead
			curHead = p
			node = node.Next
		}

		if reversedHead == nil {
			reversedHead = curHead
		}
		if i == k {
			curTail.Next = curHead
			curTail = tail
		} else {
			curTail.Next = formerHead
			return reversedHead
		}
	}

	return reversedHead
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil}}}}}
	reversedHead := reverseKGroup(head, 3)
	fmt.Println(reversedHead)
}
