package main

import (
	"fmt"
)

func main() {
	/*
		a1 := &TreeNode{
			Val: 1,
		}
		a2 := &TreeNode{
			Val: 2,
		}
		a3 := &TreeNode{
			Val: 3,
		}
		a1.Left = a2
		a1.Right = a3
	*/
	a1 := &TreeNode{
		Val: 4,
	}
	a2 := &TreeNode{
		Val: 9,
	}
	a3 := &TreeNode{
		Val: 0,
	}
	a1.Left = a2
	a1.Right = a3
	a4 := &TreeNode{
		Val: 5,
	}
	a5 := &TreeNode{
		Val: 1,
	}
	a2.Left = a4
	a2.Right = a5
	fmt.Println("Hello, playground", sumNumbers(a2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	total := new(int)
	sum := 0
	calculate(root, sum, total)
	return *total
}

func calculate(node *TreeNode, sum int, total *int) {
	if node != nil {
		sum *= 10
		sum += node.Val
	}
	if node.Left == nil && node.Right == nil {
		*total += sum
	}
	if node.Left != nil {
		calculate(node.Left, sum, total)
	}
	if node.Right != nil {
		calculate(node.Right, sum, total)
	}
}
