package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", hammingDistance(1<<31-1,0))
}

func hammingDistance(x int, y int) int {
	a := x ^ y
	count := 0

	for i := 0; i < 64; i++ {
		if (1<<i)&a > 0 {
			count++
		}
	}
	return count
}
