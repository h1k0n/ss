package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	x1 := &ListNode{
		Val: 1,
	}
	x2 := &ListNode{
		Val: 2,
	}
	x3 := &ListNode{
		Val: 4,
	}
	x1.Next = x2
	x2.Next=x3
	
	y1 := &ListNode{
		Val: 1,
	}
	y2 := &ListNode{
		Val: 3,
	}
	y3 := &ListNode{
		Val: 4,
	}
	
	y1.Next = y2
	y2.Next = y3

	z := mergeTwoLists(x1, y1)
	fmt.Println("Hello, playground", z.Val, z.Next.Val, z.Next.Next.Val, z.Next.Next.Next.Val, z.Next.Next.Next.Next.Val,z.Next.Next.Next.Next.Next.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode
	var p1, p2, p1pre, popPointer *ListNode
	if length(l1) >= length(l2) {
		p1, p2 = l1, l2
	} else {
		p1, p2 = l2, l1
	}
	newList, p1pre = p1, p1
	if p1 != nil && p2 != nil && p2.Val <= p1.Val {
		newList, p1pre = p2, p1
	}
	for p1 != nil && p2 != nil {
		if p2.Val <= p1.Val {
			popPointer = p2.Next
			p2.Next = p1
			if p1pre != p1 {
				p1pre.Next = p2
			}
			p1pre=p2
			p2 = popPointer
		} else {
			p1pre = p1
			p1 = p1.Next
		}
	}
	if p1 == nil && p2 != nil {
		p1pre.Next = p2
	}
	return newList
}

func length(list *ListNode) int {
	var count int
	p := list
	for p != nil {
		count++
		p = p.Next
	}
	return count
}
