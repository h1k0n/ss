package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(nums,2)
	rotate2(nums, 3)
	fmt.Println("Hello, playground", nums)

}

func rotate(nums []int, k int) {
	lastIndex := len(nums) - 1
	for i := 0; i < k; i++ {
		tmp := nums[lastIndex]
		for j := lastIndex - 1; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = tmp
	}
}

func rotate1(nums []int, k int) {
	length := len(nums)
	nums1 := make([]int, k)
	copy(nums1, nums[length-k:length])
	for i := length - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	copy(nums[0:k+1], nums1)
}

func rotate2(nums []int, k int) {
	length := len(nums)
	x := k % length
	nums1 := make([]int, x)
	copy(nums1, nums[length-x:length])
	for i := length - x - 1; i >= 0; i-- {
		nums[i+x] = nums[i]
	}
	copy(nums[0:x+1], nums1)
}
