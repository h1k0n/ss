package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("Hello, playground", lastStoneWeight([]int{2, 3, 3, 4, 4, 5, 5}))
}

func lastStoneWeight(stones []int) int {
	length := len(stones)
	if length == 1 {
		return stones[0]
	}
	for i := ((length - 1) - 1) / 2; i >= 0; i-- {
		sift(stones, i, length-1)
	}
	//fmt.Println(".", stones)
	loc := length - 1
	rock := [2]int{}
	x := 0
	for loc >= 0 {
		for i := 0; i < 2; i++ {
			if loc < 0 {
				return stones[0]
			}
			rock[i] = stones[0]
			stones[0], stones[loc] = stones[loc], stones[0]

			loc--
			sift(stones, 0, loc)

		}
		x = rock[0] - rock[1]
		//fmt.Println("x=", x, rock[0], rock[1], stones, loc)
		if x > 0 {
			loc += 1
			stones[loc] = x
			for i := (loc - 1) / 2; i >= 0; i-- {
				sift(stones, i, loc)
			}
			//fmt.Println("...", stones, loc)
		}
	}
	return x
}

func sift(stones []int, biggestLoc int, loc int) {
	//fmt.Println(biggestLoc, loc)
	for i := biggestLoc; 2*i+1 <= loc; {
		biggerIndex := 2*i + 1
		if 2*i+2 <= loc {
			if stones[2*i+1] < stones[2*i+2] {
				biggerIndex = 2*i + 2
			}
		}
		if stones[i] < stones[biggerIndex] {
			stones[i], stones[biggerIndex] = stones[biggerIndex], stones[i]
			i = biggerIndex
		} else {
			return
		}
	}
}
