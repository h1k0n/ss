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
		a2Left := &TreeNode{
		Val: -1,
	}
	a2LeftLeft := &TreeNode{
		Val: -2,
	}
	a3 := &TreeNode{
		Val: 33,
	}
	a1.Left = a2
	a2.Left=a2Left
	a2Left.Left=a2LeftLeft
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
	a13 := &TreeNode{
		Val: 24,
	}
	a10.Left = a13
	//a10.Left = nil

	a10.Right = a12
	//t:=a12
	//t=nil

	fmt.Println("Hello, playground", diameterOfBinaryTree(a1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, x := Xabc(root)
	//fmt.Println("h", y)
	return x
}
func Xabc(node *TreeNode) (int, int) {
	if node != nil {
		leftHeight, diameterLeft := Xabc(node.Left)
		rightHeight, diameterRight := Xabc(node.Right)
		

		max := leftHeight + rightHeight

		if diameterLeft > max {
			max = diameterLeft
		}
		if diameterRight > max {
			max = diameterRight
		}

		maxHeight := leftHeight
		if rightHeight > maxHeight {
			maxHeight = rightHeight
		}

		return maxHeight + 1, max
	}
	return 0, 0
}
