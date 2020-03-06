package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", countBits(5))
}

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	if num == 1 {
		return []int{0, 1}
	}

	res := make([]int, num+1)
	res[0] = 0
	res[1] = 1
	n := 1
	step := 0
	for i := 2; i <= num; i++ {
		if 1<<n == i {
			res[i] = 1
			step = 1 << n
			n++
			continue
		}
		res[i] = res[i-step] + 1

	}
	return res
}
