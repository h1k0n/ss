package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", sortByBits([]int{10, 100, 1000, 10000}))

}

func sortByBits(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		//x := 1
		beginBit := 14

		for j := 0; j < 14; j++ {
			if (1<<j)&arr[i] > 0 {
				beginBit++
				arr[i] += 1 << beginBit
			}
		}
		j := i
		t := arr[j]

		for ; j > 0; j-- {

			if t < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break
			}

		}
		if j < i {
			arr[j] = t
		}

	}
	a := 1<<16 - 1
	a = a << 15
	for i := 0; i < len(arr); i++ {

		tmp := arr[i] & a
		arr[i] -= tmp
	}
	//fmt.Println(arr, 1<<14)
	return arr
}
