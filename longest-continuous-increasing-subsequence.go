package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", findLengthOfLCIS([]int{2,2,2,2,2}))
}

func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	k := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			k++
		} else {
			k = 1
		}
		if max < k {
			max = k
		}

	}
	return max
}
