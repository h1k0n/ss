package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		[]int{4, 3, 2, -1}, []int{3, 2, 1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3},
		//[]int{3, 2}, []int{1, 0},
		//[]int{1,-1}, []int{-1,-1},
		//[]int{1},
	}
	fmt.Println("Hello, playground", countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	var count int
	for i := 0; i < row; i++ {
		count += countNegative(grid[i], column)
	}
	return count
}

func countNegative(nums []int, length int) int {
	if nums[0] < 0 {
		return length
	}
	if length == 1 {
		return 0
	}

	p, q := 0, length-1
	var mid int
	for p <= q {
		mid = (p + q) / 2
		if mid+1 < length && nums[mid] >= 0 && nums[mid+1] < 0 {
			return length - 1 - mid
		} else if nums[mid] < 0 {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	return 0
}
