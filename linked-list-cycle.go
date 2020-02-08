package main

import (
	"fmt"
)

func main() {
	x1 := &ListNode{
		Val: 3,
	}
	x2 := &ListNode{
		Val: 0,
	}
	x3 := &ListNode{
		Val: 1,
	}
	x4 := &ListNode{
		Val: 4,
	}
	x1.Next = x2
	x2.Next = x3
	x3.Next = x4
	x4.Next = x2

	fmt.Println("Hello, playground", hasCycle(x1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var c uint64

	i := uint64(0)

	checkPoint := head
	p := head

	for p != nil {
		if c == 1<<i {
			checkPoint = p
			i++
		}
		c++
		p = p.Next
		if checkPoint == p {
			return true
		}
	}

	return false
}
