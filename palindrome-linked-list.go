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
		Val: 3,
	}
	a5 := &ListNode{
		Val: 2,
	}
	a6 := &ListNode{
		Val: 1,
	}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6
	fmt.Println("Hello, playground", isPalindrome(a1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	mid, prev := middle(head)
	p, q := mid, prev
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true

}

func middle(head *ListNode) (*ListNode, *ListNode) {
	fast, slow := head, head

	var prev, tmp *ListNode
	for fast != nil && fast.Next != nil {
		//fmt.Println(fast.Val)
		fast = fast.Next.Next

		tmp = slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp

	}
	if fast != nil && tmp != nil {
		tmp = tmp.Next
	}
	return tmp, prev
}
