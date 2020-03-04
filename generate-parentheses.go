package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", gen(3), generateParenthesis(5))

}

func gen(n int) uint64 {
	x := uint64(0)
	for i := 0; i < n; i++ {
		x += 1 << (2*n - i - 1)
	}
	return x
}

func generateParenthesis(n int) []string {
	return xxafs(2*(n-1), n-1)
}


func xxafs(n, r int) []string {
	res := make([]string, 0)
	if n < 0 {
		return res
	}
	if n == 0 {
		return []string{"()"}
	}
	a := make([]int, r)
	st := make([]byte, n+2)

	top := -1
	arr := make([]byte, n+2)
	arr[0] = '('
	for j := 1; j < n+2; j++ {
		arr[j] = ')'
	}
	// initialize first combination
	for i := 0; i < r; i++ {
		a[i] = i
	}
	i := r - 1 // Index to keep track of maximum unsaturated element in array
	// a[0] can only be n-r+1 exactly once - our termination condition!
	for a[0] < n-r+1 {
		// If outer elements are saturated, keep decrementing i till you find unsaturated element
		for i > 0 && a[i] == n-r+i {
			i--
		}
		//fmt.Println(a) // pseudo-code to print array as space separated numbers

		for j := 0; j < r; j++ {
			arr[a[j]+1] = '('
		}
		top = -1
		flg := false
		for _, v := range arr {
			if v == '(' {
				top++
				st[top] = '('

			} else {
				top--
			}
			if top < -1 {
				flg = true
				break
			}

		}

		if flg == false {
			//fmt.Println(string(arr))
			res = append(res, string(arr))
		}
		for j := 0; j < r; j++ {
			arr[a[j]+1] = ')'
		}

		a[i]++
		// Reset each outer element to prev element + 1
		for i < r-1 {
			a[i+1] = a[i] + 1
			i++
		}
	}
	//fmt.Println(len(res))
	return res
}
