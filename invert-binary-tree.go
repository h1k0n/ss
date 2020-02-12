package main

import (
	"fmt"
)

func main() {
	x1 := &TreeNode{
		Val: 4,
	}
	x2 := &TreeNode{
		Val: 2,
	}
	x3 := &TreeNode{
		Val: 7,
	}
	x4 := &TreeNode{
		Val: 1,
	}
	x5 := &TreeNode{
		Val: 3,
	}
	x6 := &TreeNode{
		Val: 6,
	}
	x7 := &TreeNode{
		Val: 9,
	}
	x1.Left = x2
	x1.Right = x3

	x2.Left = x4
	x2.Right = x5

	x3.Left = x6
	x3.Right = x7
	fmt.Println("Hello, playground")
	r := invertTree(x1)
	inOrder(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode) {
	if root != nil {

		inOrder(root.Left)
		fmt.Println(root.Val)
		inOrder(root.Right)
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		if root.Left != nil || root.Right != nil {
			root.Left, root.Right = root.Right, root.Left
		}
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
