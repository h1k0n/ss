package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	a := minStack.GetMin() //--> Returns -3.
	fmt.Println("Hello, playground", a)
	minStack.Pop()
	b := minStack.Top() //--> Returns 0.
	fmt.Println("Hello, playground", b)
	c := minStack.GetMin()
	fmt.Println("Hello, playground", c)
}

type MinStack struct {
	st         [][2]int
	top        int
	minPointer int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{
		st:  make([][2]int, 0),
		top: -1,
	}
	return m
}

func (this *MinStack) Push(x int) {
	this.top++

	if this.top == 0 {
		this.minPointer = 0
		t := [2]int{x, this.minPointer}
		if this.top >= len(this.st) {
			this.st = append(this.st, t)
		} else {
			this.st[this.top] = t
		}
		return
	}
	t := [2]int{x, this.minPointer}
	if this.st[this.minPointer][0] > x {
		this.minPointer = this.top
	}
	if this.top >= len(this.st) {
		this.st = append(this.st, t)
	} else {
		this.st[this.top] = t
	}

}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
	}
	if this.minPointer == this.top {
		this.minPointer = this.st[this.minPointer][1]

	}
	this.top--
}

func (this *MinStack) Top() int {
	return this.st[this.top][0]
}

func (this *MinStack) GetMin() int {
	return this.st[this.minPointer][0]
}
