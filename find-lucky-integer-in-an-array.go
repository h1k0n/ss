package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", findLucky([]int{7, 7, 7, 7, 7, 7, 7}))
}

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	max := -1
	for k, v := range m {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
