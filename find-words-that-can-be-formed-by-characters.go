package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	//fmt.Println("Hello, playground", countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters(words []string, chars string) int {
	m := make(map[byte]int, len(chars))
	count := 0
	for _, v := range []byte(chars) {
		m[v] = m[v] + 1
	}
	for _, v := range words {
		str := []byte(v)
		length := len(str)
		flg := false
		i := 0
		for ; i < length; i++ {
			if m[str[i]] == 0 {
				flg = true
				break
			} else {
				m[str[i]]--
			}
		}
		if i > 0 {
			for j := i - 1; j >= 0; j-- {
				m[str[j]]++
			}
		}
		if flg == true {
			continue
		} else {
			count += length
		}
	}
	return count
}
