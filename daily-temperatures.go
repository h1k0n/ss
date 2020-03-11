package main

import (
	"fmt"
)

func main() {

	//res:=dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	res := dailyTemperatures([]int{73, 72})
	fmt.Println("Hello, playground", res)
}

func dailyTemperatures(T []int) []int {
	length := len(T)
	if length == 1 {
		return []int{0}
	}
	st := make([][2]int, length)
	res := make([]int, length)
	top := -1
	top++
	st[top] = [2]int{0, T[0]}

	for i := 1; i < length; i++ {
		//fmt.Println("*", i, T[i], st)
		for top != -1 && T[i] > st[top][1] {
			//fmt.Println(".", st[top][1])
			res[st[top][0]] = i - st[top][0]
			top--
		}

		top++
		st[top] = [2]int{i, T[i]}

	}
	//fmt.Println("z", st, top)
	return res
}
