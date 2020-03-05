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

	fmt.Println("Hello, playground")
	x := swapPairs(a)
	printList(x)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	var q, r, tail *ListNode

	p := head
	for i := 0; p != nil; i++ {
		if i == 1 {

			head = p
			//fmt.Println(i,head.Val)
		}
		if i%2 == 1 {

			if i > 2 {
				tail.Next = p
			}
			r = p.Next
			p.Next = q
			q.Next = r
			tail = q
			p = r
		} else {
			q = p
		}

		if i%2 == 0 {
			p = p.Next
		}
	}
	return head
}

func printList(head *ListNode) {
	p := head
	for p != nil {

		fmt.Println(p.Val)
		p = p.Next
	}
}
