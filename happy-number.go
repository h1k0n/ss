package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", isHappy(21))
}

func isHappy(n int) bool {
	m := make(map[int]int)
	var num int
	for {
		if m[n] != 0 {
			break
		}
		//fmt.Println(n, m[n])
		num = n
		sum := 0
		for num != 0 {
			digit := num % 10
			sum += digit * digit
			num = num / 10

		}

		m[n] = sum
		//fmt.Println(n, m[n])
		n = sum
		if sum == 1 {
			return true
		}
	}
	return false
}
