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
			Val: 0,
		}
		a3 := &TreeNode{
			Val: 2,
		}
		a1.Left = a2
		a1.Right = a3
	*/
	/*
		a1 := &TreeNode{
			Val: 3,
		}
		a2 := &TreeNode{
			Val: 0,
		}
		a3 := &TreeNode{
			Val: 6,
		}
		a4 := &TreeNode{
			Val: 2,
		}
		a5 := &TreeNode{
			Val: 1,
		}
		a6 := &TreeNode{
			Val: 4,
		}
		a7 := &TreeNode{
			Val: 5,
		}
		a1.Left = a2
		a1.Right = a3
		a2.Right = a4
		a4.Left = a5

		a3.Left = a6
		a6.Right = a7
	*/

	a1 := &TreeNode{
		Val: 5,
	}
	a2 := &TreeNode{
		Val: 3,
	}
	a3 := &TreeNode{
		Val: 6,
	}
	a4 := &TreeNode{
		Val: 1,
	}
	a5 := &TreeNode{
		Val: 4,
	}
	a6 := &TreeNode{
		Val: 2,
	}
	a1.Left = a2
	a1.Right = a3
	a2.Left = a4
	a2.Right = a5
	a4.Right = a6

	//a1 := construct([]int{2, 0, 33, -100, 1, 25, 40, -100, -100, 11, 31, 34, 45, 10, 18, 29, 32, -100, 36, 43, 46, 4, -100, 12, 24, 26, 30, -100, -100, 35, 39, 42, 44, -100, 48, 3, 9, -100, 14, 22, -100, -100, 27, -100, -100, -100, -100, 38, -100, 41, -100, -100, -100, 47, 49, -100, -100, 5, -100, 13, 15, 21, 23, -100, 28, 37, -100, -100, -100, -100, -100, -100, -100, -100, 8, -100, -100, -100, 17, 19, -100, -100, -100, -100, -100, -100, -100, 7, -100, 16, -100, -100, 20, 6})
	/*
		a1 := &TreeNode{
			Val: 2,
		}
		a2 := &TreeNode{
			Val: 0,
		}
		a3 := &TreeNode{
			Val: 33,
		}
		a1.Left = a2
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
		a10.Right = a12
	*/
	fmt.Println("Hello, playground")
	preOrder(a1)
	x := trimBST(a1, 1, 4)
	fmt.Println("Hello, preOrder")
	preOrder(x)
	fmt.Println("Hello, inOrder")
	inOrder(x)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(nums []int) *TreeNode {

	tmp := make([]*TreeNode, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] != -100 {
			tmp[i] = &TreeNode{
				Val: nums[i],
			}
		}
	}
	for i := 0; 2*i+2 < len(nums); i++ {
		if nums[i] != -100 {
			tmp[i].Left = tmp[2*i+1]
			tmp[i].Right = tmp[2*i+2]
		}
	}

	return tmp[0]
}
func preOrder(t *TreeNode) {
	if t != nil {
		fmt.Println(".", t.Val)
		preOrder(t.Left)
		preOrder(t.Right)
	}
}
func inOrder(t *TreeNode) {
	if t != nil {

		inOrder(t.Left)
		fmt.Println("x", t.Val)
		inOrder(t.Right)

	}
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	p := root
	t := -1
	var g, f *TreeNode
	if L < p.Val {
		for p != nil {
			if p.Val == L {
				p.Left = nil
				break
			} else if L > p.Val {
				f = p
				p = p.Right
				t--
			} else {
				g = p
				p = p.Left
				t = 1
			}
			if t == 0 {
				g.Left = f.Right
				p = g
			}
		}
		p = root
		if R >= p.Val {
			t = -1
			for p != nil {
				if p.Val == R {
					p.Right = nil
					break
				} else if R > p.Val {
					g = p
					p = p.Right
					t = 1
				} else {
					f = p
					p = p.Left
					t--
				}
				if t == 0 {
					g.Right = f.Left
					p = g
				}
			}
			return root
		} else {
			//处理 左边的右边界
			t = -1
			p = root
			res := root
			found := false
			for p != nil {
				if !found && p.Val <= R {
					res = p
					found = true

				}
				if p.Val == R {
					p.Right = nil
					break
				} else if R > p.Val {
					g = p
					p = p.Right
					t = 1
				} else {
					f = p
					p = p.Left
					t--
				}
				if t == 0 {
					g.Right = f.Left
					p = g
				}
			}
			return res
		}
	}
	t = -1
	p = root
	for p != nil {
		if p.Val == R {
			p.Right = nil
			break
		} else if R > p.Val {
			g = p
			p = p.Right
			t = 1
		} else {
			f = p
			p = p.Left
			t--
		}
		if t == 0 {
			g.Right = f.Left
			p = g
		}
	}
	t = -1
	p = root
	res := root
	found := false
	for p != nil {
		if !found && p.Val >= L {
			res = p
			found = true
		}
		if p.Val == L {
			p.Left = nil
			break
		} else if L > p.Val {
			f = p
			p = p.Right
			t--
		} else {
			g = p
			p = p.Left
			t = 1
		}
		if t == 0 {

			g.Left = f.Right
			p = g
		}
	}
	return res

}
