package main

import (
	"fmt"
)

func main() {
	t1 := &Node{
		Val: 1,
	}
	t2 := &Node{
		Val: 2,
	}
	t3 := &Node{
		Val: 3,
	}
	t4 := &Node{
		Val: 4,
	}
	t5 := &Node{
		Val: 5,
	}
	t1.Children = []*Node{t2, t3, t4, t5}
	fmt.Println("Hello, playground", preorder(t1))
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	st := make([]*Node, 100)
	top := -1
	p := root
	var ptr *Node
	if p != nil {
		top++
		st[top] = p
		for top > -1 {
			ptr = st[top]
			top--
			res = append(res, ptr.Val)
			for i := len(ptr.Children) - 1; i >= 0; i-- {
				top++
				if len(st) == top {
					st = append(st, ptr.Children[i])
				} else {
					st[top] = ptr.Children[i]
				}
			}
		}
	}
	return res
}
