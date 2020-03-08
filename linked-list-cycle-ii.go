package main

import (
	"fmt"
)

func main() {

	n1 := &ListNode{
		Val: 3,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 0,
	}
	n4 := &ListNode{
		Val: -4,
	}
	n5 := &ListNode{
		Val: 5,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n4.Next = n4
	//n1 = nil

	fmt.Println("Hello, playground")
	x := detectCycle(n1)
	if x != nil {
		fmt.Println(x.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	var q *ListNode
	i := 1
	j := 0
	star := 0
	for p != nil {
		if p == q {
			//fmt.Println(i, p.Val, 1<<j)
			break
		}

		if i == 1<<j {
			//fmt.Println("-", i, p.Val)
			star = i
			q = p
			j++
		}
		p = p.Next
		i++

	}
	//fmt.Println("--", i-star)
	if p != nil && i-star > 0 {
		for r := head; r != nil; r = r.Next {
			s := r
			for k := 0; k < i-star; k++ {
				s = s.Next
			}
			//fmt.Println("x", r.Val, s.Val)
			if s == r {
				//fmt.Println(".", r.Val, s.Val)
				return s
			}

		}
	}
	return nil
}
