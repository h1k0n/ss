package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	nextPermutation([]int{2, 3, 1})
	nextPermutation([]int{1, 3, 2})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{5, 1, 1})
	nextPermutation([]int{5, 4, 7, 5, 3, 2})
	nextPermutation([]int{2, 2, 7, 5, 4, 3, 2, 2, 1})
	nextPermutation([]int{4, 2, 0, 2, 3, 2, 0})
	nextPermutation([]int{6})
}

func nextPermutation(nums []int) {
	length := len(nums)
	j := length - 2
	var i int
OUTER:
	for ; j >= 0; j-- {
		for i = length - 1; i > j; i-- {
			if nums[j] < nums[i] {
				//fmt.Println(i,j)
				break OUTER
			}
		}

	}

	if i == 0 {
		for i := 0; i < length/2; i++ {
			nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
		}
	} else {
		for x := i + 1; x < length-1; x++ {
			for y := length - 1; y > x; y-- {
				//fmt.Println("((", nums)
				if nums[y] < nums[y-1] {
					nums[y], nums[y-1] = nums[y-1], nums[y]
				}
			}
		}
		//fmt.Println("*", nums)
		for x := j + 1; x < i-1; x++ {
			for y := i - 1; y > x; y-- {
				//fmt.Println("((", nums)
				if nums[y] < nums[y-1] {
					nums[y], nums[y-1] = nums[y-1], nums[y]
				}
			}
		}
		//fmt.Println("*", nums)
		for z := j; z < i; z++ {
			k := j
			t := nums[j]
			for ; k < length-1; k++ {
				nums[k%length] = nums[(k+1)%length]
			}
			nums[k] = t
			//fmt.Println("-", nums)
		}
		//nums[(j+1)%length] = t
	}
	fmt.Println(nums)
}
