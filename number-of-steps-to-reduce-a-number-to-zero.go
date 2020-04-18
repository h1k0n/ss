package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",numberOfSteps(14))
}

func numberOfSteps(num int) int {
	count := 0
	for {
		if num%2 == 0 {
			num /= 2
			count++
		} else {
			num -= 1
			count++
		}
		if num == 0 {
			break
		}

	}
	return count
}
