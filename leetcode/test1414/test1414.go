package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	f := []int{1, 1}
	for i := 2; f[i-1] <= k; i++ {
		cur := f[i-1] + f[i-2]
		f = append(f, cur)
	}
	res := 0
	for i := len(f) - 1; i >= 0; i-- {
		if k >= f[i] {
			k -= f[i]
			res++
		}
	}
	return res
}

func main() {
	k := 513314
	res := findMinFibonacciNumbers(k)
	fmt.Println(res)
}
