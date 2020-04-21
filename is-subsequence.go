package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", isSubsequence("axc", "ahbgdcd"))
}

func isSubsequence(s string, t string) bool {
	var i, j int
	var v byte
	lengthT := len(t)
	for i, v = range []byte(s) {

		for j < lengthT {
			if t[j] == v {
				j++
				break
			}
			j++
		}
		if j == lengthT {
			if i == len(s)-1 && v == t[j-1] {
				//fmt.Println("debug")
				return true
			}
			return false
		}

	}
	return true
}
