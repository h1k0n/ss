package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", climbStairs(1))
}

func climbStairs(n int) int {
	a := make([]int, n+1)
	res := lower(a, n)

	return res
}

func lower(a []int, n int) int {
	if n < 3 {
		a[n] = n
		return n
	}
	if a[n-2] == 0 {
		a[n-2] = lower(a, n-2)
	}
	if a[n-1] == 0 {
		a[n-1] = lower(a, n-1)
	}
	return a[n-2] + a[n-1]

}
