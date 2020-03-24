package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println("Hello, playground", findUnsortedSubarray([]int{1, 1, 1, 1}))
	fmt.Println("Hello, playground", findUnsortedSubarray([]int{2, 1}))
}

func findUnsortedSubarray(nums []int) int {

	length := len(nums)
	if length == 1 {
		return 0
	}
	x, y := 0, length-1
O1:
	for j := 0; j < length-1; j++ {
		i := j + 1
		for ; i < length; i++ {
			if nums[j] > nums[i] {

				break O1
			}
		}
		//fmt.Println("a", j, i)
		if i == length {
			x = j+1
		}
	}
O2:
	for j := length - 1; j > 0; j-- {
		i := j - 1
		for ; i >= 0; i-- {
			if nums[j] < nums[i] {

				break O2
			}
		}
		//fmt.Println("a", j, i)
		if i == -1 {
			y = j-1
		}
		if x >= y {
			return 0
		}
	}

	return y - x + 1
}
