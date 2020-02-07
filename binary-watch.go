package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", countBitNum(7))
	output := readBinaryWatch(2)
	fmt.Println("Hello, playground", output)
}

func readBinaryWatch(num int) []string {
	collect := make([]string, 0)
	var i uint16
	guard := uint16(1 << 10)
	for i = 0; i < guard; i++ {
		if num == countBitNum(i) {
			j := i
			h := j & (1<<4 - 1)
			if h > 11 {
				continue
			}
			j = j >> 4
			m := j & (1<<6 - 1)
			if m > 59 {
				continue
			}
			clock := fmt.Sprintf("%d:%02d", h, m)
			collect = append(collect, clock)
		}
	}
	return collect
}

func countBitNum(watch uint16) int {
	tmp := uint16(1)
	count := 0
	for i := 0; i < 10; i++ {
		if watch&(tmp<<i) > 0 {
			count++
		}
	}
	return count
}
