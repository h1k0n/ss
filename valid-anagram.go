package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", isAnagram("radat", "car"))
}

func isAnagram(s string, t string) bool {
	arr := make([]int, 26)
	length1 := len(s)
	length2 := len(t)
	if length1 != length2 {
		return false
	} else {
		for _, v := range []byte(s) {
			arr[v-'a']++
		}
		for _, v := range []byte(t) {
			arr[v-'a']--
		}
		for _, v := range arr {
			if v != 0 {
				return false
			}
		}
	}
	return true
}
