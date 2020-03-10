package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", countPrimes(20))
}

/*
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	fmt.Println(isPrime)
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
			fmt.Println("*", j, isPrime)
		}
		fmt.Println(".", i, isPrime)
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
*/

func countPrimes(n int) int {

	if n <= 2 {
		return 0
	}

	a := make([]int, n)
	a[0] = 2
	i := 3
	count := 1
Top:
	if i < n {
		k := 0
		for j := a[k]; j*j <= i; {
			if i%j == 0 {
				i++
				goto Top
			}
			k++
			j = a[k]
		}
		count++
		a[count-1] = i
		i++
		goto Top
	}
	return count
}
