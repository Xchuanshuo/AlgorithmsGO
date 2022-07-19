package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var arr = make([]int, 0)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scanf("%d", &v)
		arr = append(arr, v)
	}
	sort.Ints(arr)
	var dirs = make([]int, n)
	dirs[0] = 1
	for i := 1; i < n-1; i++ {
		if arr[i]-arr[i-1] > arr[i+1]-arr[i] {
			dirs[i] = 1
		}
	}
	res := 0
	for i := 1; i < n; i++ {
		if dirs[i] == 0 && dirs[i-1] == 1 {
			res++
			if i-2 >= 0 && dirs[i-2] == 1 && i+1 < n && dirs[i+1] == 0 {
				res++
			}
		}
	}
	fmt.Println(res)
}
