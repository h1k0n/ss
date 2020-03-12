package main

import (
	"fmt"
)

func main() {
	a1 := &TreeNode{
		Val: 2,
	}
	a2 := &TreeNode{
		Val: 0,
	}
	a3 := &TreeNode{
		Val: 33,
	}
	a1.Left = a2
	a1.Right = a3
	a4 := &TreeNode{
		Val: 25,
	}
	a5 := &TreeNode{
		Val: 40,
	}
	a3.Left = a4
	a3.Right = a5
	a6 := &TreeNode{
		Val: 11,
	}
	a7 := &TreeNode{
		Val: 31,
	}
	a4.Left = a6
	a4.Right = a7
	a8 := &TreeNode{
		Val: 29,
	}
	a9 := &TreeNode{
		Val: 32,
	}
	a7.Left = a8
	a7.Right = a9
	a10 := &TreeNode{
		Val: 26,
	}
	a11 := &TreeNode{
		Val: 30,
	}
	a8.Left = a10
	a8.Right = a11
	a12 := &TreeNode{
		Val: 27,
	}
	a10.Left = a12
	fmt.Println("Hello, playground", inorderTraversal(a4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	st := make([]*TreeNode, 1531)
	top := -1
	p := root
	for {
		for p != nil {
			top++
			st[top] = p
			p = p.Left
		}
		if top == -1 {
			break
		}
		for top > -1 {
			q := st[top]
			top--
			//fmt.Println(q.Val)
			res = append(res, q.Val)
			if q.Right != nil {
				p = q.Right
				break
			}

		}

	}
	return res
}
