package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	x := letterCasePermutation("adasf")
	fmt.Println("Hello, playground", x)
}

func letterCasePermutation(S string) []string {
	res := make([]string, 0)

	raw := []byte(S)
	indexArray := make([]int, 0)
	for i, v := range raw {
		if v >= 'A' { // V-'A'>=0 总是true
			//fmt.Println(v, '3', v-'3')
			indexArray = append(indexArray, i)
		}
	}

	n := len(indexArray)
	if n == 0 {
		//fmt.Println("n", n)
		return []string{S}
	}
	a := make([]int, n)
	length := len(a)
	k := 0

	for k >= 0 {
		a[k] = a[k] + 1
		//fmt.Println("k,a[k]", k, a[k])
		if a[k] < 3 && k == length-1 {
			//fmt.Println("a", a)
			tmp := []byte(S)
			for j := 0; j < n; j++ {
				if a[j] > 1 {
					index := indexArray[j]
					if tmp[index] < 'a' {
						tmp[index] = tmp[index] + 32
					} else {
						tmp[index] = tmp[index] - 32
					}
				}
			}
			res = append(res, string(tmp))
		} else if a[k] < 3 && k < length-1 {
			k = k + 1
		} else if a[k] >= 3 {
			//fmt.Println("---",k,a[k])
			for i := k; i < length; i++ {
				a[i] = 0
			}
			k = k - 1
			//a[k] = a[k] + 1
			//k++
		}

	}
	//fmt.Println("over")
	return res
}
