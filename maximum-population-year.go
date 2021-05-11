package main

import (
	"fmt"
)

func main() {
	slice := [][]int{
		//{1950, 1961}, {1960, 1971}, {1970, 1981},
		{1993, 1999}, {2000, 2010},
	}
	fmt.Println("Hello, playground", maximumPopulation(slice))
}

func maximumPopulation(logs [][]int) int {
	x := make([]int, 2050-1950+1)
	for i := 0; i < len(logs); i++ {
		for j := logs[i][0]; j < logs[i][1]; j++ {
			x[j-1950]++
		}
	}
	index, max := 0, 0
	for i, v := range x {
		if v > max {
			max = v
			index = i
		}
	}
	return index + 1950
}

