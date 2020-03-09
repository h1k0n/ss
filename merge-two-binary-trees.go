package main

import (
	"fmt"
)

func main() {
	a1 := &TreeNode{
		Val: 1,
	}
	a2 := &TreeNode{
		Val: 3,
	}
	a3 := &TreeNode{
		Val: 2,
	}
	a4 := &TreeNode{
		Val: 5,
	}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4

	b1 := &TreeNode{
		Val: 2,
	}
	b2 := &TreeNode{
		Val: 1,
	}
	b3 := &TreeNode{
		Val: 3,
	}
	b4 := &TreeNode{
		Val: 4,
	}
	b5 := &TreeNode{
		Val: 7,
	}
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b3.Right = b5
	fmt.Println("Hello, playground")
	res := mergeTrees(a1, b1)
	preOrder(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	res := preProcess(t1, t2)
	return res

}

func preProcess(t1, t2 *TreeNode) *TreeNode {
	p := t1
	q := t2
	if p == nil || q == nil {
		if p != nil {
			return p
		}
		return q
	}
	if p != nil && q != nil {
		q.Val += p.Val
		if p.Left == nil || q.Left == nil {
			if p.Left != nil && q.Left == nil {
				q.Left = p.Left
			}
		} else {
			preProcess(p.Left, q.Left)
		}

		if p.Right == nil || q.Right == nil {
			if p.Right != nil && q.Right == nil {
				q.Right = p.Right
			}
		} else {
			preProcess(p.Right, q.Right)
		}
	}
	return q
}
func preOrder(x *TreeNode) {
	if x != nil {
		fmt.Println("x", x.Val)
		preOrder(x.Left)
		preOrder(x.Right)
	}
}
