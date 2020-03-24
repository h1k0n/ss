package main

import (
	"fmt"
)

func main() {
	
		node := &ListNode{
			Val: 1,
		}
		node2 := &ListNode{
			Val: 2,
		}
		node3 := &ListNode{
			Val: 3,
		}
		node4 := &ListNode{
			Val: 4,
		}
		node5 := &ListNode{
			Val: 5,
		}

		node.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
	

	fmt.Println("Hello, playground", removeNthFromEnd(node, 4))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	p := head
	i := 0
	for p != nil {
		i++
		p = p.Next
	}
	if n == i {
		return head.Next
	}

	p = head
	j := 0
	for p != nil && j < i-n-1 {
		j++
		p = p.Next
	}


	p.Next = p.Next.Next

	return head
}
