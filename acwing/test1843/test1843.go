package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr, sums := make([]int,n), make([]int, n+1)
	var sum = 0
	for i := 0;i < n;i++ {
		var v int
		fmt.Scanln(&v)
		arr[i] = v
		sums[i+1] = sums[i] + arr[i]
		sum += i*arr[i]
	}
	res := int(1e9)
	for i := 0;i < n;i++ {
		cur := sum - i*arr[i] + (n-i)*sums[i] - i*(sums[n] - sums[i+1])
		if cur < res { res = cur }
	}
	fmt.Println(res)
}
