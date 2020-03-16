package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", sortArrayByParity([]int{3}))
}

func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for {
		for ; i < j; i++ {
			if A[i]%2 == 1 {

				break
			}
		}
		for ; j > i; j-- {
			if A[j]%2 == 0 {
				break
			}
		}
		A[i], A[j] = A[j], A[i]
		if j <= i {
			break
		}
	}
	return A
}
