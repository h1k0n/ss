package main

import (
	"fmt"
)

func main() {
	k := 3
	fmt.Println("Hello, playground", len("y海"), k, k<<2, letterCombinations("237"))
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}
	table := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	nums := make([]int, len(digits))
	for k, v := range []byte(digits) {
		nums[k] = int(v - '0')
	}

	border := uint64(1)
	for i := 0; i < len(digits); i++ {
		border *= 4
	}
	//fmt.Println("--", nums,border)

	var j uint64
	var k uint64
	for j = 0; j < border; j++ {
		k = 3
		flg := false
		for z := 0; z < len(digits); z++ {
			m := j & k
			n := m >> (z * 2)
			if n >= uint64(len(table[nums[z]])) {
				flg = true
				break
			}
			//fmt.Println(".",z,n)
			k = k << 2

		}
		if flg == false {
			k = 3
			raw := make([]byte, len(digits))
			for z := 0; z < len(digits); z++ {
				m := j & k
				n := m >> (z * 2)

				raw[z] = table[nums[z]][n]
				k = k << 2
			}
			res = append(res, string(raw))
		}
	}
	return res
}
