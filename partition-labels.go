package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println("Hello, playground", partitionLabels("mlullbhiuiujgvwvurcdvhzdk"))
}

func partitionLabels(S string) []int {
	res := make([]int, 0)
	arr := []byte(S)

	m := make(map[byte]int)
	//mm := make(map[byte]int)
	mm := make([]int, 26)
	for i, v := range arr {
		if _, ok := m[v]; !ok {
			mm[v-'a'] = i
		}
		m[v] = i
	}
	//fmt.Printf("%+v\n%+v\n", mm, m)

	begin := -1
	ca := arr[0]
	max := m[ca]
	for {
	Top:
		for k, v := range mm {
			if v < max {
				if m[byte('a'+k)] > max {
					//fmt.Printf("%c,%d\n", 'a'+k, m[byte('a'+k)])
					max = m[byte('a'+k)]
					goto Top
				}
			}
		}
		//fmt.Printf("%c\n", ca)
		//fmt.Println(max)

		res = append(res, max-begin)
		begin = max
		if max+1 < len(arr) {

			ca = arr[max+1]

			max = m[ca]

		} else {
			break
		}

		//
	}
	return res

}
