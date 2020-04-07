package main

import (
	"fmt"
)

func main() {
	list := buildLinkList([]int{6, 1, 2, 6, 3, 4, 5, 6})
	removeElements(list, 6)
	fmt.Println("Hello, playground")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildLinkList(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for i := 0; i < len(nums); i++ {
		p.Next = &ListNode{
			Val: nums[i],
		}
		p = p.Next
	}
	return head.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	fakeHead := &ListNode{}
	fakeHead.Next = head
	p := fakeHead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	//p = fakeHead.Next
	//for p != nil {
	//	fmt.Println(p.Val)
	//	p = p.Next
	//}
	return fakeHead.Next
}
