package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", canPlaceFlowers([]int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flower := 0
	countZero := 0
	if flowerbed[0] == 0 {
		countZero = 1
	}
	length := len(flowerbed)
	i := 0
	for ; i < length; i++ {
		if flowerbed[i] == 1 {
			flower += (countZero - 1) / 2
			countZero = 0
			continue
		}
		countZero++
	}
	if i == length && flowerbed[i-1] == 0 {
		flower += countZero / 2
	}
	if flower >= n {
		return true
	}
	return false
}
