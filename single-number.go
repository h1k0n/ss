package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", xor(xor(100, 46), 46))
	fmt.Println("Hello, playground",singleNumber([]int{4,1,2,1,2}))
}

func xor(a, b int) int {
	return a ^ b
}

func singleNumber(nums []int) int {
    var res int
    for i:=range nums {
       res=res^nums[i]
    }
    return res
}
