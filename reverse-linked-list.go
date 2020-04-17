package main

import (
	"fmt"
)

func main() {
	a := &ListNode{
		Val: 1,
	}
	b := &ListNode{
		Val: 2,
	}
	c := &ListNode{
		Val: 3,
	}
	d := &ListNode{
		Val: 4,
	}
	e := &ListNode{
		Val: 5,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	fmt.Println("Hello, playground", reverseList(a))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	p := head
	var prev, tmp *ListNode
	for p != nil {
		tmp = p
		p = tmp.Next
		tmp.Next = prev
		prev = tmp
	}
	//p = prev
	//for p != nil {
	//	fmt.Println(p.Val)
	//	p = p.Next
	//}
	return prev
}
