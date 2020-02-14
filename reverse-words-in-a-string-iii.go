//https://play.golang.org/p/_3L1lzraZtu
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", reverseWords(" Let's take LeetCode contest "))
}

func reverseWords(s string) string {
	var start, end int
	t := []byte(s)
	for i := 0; i < len(t); i++ {
		//fmt.Printf("%c",t[i])
		if t[i] == ' ' {
			//fmt.Printf("%c\n",t[i])
			start = i + 1
			continue
		}
		if i+1 == len(t) || t[i+1] == ' ' {

			end = i
			//fmt.Println(start, end)
			for j := 0; ; j++ {
				if start+j >= end-j {
					break
				}
				t[start+j], t[end-j] = t[end-j], t[start+j]
			}
		}
	}
	return string(t)
}
