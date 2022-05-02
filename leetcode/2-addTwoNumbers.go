package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 9,
	}

	l2 := &ListNode{
		Val: 9,
	}

	println(addtwonumbers(l1, l2).Next.Val)
}

func addtwonumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

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

		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return dummy.Next
}
