package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{1, 0, 0, 0},
		[]int{1, 0, 0, 0},
	}
	fmt.Println("Hello, playground", kWeakestRows(input,2))
}

type myarr [][2]int

func (arr myarr) Len() int {
	return len(arr)
}

func (arr myarr) Less(i, j int) bool {
	if arr[i][1] < arr[j][1] {
		return true
	} else if arr[i][1] == arr[j][1] && arr[i][0] < arr[j][0] {
		return true
	}
	return false
}

func (arr myarr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func kWeakestRows(mat [][]int, k int) []int {
	arr := make(myarr, len(mat))
	for i := range mat {
		arr[i][0] = i
		for _, v := range mat[i] {
			if v == 1 {
				arr[i][1]++
			}
		}
	}
	//fmt.Println(arr)
	sort.Sort(arr)
	//fmt.Println(arr)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][0]
	}
	return res
}
