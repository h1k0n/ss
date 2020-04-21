package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", reformat("leetcode"))
}

func reformat(s string) string {
	length := len(s)
	digit := make([]byte, 0)
	alphabet := make([]byte, 0)
	for _, v := range []byte(s) {
		if v >= '0' && v <= '9' {
			digit = append(digit, v)
		} else {
			alphabet = append(alphabet, v)
		}
	}
	//fmt.Println(digit, alphabet)
	if len(alphabet)-len(digit) == 1 || len(alphabet) == len(digit) {
		res := make([]byte, length)
		for i := 0; i < len(alphabet); i++ {
			res[2*i] = alphabet[i]
		}
		for i := 0; i < len(digit); i++ {
			res[2*i+1] = digit[i]
		}
		return string(res)
	} else if len(alphabet)+1 == len(digit) {
		res := make([]byte, length)
		for i := 0; i < len(digit); i++ {
			res[2*i] = digit[i]
		}
		for i := 0; i < len(alphabet); i++ {
			res[2*i+1] = alphabet[i]
		}
		return string(res)
	} else {
		return ""
	}
}
