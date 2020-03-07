package main

import (
	"fmt"
)

func main() {
	n1 := &TreeNode{
		Val: 3,
	}
	n2 := &TreeNode{
		Val: 9,
	}
	n3 := &TreeNode{
		Val: 20,
	}
	n4 := &TreeNode{
		Val: 15,
	}
	n5 := &TreeNode{
		Val: 7,
	}

	n1.Left = n2
	n1.Right = n3

	n3.Left = n4
	n3.Right = n5
	//var nx *TreeNode
	fmt.Println("Hello, playground", levelOrder(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QNode struct {
	Level int
	Data  *TreeNode
	Next  *QNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	var front, rear *QNode
	front = &QNode{
		Level: 0,
		Data:  root,
	}
	rear = front
	flg := -1

	for front != nil {
		if front.Level != flg {
			t := make([]int, 0)
			res = append(res, t)

			flg = front.Level
		}
		res[flg] = append(res[flg], front.Data.Val)
		if front.Data.Left != nil {
			qnode := &QNode{
				Level: front.Level + 1,
				Data:  front.Data.Left,
			}
			rear.Next = qnode
			rear = qnode
		}
		if front.Data.Right != nil {
			qnode := &QNode{
				Level: front.Level + 1,
				Data:  front.Data.Right,
			}
			rear.Next = qnode
			rear = qnode
		}
		front = front.Next
	}
	return res

}
