package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground",majorityElement([]int{2,2,1,4,5,2,2}))
}

func majorityElement(nums []int) int {
    m:=make(map[int]int,len(nums)/2+1)
    for i:=0;i<len(nums) ;i++{
        m[nums[i]]++
    }
    maj:=nums[0]
    for k,v:=range m {
        if v>m[maj] {
            maj=k
        }
    }
    return maj
}
