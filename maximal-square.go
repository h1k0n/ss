package main

import (
	"fmt"
)

func main() {
	matrix := [][]byte{
		[]byte{'0', '1', '1', '0', '0', '1', '0', '1', '0', '1'},
		[]byte{'0', '0', '1', '0', '1', '0', '1', '0', '1', '0'},
		[]byte{'1', '0', '0', '0', '0', '1', '0', '1', '1', '0'},
		[]byte{'0', '1', '1', '1', '1', '1', '1', '0', '1', '0'},
		[]byte{'0', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0', '1', '1', '1', '1', '0'},
		[]byte{'0', '0', '0', '1', '1', '0', '0', '0', '1', '0'},
		[]byte{'1', '1', '0', '1', '1', '0', '0', '1', '1', '1'},
		[]byte{'0', '1', '0', '1', '1', '0', '1', '0', '1', '1'}}
	fmt.Println("Hello, playground", maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {

	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	columns := len(matrix[0])
	if columns == 0 {
		return 0
	}
	maxEdge := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				//fmt.Println("*", i, j)
				if maxEdge == 0 {
					maxEdge = 1
				}
				edgeLength := rows - 1 - i
				if edgeLength > columns-1-j {
					edgeLength = columns - 1 - j
				}
				k := edgeLength
				for ; k > 0; k-- {
					if detect(matrix, i, j, k) == false {
						continue
					} else {
						break
					}
				}
				if maxEdge < k+1 {
					//fmt.Println("*", i, j)
					maxEdge = k + 1
				}

			}
		}
	}
	return maxEdge * maxEdge
}

func detect(matrix [][]byte, i, j, k int) bool {
	for x := 1; x <= k; x++ {
		if matrix[i+x][j] == '0' || matrix[i][j+x] == '0' {
			return false
		}
	}
	for x := 1; x <= k; x++ {
		for y := 1; y <= k; y++ {
			if matrix[i+x][j+y] == '0' {
				return false
			}
		}
	}
	return true
}
