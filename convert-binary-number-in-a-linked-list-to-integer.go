package main

import (
	"fmt"
)

func main() {
	a1 := &ListNode{
		Val: 1,
	}
	a2 := &ListNode{
		Val: 0,
	}
	a3 := &ListNode{
		Val: 1,
	}
	a1.Next = a2
	a2.Next = a3
	fmt.Println("Hello, playground", getDecimalValue(nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	st := make([]int, 31)
	top := -1

	num := 0
	p := head
	for p != nil {
		top++
		st[top] = p.Val
		p = p.Next
	}
	for i := top; i >= 0; i-- {
	        //fmt.Println(i,st[i])
		num += (st[i] << (top - i))
	}
	return num
}
