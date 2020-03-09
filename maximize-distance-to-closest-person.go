package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", maxDistToClosest([]int{0, 0, 0, 1}))
}

func maxDistToClosest(seats []int) int {
	length := len(seats)
	max := 0
	begin := 0
	end := length - 1
	if seats[0] == 1 {
		for i := begin; ; i++ {
			if seats[i] == 0 {
				begin = i - 1
				break
			}
		}
	} else {
		for i := begin; ; i++ {
			if seats[i] == 0 {
				max++
			} else {
				begin = i
				break
			}
		}
	}
	if seats[end] == 1 {
		for i := end; ; i-- {
			if seats[i] == 0 {
				end = i + 1
				break
			}
		}
	} else {
		maxRight := 0
		for i := end; ; i-- {
			if seats[i] == 0 {
				maxRight++
			} else {
				end = i
				break
			}
		}
		if max < maxRight {
			max = maxRight
		}
	}

	k := 0
	latestK := 0

	for i := begin; i <= end; i++ {
		if seats[i] == 0 {
			k++
		} else {
			k = 0
		}
		if k == 0 {
			if max < (latestK-1)/2+1 {
				max = (latestK-1)/2 + 1
			}
		}
		latestK = k
	}
	return max
}
