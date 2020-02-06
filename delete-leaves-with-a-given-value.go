package main

import (
	"fmt"
)

func main() {
	x1 := &TreeNode{
		Val: 1,
	}
	x2 := &TreeNode{
		Val: 3,
	}
	x3 := &TreeNode{
		Val: 3,
	}
	x4 := &TreeNode{
		Val: 3,
	}
	x5 := &TreeNode{
		Val: 2,
	}

	x1.Left = x2
	x1.Right = x3
	x2.Left = x4
	x2.Right = x5

	preOrder(x1)
	travel(x1)
	fmt.Println("Hello, playground")
	newTree := removeLeafNodes(x1, 3)
	preOrder(newTree)
}

func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val)
		fmt.Print("(")
		preOrder(root.Left)
		fmt.Print(",")
		preOrder(root.Right)
		fmt.Print(")")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(root *TreeNode) {
	iter := root

	stack := make([]*TreeNode, 3050)
	top := -1

	var p *TreeNode
	var flag int

	if iter != nil {
		for {
			for iter != nil {
				top++
				stack[top] = iter
				iter = iter.Left
			}
			p = nil
			flag = 1
			for top > -1 && flag == 1 {
				iter = stack[top]
				if iter.Right == p {
					fmt.Println(iter.Val, iter.Left == nil && iter.Right == nil && iter.Val == 3)
					if iter.Left == nil && iter.Right == nil && iter.Val == 3 {
						if stack[top-1].Left == iter {
							stack[top-1].Left = nil
						} else {
							stack[top-1].Right = nil
						}
					}
					top--
					p = iter
				} else {
					iter = iter.Right
					flag = 0
				}

			}
			if top <= -1 {
				break
			}
		}
	}

}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	iter := root

	stack := make([]*TreeNode, 3050)
	top := -1

	var p *TreeNode
	var flag int

	if iter != nil {
		for {
			for iter != nil {
				top++
				stack[top] = iter
				iter = iter.Left
			}
			p = nil
			flag = 1
			for top > -1 && flag == 1 {
				iter = stack[top]
				if iter.Right == p {
					if iter.Left == nil && iter.Right == nil && iter.Val == target {
						if top == 0 {
							return nil
						}
						if stack[top-1].Left == iter {
							stack[top-1].Left = nil
						} else {
							stack[top-1].Right = nil
						}
					}
					top--
					p = iter
				} else {
					iter = iter.Right
					flag = 0
				}
			}
			if top <= -1 {
				break
			}
		}
	}
	return iter
}
