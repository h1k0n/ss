package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	}
	fmt.Println("Hello, playground", shiftGrid(input, 1))
}

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	k = k % (m * n)


	for index := 0; index < m*n; index++ {

		res[k/n][k%n] = grid[index/n][index%n]
		//fmt.Println(k/n,k%n,index/n,index%n)
	
		k++
		if k == n*m {
			k = 0
		}

	}
	return res
}
