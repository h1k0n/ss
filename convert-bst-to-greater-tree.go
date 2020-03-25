package main

import (
	"fmt"
)

func main() {
	n1 := &TreeNode{
		Val: 5,
	}
	n2 := &TreeNode{
		Val: 2,
	}
	n3 := &TreeNode{
		Val: 13,
	}
	n1.Left = n2
	n1.Right = n3
	fmt.Println("Hello, playground")
	x := convertBST(n1)
	inOrderX(x)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	s := sum(root)
	y := new(int)
	inOrder(root, y, s)
	return root
}

func sum(root *TreeNode) int {
	p := root
	if p != nil {
		return p.Val + sum(p.Left) + sum(p.Right)
	}
	return 0
}

func inOrder(root *TreeNode, pre *int, sum int) {
	p := root
	if p != nil {

		inOrder(p.Left, pre, sum)
		x := p.Val
		p.Val = sum - *pre
		*pre += x

		inOrder(p.Right, pre, sum)
		//return x
	}
	//return 0

}

func inOrderX(root *TreeNode) {
	p := root
	if p != nil {
		inOrderX(p.Left)
		fmt.Println(p.Val, "_")
		inOrderX(p.Right)

	}
}
