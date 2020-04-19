package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",removePalindromeSub(""))
}

//remove one palindromic subsequence
func removePalindromeSub(s string) int {
	length := len(s)
	if s == "" {
		return 0
	}
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return 2
		}
	}
	return 1
}
