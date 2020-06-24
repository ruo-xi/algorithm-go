package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {
	dummyHead := ListNode{Val: 0}
	p, q, curr := l1, l2, dummyHead
	carry := 0
	for p != nil && q != nil {
		var x, y int
		if p == nil {
			x = 0
		} else {
			x = p.Val
		}
		if q == nil {
			y = 0
		} else {
			y = p.Val
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: carry % 10}
		curr = *curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
