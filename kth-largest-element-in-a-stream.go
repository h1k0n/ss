package main

import (
	"fmt"
)

func main() {

	k := 3
	nums := []int{4, 5, 8, 2}
	//k := 1
	//nums := []int{}
	obj := Constructor(k, nums)
	param_1 := obj.Add(3)
	fmt.Println("Hello, playground", param_1)
	param_1 = obj.Add(5)
	fmt.Println("Hello, playground", param_1)

	param_1 = obj.Add(10)
	fmt.Println("Hello, playground", param_1)
	param_1 = obj.Add(9)
	fmt.Println("Hello, playground", param_1)
	param_1 = obj.Add(4)
	fmt.Println("Hello, playground", param_1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type KthLargest struct {
	Head *ListNode
	Tail *ListNode
	Kth  *ListNode
	K    int
}

//只维护k个降序


func Constructor(k int, nums []int) KthLargest {
	if len(nums) == 0 {
		return KthLargest{
			K: k,
		}
	}
	head := &ListNode{
		Val: nums[0],
	}

	for i := 1; i < len(nums); i++ {
		p := head
		var latest *ListNode
		for p != nil {
			if nums[i] >= p.Val {
				node := &ListNode{
					Val: nums[i],
				}
				if latest == nil {
					node.Next = p
					head = node
					break
				} else {
					node.Next = p
					latest.Next = node
					break
				}

			}
			latest = p
			p = p.Next
		}
		if p == nil {
			latest.Next = &ListNode{
				Val: nums[i],
			}
		}

	}
	/*
		p := head
		for i := 0; p != nil; i++ {
			fmt.Println("test", p.Val)
			p = p.Next
		}
	*/
	p := head
	for i := 0; i < k-1; i++ {
		p = p.Next
	}
	kth := KthLargest{
		Head: head,
		Kth:  p,
		K:    k,
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.Head == nil {
		this.Head = &ListNode{
			Val: val,
		}
	}
	p := this.Head

	if this.Kth != nil && val < this.Kth.Val {
		p = this.Kth
	}

	var latest *ListNode
	for p != nil {
		if val >= p.Val {
			node := &ListNode{
				Val: val,
			}
			if latest == nil {
				node.Next = p
				this.Head = node
				break
			} else {
				node.Next = p
				latest.Next = node
				break
			}

		}
		latest = p
		p = p.Next
	}
	if p == nil {
		latest.Next = &ListNode{
			Val: val,
		}
	}
	q := this.Head
	for i := 0; i < this.K-1; i++ {
		q = q.Next
	}
	this.Kth = q

	return this.Kth.Val
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
