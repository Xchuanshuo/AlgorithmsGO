package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	arr := make([]int, n+1)
	for i := 0;i < k;i++ {
		var l, r int
		fmt.Scanln(&l, &r)
		l -= 1
		r -= 1
		arr[l] += 1
		arr[r+1] -= 1
	}
	for i := 1;i < len(arr);i++ {
		arr[i] += arr[i-1]
	}
	sort.Ints(arr)
	fmt.Println(arr[len(arr)/2])
}
