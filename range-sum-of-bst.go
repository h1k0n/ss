package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{
		Val: 10,
	}
	t2 := &TreeNode{
		Val: 5,
	}
	t3 := &TreeNode{
		Val: 15,
	}
	t4 := &TreeNode{
		Val: 3,
	}
	t5 := &TreeNode{
		Val: 7,
	}
	t6 := &TreeNode{
		Val: 18,
	}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Right = t6
	fmt.Println("Hello, playground", rangeSumBST(t1, 7, 15))
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	nums := make([]int, 0)
	inOrder(root, &nums)
	res := 0
	flg := false
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == R {
			res += nums[i]
			break
		}
		if flg {
			res += nums[i]
		}
		if nums[i] == L {
			res += nums[i]
			flg = true
		}

	}
	return res
}

func inOrder(p *TreeNode, nums *[]int) {
	if p != nil {
		inOrder(p.Left, nums)
		*nums = append(*nums, p.Val)
		inOrder(p.Right, nums)
	}
}
