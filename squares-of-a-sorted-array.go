package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(A []int) []int {
	length := len(A)
	if length == 0 {
		return []int{}
	}
	res := make([]int, length)
	if A[0] >= 0 {
		index := 0
		for i := 0; i < length; i++ {
			res[index] = A[i] * A[i]
			index++
		}
		return res
	}
	if A[length-1] < 0 {
		index := 0
		for i := length - 1; i >= 0; i-- {
			res[index] = A[i] * A[i]
			index++
		}
		return res
	}
	i := divide(A)
	j := i + 1

	for index := 0; j < length || i >= 0; index++ {
		if i < 0 && j < length {
			res[index] = A[j] * A[j]
			j++
			continue
		} else if j >= length && i >= 0 {
			res[index] = A[i] * A[i]
			i--
			continue
		}
		if -A[i] > A[j] {
			res[index] = A[j] * A[j]
			j++
		} else {
			res[index] = A[i] * A[i]
			i--
		}
	}

	return res
}

func divide(A []int) int {
	length := len(A)
	p, q := 0, length-1
	for p <= q {
		mid := (p + q) / 2
		if mid+1 < length && A[mid] < 0 && A[mid+1] >= 0 {
			return mid
		} else if A[mid] >= 0 {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	return -1
}
