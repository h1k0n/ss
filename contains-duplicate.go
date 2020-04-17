package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	length := len(nums)
	var ok bool
	for i := 0; i < length; i++ {
		if _, ok = m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}
