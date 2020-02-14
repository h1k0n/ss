//https://play.golang.org/p/EZaoBBxe7wK
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", findDisappearedNumbers1([]int{4, 3, 2, 7, 7, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	counters := make([]int, len(nums))
	var ret []int
	for _, v := range nums {
		counters[v-1]++
	}
	for i, v := range counters {
		if v == 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func findDisappearedNumbers1(nums []int) []int {
	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums)-i; j++ {
			if i+1 == nums[i] {
				//fmt.Print("b", j)
				break
			}
			index := nums[i] - 1
			if nums[i]== nums[index] {
				//fmt.Print("b", j)
				break
			}
			nums[i], nums[index] = nums[index], nums[i]
			//fmt.Print("x", j)
			//fmt.Println("z", nums)

		}

		//fmt.Println("s",i, nums)
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {

		for {
			if nums[i] == 0 {
				break
			}
			if nums[i] == i+1 {
				nums[i] = 0
				break
			} else {
				index := nums[i] - 1
				if nums[index] != 0 {
					nums[i], nums[index] = nums[index], 0
				} else {
					break
				}
			}

		}

		//fmt.Println("s", i, nums)
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}
