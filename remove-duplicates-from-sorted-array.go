package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",removeDuplicates([]int{1<<31-1,1<<31-1}))
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	g := nums[0] - 1
	top := -1
	for i := 0; i < len(nums); i++ {
		if g != nums[i] {
			g = nums[i]
			top++
			nums[top] = g

		}

	}
	//fmt.Println(nums)
	return top + 1
}
