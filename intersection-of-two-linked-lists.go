package main

import (
	"fmt"
)

func main() {
	a1 := &ListNode{
		Val: 0,
	}
	a2 := &ListNode{
		Val: 9,
	}
	a3 := &ListNode{
		Val: 1,
	}
	a1.Next = a2
	a2.Next = a3

	b1 := &ListNode{
		Val: 3,
	}
	b2 := &ListNode{
		Val: 2,
	}
	//b3 := &ListNode{
	//	Val: 4,
	//}
	//a3.Next = b2
	//b2.Next = b3

	b1.Next = b2
	fmt.Println("Hello, playground")
	a := getIntersectionNode(a1, b1)
	if a!=nil{
	fmt.Println(a.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	lengthA := 0
	for p != nil {
		lengthA++
		p = p.Next
	}
	//fmt.Println("a", lengthA)

	p = headB
	lengthB := 0
	for p != nil {
		lengthB++
		p = p.Next
	}
	//fmt.Println("b", lengthB)

	length := lengthA - lengthB
	p = headA
	q := headB
	if lengthA < lengthB {
		length = lengthB - lengthA
		p = headB
		q = headA
	}
	for i := 0; i < length; i++ {
		p = p.Next
	}

	for p != nil && q != nil {
		if p == q {
			return p

		}

		p = p.Next
		q = q.Next
	}

	return nil
}
