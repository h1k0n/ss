package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{-1,4}
	fmt.Println("Hello, playground")
	merge(nums1, 3, nums2, 2)
	fmt.Println("Hello, playground", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
        var i,j int
	for i = 0; i < n; i++ {
		for j = 0; j < m+i; j++ {
			if nums2[i] < nums1[j] {
				for k := m + i; k > j; k-- {
					nums1[k] = nums1[k-1]
				}
				nums1[j] = nums2[i]
				break
			}
		}
		if j == m+i {
			nums1[m+i] = nums2[i]
		}
	}
}
