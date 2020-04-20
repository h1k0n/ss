package main

import (
	"fmt"
)

func main() {
	a1 := &ListNode{
		Val: 1,
	}
	a2 := &ListNode{
		Val: 1,
	}
	a3 := &ListNode{
		Val: 2,
	}
	a1.Next = a2
	a2.Next = a3
	fmt.Println("Hello, playground", deleteDuplicates(a1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev := &ListNode{
		Val: head.Val - 1,
	}
	p := head
	for p != nil {
		if p.Val != prev.Val {
			prev = p

		} else {
			prev.Next = p.Next
		}
		p = p.Next
	}
	//p = head
	//for p != nil {
	//	fmt.Println(p.Val)
	//	p = p.Next
	//}
	return head

}
