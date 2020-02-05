package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", convert("PAYPALISHIRI", 3))
}

type Node struct {
	Val  byte
	Next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%c %p", n.Val, n.Next)
}
func newBox(row int) []Node {
	box := make([]Node, 2*row-2)
	for i := 1; i <= row-2; i++ {
		box[row-1-i].Next = &box[row-1+i]
	}
	return box
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arrLen := 2*numRows - 2
	var boxLength int
	if len(s)%arrLen == 0 {
		boxLength = len(s) / arrLen
	} else {
		boxLength = len(s)/arrLen + 1
	}

	boxSlice := make([][]Node, boxLength)
	for i, k := range []byte(s) {
		fmt.Printf("i,k %d %c\n", i, k)
		if i%arrLen == 0 {
			boxSlice[i/arrLen] = newBox(numRows)
		}
		boxSlice[i/arrLen][i%arrLen].Val = k
	}
	fmt.Println(boxSlice[0])
	fmt.Println(boxSlice[1])
	//fmt.Println(boxSlice[6])
	for i := 0; i < boxLength-1; i++ {
		boxSlice[i][0].Next = &boxSlice[i+1][0]
		for j := 0; j <= numRows-2; j++ {
			boxSlice[i][numRows-1+j].Next = &boxSlice[i+1][numRows-1-j]
		}
	}
	fmt.Println(boxSlice[0])
	fmt.Println(boxSlice[1])
	//fmt.Println(boxSlice[2])
	k := 0
	buf := make([]byte, len([]byte(s)))
	for i := 0; i < numRows; i++ {
		p := &boxSlice[0][i]
		for p != nil {
			if p.Val != 0 {

				fmt.Printf("%d %c\n", i, p.Val)
				buf[k] = p.Val
				k++
			}
			p = p.Next

		}
	}
	return string(buf)
}
