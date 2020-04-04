package main

import (
	"fmt"
)

func main() {
	a1 := &ListNode{
		Val: 9,
	}
	a2 := &ListNode{
		Val: 9,
	}
	a3 := &ListNode{
		Val: 9,
	}
	a4 := &ListNode{
		Val: 9,
	}
	a5 := &ListNode{
		Val: 9,
	}
	b1 := &ListNode{
		Val: 5,
	}
	b2 := &ListNode{
		Val: 6,
	}
	b3 := &ListNode{
		Val: 4,
	}
	c1 := &ListNode{
		Val: 1,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	b1.Next = b2
	b2.Next = b3
	fmt.Println("Hello, playground", addTwoNumbers(c1, a1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	head := l1
	inc := 0
	sum := 0
	var tmpP *ListNode
	for p != nil && q != nil {
		sum = p.Val + q.Val + inc
		p.Val = sum % 10
		inc = sum / 10

		tmpP = p
		p = p.Next
		q = q.Next
	}
	if p==nil&&q != nil {
	        tmpP.Next=q
		p = q
		
	}
	for p != nil {
		sum = p.Val + inc
		p.Val = sum % 10
		inc = sum / 10
		tmpP = p
		p = p.Next
	}

	if inc != 0 {
		tmpP.Next = &ListNode{
			Val: inc,
		}

	}
	//p = head
	//for p != nil {
	//	fmt.Println(p.Val)
	//	p = p.Next
	//}
	return head

}
