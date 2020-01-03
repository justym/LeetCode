package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{0, nil}
	var dummy = head
	var p, q = l1.Next, l2.Next
	var x, y int
	var sum, carry int

	for p == nil || q == nil {
		x = p.Val
		y = q.Val
		sum = x + y + carry
		carry = sum / 10
		dummy.Next = &ListNode{sum % 10, nil}
		p = p.Next
		q = q.Next
		dummy = dummy.Next
	}

	return head
}

func main() {
	var l1, l2 *ListNode
	fmt.Scanln(l1, l2)
	addTwoNumbers(l1, l2)
}
