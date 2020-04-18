package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, playground", stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

func stringMatching(words []string) []string {
	lenArr := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		lenArr[i] = len(words[i])
	}
	res := make([]string, 0)
	wordsLen := len(words)
	var a, b int
	var flg bool
	for i := 0; i < wordsLen; i++ {
		for j := i + 1; j < wordsLen; j++ {
			//fmt.Println("-", words[i], words[j])
			a, b = j, i
			if lenArr[j] < lenArr[i] {
				a, b = i, j
			}
			if strings.Contains(words[a], words[b]) {
				for _, v := range res {
					//fmt.Println("-", v, words[b])
					if v == words[b] {
						flg = true
						break
					}
				}
				if flg == true {
					flg = false
					continue
				} else {
					res = append(res, words[b])
					flg = false
				}
			}
		}
	}
	return res
}
