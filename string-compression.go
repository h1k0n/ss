package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", compress([]byte{'a'}))
}

func compress(chars []byte) int {
	count := 0
	shortCount := 0
	tmp := byte(0)
	top := 0
	for i, v := range chars {
		if v == tmp {
			shortCount++
			
			//fmt.Println("if", shortCount)
			wei:=xxx(shortCount)
			if i == len(chars)-1 && shortCount >= 1 {
				count += wei
			}
			modify(chars[top+1:top+1+wei], shortCount)
		} else {
			//fmt.Printf("else %c\n", v)
			tmp = v

			if shortCount > 1 {
				count += xxx(shortCount)
			}
			shortCount = 1
			top = count
			count++
			
			chars[top]=tmp
		}

	}
	//fmt.Printf("%s\n", chars[:count])
	return count
}

func xxx(num int) int {
	i := num
	c := 0
	for i != 0 {
		c++
		i = i / 10
	}
	return c
}

func modify(b []byte, num int) {
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = byte('0' + num%10)
		num = num / 10
	}
}
