package main

import (
	"fmt"
)

func main() {
	
	n:=Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println("Hello, playground",n.SumRange(0,0))
	
}

type NumArray struct {
	PrefixSum []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	a := make([]int, length)
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		a[i] = sum
	}
	return NumArray{
		PrefixSum: a,
	}

}

func (this *NumArray) SumRange(i int, j int) int {
	if i != 0 {
		return this.PrefixSum[j] - this.PrefixSum[i-1]
	}
	return this.PrefixSum[j]
}
