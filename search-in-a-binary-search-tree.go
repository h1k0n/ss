package main

import (
	"fmt"
)

func main() {
	t1 := &TreeNode{
		Val: 4,
	}
	t2 := &TreeNode{
		Val: 2,
	}
	t3 := &TreeNode{
		Val: 7,
	}
	t4 := &TreeNode{
		Val: 1,
	}
	t5 := &TreeNode{
		Val: 3,
	}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t := searchBST(t1, 2)
	preorder(t)
	fmt.Println("Hello, playground")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	p := root
	if p != nil {
		if p.Val == val {
			return p
		}
		lptr := searchBST(p.Left, val)
		rptr := searchBST(p.Right, val) //todo:这里没有充分使用题目中BST的信息，导致同时搜索两侧
		if lptr != nil {
			return lptr
		}
		if rptr != nil {
			return rptr
		}
		return nil
	}
	return nil
}

func preorder(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)
		preorder(t.Left)
		preorder(t.Right)
	}
}
