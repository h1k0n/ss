package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	root := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
	fmt.Println("Hello, playground", root == nil)
	inOrder(root)
	preOrder(root)
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
func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		preOrder(root.Left)

		preOrder(root.Right)
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	root = construct(root, nums, 0)
	return root
}
func construct(root *TreeNode, nums []int, loc int) *TreeNode {

	length := len(nums)

	if length == 1 {
		//fmt.Println("--", nums[0])
		node := &TreeNode{
			Val: nums[length/2],
		}
		if loc == -1 {
			root.Left = node
		} else if loc == 1 {
			root.Right = node
		} else {
			root = node
		}
	} else if length > 1 {
		//fmt.Println(nums[length/2])
		node := &TreeNode{
			Val: nums[length/2],
		}

		if loc == -1 {
			root.Left = node
		} else if loc == 1 {
			root.Right = node
		} else {
			root = node
		}
		construct(node, nums[0:length/2], -1)
		if length/2+1 < length {
			construct(node, nums[length/2+1:length], 1)
		}
	}
	return root

}
