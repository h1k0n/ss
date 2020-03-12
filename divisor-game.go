package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",divisorGame(3))
}
func divisorGame(N int) bool {
	if N%2 == 0 {
		return true
	}
	return false
}
