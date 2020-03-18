package main

import (
	"fmt"
)

func main() {
	//ab := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	//ab := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	//ab := [][]int{{7, 8}, {1, 2}}
	ab:=[][]int{{2}}
	fmt.Println("Hello, playground", luckyNumbers(ab))

}

func luckyNumbers(matrix [][]int) []int {
	res := make([]int, 0)
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		minIndex := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] < matrix[i][minIndex] {
				minIndex = j
			}
		}
		//maxIndex:=minIndex
		k := 0
		for ; k < m; k++ {
			if matrix[k][minIndex] > matrix[i][minIndex] {
				break
			}
		}
		if k == m {
			res = append(res, matrix[i][minIndex])
		}

	}
	return res
}
