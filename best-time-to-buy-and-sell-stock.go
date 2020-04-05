package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	min := prices[0]
	maxProfit := 0
	tmp := 0
	for i := 0; i < length; i++ {
		if min > prices[i] {
			min = prices[i]
		}
		tmp = prices[i] - min
		if maxProfit < tmp {
			maxProfit = tmp
		}
	}
	return maxProfit
}
