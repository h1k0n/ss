package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	res := []int{}

	length := len(nums)
	m := make(map[int]int, length)
	for i := 0; i < length; i++ {
		m[nums[i]] = i
	}
	for i := 0; i < length; i++ {

		if index, ok := m[target-nums[i]]; ok {
			if i == index {
				continue
			}
			res = append(res, i, index)
			return res
		}

	}
	return res
}
