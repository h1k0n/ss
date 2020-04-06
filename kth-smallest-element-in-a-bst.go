package main

import (
	"fmt"
)

func main() {
	t1 := &TreeNode{
		Val: 5,
	}
	t2 := &TreeNode{
		Val: 3,
	}
	t3 := &TreeNode{
		Val: 6,
	}
	t4 := &TreeNode{
		Val: 2,
	}
	t5 := &TreeNode{
		Val: 4,
	}
	t6 := &TreeNode{
		Val: 1,
	}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t4.Left = t6
	fmt.Println("Hello, playground", kthSmallest(t1, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var i, res int

	inOrder(root, &res, &i, k)
	return res
}

func inOrder(p *TreeNode, res, i *int, k int) {

	if p != nil && *i < k {

		inOrder(p.Left, res, i, k)

		*i++
		if *i == k {
			*res = p.Val
			//fmt.Println(p.Val)
		}
		inOrder(p.Right, res, i, k)

	}
}
