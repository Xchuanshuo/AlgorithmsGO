package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var a = make([]int, n+1)
	var dif = make(map[int]int, 0)
	var arr = make([]int, 0)
	for i := 1;i <= n;i++ {
		fmt.Scanln(&a[i])
		if a[i-1] < a[i] {
			if _, exist := dif[a[i-1]];!exist {
				arr = append(arr, a[i-1])
			}
			if _, exist := dif[a[i]];!exist {
				arr = append(arr, a[i])
			}
			dif[a[i-1]]++
			dif[a[i]]--
		}
	}
	sort.Ints(arr)
	res, sum := 0, 0
	for _, v := range arr {
		sum += dif[v]
		if sum > res { res = sum }
	}
	fmt.Println(res)
}
