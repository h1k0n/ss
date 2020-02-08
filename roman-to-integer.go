package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var num int

	preNum := 5000
	for _, v := range []byte(s) {
		xx := tranfer(v)
		if preNum < xx {
			num = num - preNum + xx - preNum
		} else {
			num += xx
			preNum = xx
		}
	}
	return num
}

func tranfer(b byte) int {
	bricks := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	bricksVal := []int{1, 5, 10, 50, 100, 500, 1000}
	for i, v := range bricks {
		if v == b {
			return bricksVal[i]
		}
	}
	return 0
}
