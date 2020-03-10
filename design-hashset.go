package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	//obj.Add(2)
	//obj.Remove(2)
	//param_3 := obj.Contains(2)
	//fmt.Println("Hello, playground", param_3)
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	param_3 := hashSet.Contains(1) // returns true
	fmt.Println("Hello, playground", param_3)
	param_3 = hashSet.Contains(3) // returns false (not found)
	fmt.Println("Hello, playground", param_3)
	hashSet.Add(2)
	param_3 = hashSet.Contains(2) // returns true
	fmt.Println("Hello, playground", param_3)
	hashSet.Remove(2)
	param_3 = hashSet.Contains(2) // returns false (already removed)
	fmt.Println("Hello, playground", param_3)
	hashSet.Add(3065)
	hashSet.Add(4593)
	param_3 = hashSet.Contains(3065)
	fmt.Println("Hello, playground", param_3)
	param_3 = hashSet.Contains(4593)
	fmt.Println("Hello, playground", param_3)
	hashSet.Remove(3065)
	hashSet.Add(3065)
	hashSet.Remove(3065)
	hashSet.Add(4593)
	hashSet.Remove(4593)
	hashSet.Add(4593)
	param_3 = hashSet.Contains(3065)
	fmt.Println("Hello, playground", param_3)
	param_3 = hashSet.Contains(4593)
	fmt.Println("Hello, playground", param_3)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyHashSet struct {
	Table []*ListNode
	Size  int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	size := 1531
	table := make([]*ListNode, size)
	for i := 0; i < len(table); i++ {
		table[i] = &ListNode{
			Val: -1,
		}
	}
	return MyHashSet{
		Table: table,
		Size:  size,
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	node := &ListNode{
		Val: key,
	}
	head := this.Table[key%this.Size]
	if head.Next == nil {
		head.Next = node
		return
	}
	p := head.Next
	node.Next = p
	head.Next = node

}

func (this *MyHashSet) Remove(key int) {
	head := this.Table[key%this.Size]
	p := head
	q := p.Next
	for q != nil {
		if q.Val == key {
			p.Next = q.Next
			break
		}
		p = q
		q = q.Next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	head := this.Table[key%this.Size]
	p := head.Next
	for p != nil {
		if p.Val == key {
			return true
		}
		p = p.Next
	}
	return false

}
