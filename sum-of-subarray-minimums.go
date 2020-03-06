package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", sumSubarrayMins([]int{3, 1, 2, 4}))
}

func sumSubarrayMins(A []int) int {
	n := len(A)

	sum := A[0]
	modSum := sum % (1e9 + 7)
	for i := 1; i < n; i++ {
		sum = 0
		sum += A[i]
		for j := 1; j < i+1; j++ {
			if A[i-j] > A[i-j+1] {
				A[i-j] = A[i-j+1]
			}

			sum += A[i-j]
		}
		modSum += sum % (1e9 + 7)
	}
	return modSum % (1e9 + 7)

}

