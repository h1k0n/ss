package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, playground")
	fmt.Println("Hello, playground", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println("Hello, playground", maxSubArray([]int{-1, -4}))                        //-1
	fmt.Println("Hello, playground", maxSubArray([]int{-1, 1, -3}))                     //1
	fmt.Println("Hello, playground", maxSubArray([]int{1, -1, -2}))                     //1

	fmt.Println("Hello, playground", maxSubArray([]int{3, 1, -1})) //4
	fmt.Println("Hello, playground", maxSubArray([]int{1, 2}))     //3
	fmt.Println("Hello, playground", maxSubArray([]int{1, 1, -2})) //2

	fmt.Println("Hello, playground", maxSubArray([]int{0, -3, 1, 1}))            // 2
	fmt.Println("Hello, playground", maxSubArray([]int{0, -2, -3}))              //0
	fmt.Println("Hello, playground", maxSubArray([]int{3, -2, -3, -3, 1, 3, 0})) //4

}

func maxSubArray(nums []int) int {
	minSum := nums[0]
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//if nums[i] > max {
		//	max = nums[i]
		//}
		sum += nums[i]
		//fmt.Println("sum minSum", sum, minSum)

		if minSum > 0 && sum > max {
			max = sum
		} else if sum-minSum > max {
			max = sum - minSum
		}
		if sum < minSum {
			minSum = sum
		}

	}
	return max
}
