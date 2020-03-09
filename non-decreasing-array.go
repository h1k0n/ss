package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", checkPossibility([]int{4, 5, 6, 3, 5, 7, 8}))
	//4, 2, 3
	//4, 1, 2, 3
	//0, 2, 1
	//0, 2, 3, 1
	//4, 5, 6, 3, 8, 9, 10
	//4, 5, 6, 3, 5, 7, 8
}

func checkPossibility(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	if n == 2 {
		return true
	}
	k := 2
	for i := 1; i < n; i++ {

		if nums[i] < nums[i-1] {
			k--
			if i-2 < 0 {
				continue
			}
			if i+1 < n {
				if nums[i-2] > nums[i] && nums[i-1] > nums[i+1] {
					return false
				}
			}
			if k == 0 {
				return false
			}
		}
	}
	return true

}
