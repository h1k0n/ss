package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",countSubstrings("abc"))
}

func countSubstrings(s string) int {
	length := len(s)
	count := 0
	for i := 0; i < length; i++ {
		for x := 0; i-x >= 0 && i+x < length; x++ {
			if s[i-x] == s[i+x] {
				count++
			} else {
				break
			}

		}
	}
	for i := 1; i < length; i++ {
		for x := 0; i-1-x >= 0 && i+x < length; x++ {
			if s[i-1-x] == s[i+x] {
				count++
			} else {
				break
			}
		}
	}
	return count
}
