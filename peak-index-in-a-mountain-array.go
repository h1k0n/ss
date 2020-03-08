package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", peakIndexInMountainArray([]int{18, 100, 99, 98, 90, 89, 87, 86, 71}))
}

func peakIndexInMountainArray(A []int) int {
	length := len(A)

	left := 0
	right := length - 1
	m := (length - 1) / 2
Top:
	fmt.Println(m, A[m])
	if A[m+1] > A[m] {
		left = m
		m = m + (right-m)/2
		//right = m
		goto Top
	} else if A[m] < A[m-1] {
		right = m
		m = m - (m-left)/2
		//left = m
		goto Top
	}
	return m

}
