package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", isValid("{"))
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	top := -1
	st := make([]byte, len([]byte(s)))
	for _, v := range []byte(s) {
		if v == '(' || v == '{' || v == '[' {
			top++
			st[top] = v
		} else {
			if top < 0 {
				return false
			}
			if (st[top] == '(' && v == ')') || (st[top] == '[' && v == ']') || (st[top] == '{' && v == '}') {
				top--
			} else {
				return false
			}

		}
	}
	if top > -1 {
		return false
	}
	return true

}
