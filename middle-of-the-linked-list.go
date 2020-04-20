package main

import (
	"fmt"
)

func main() {
	a1 := &ListNode{
		Val: 1,
	}
	a2 := &ListNode{
		Val: 2,
	}
	a3 := &ListNode{
		Val: 3,
	}
	a4 := &ListNode{
		Val: 4,
	}
	a5 := &ListNode{
		Val: 5,
	}
	a6 := &ListNode{
		Val: 6,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	fmt.Println("Hello, playground", middleNode(a1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil {

		if fast.Next.Next == nil {
			return slow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}
