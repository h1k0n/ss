package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	matrix := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length-1; i++ {
		for j := i; i+j < length-1; j++ {
			//fmt.Println("-", i, j)
			matrix[j][length-1-i], matrix[length-1-i][length-1-j], matrix[length-1-j][i], matrix[i][j] = matrix[i][j], matrix[j][length-1-i], matrix[length-1-i][length-1-j], matrix[length-1-j][i]
		}
	}
}
