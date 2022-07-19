package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var arr = make([][]int, 0)
	for i := 0; i < n; i++ {
		var s, e, v int
		fmt.Scanf("%d %d %d", &s, &e, &v)
		arr = append(arr, []int{s, v})
		arr = append(arr, []int{e, -v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	sum, res := 0, 0
	for _, a := range arr {
		sum += a[1]
		if sum > res {
			res = sum
		}
	}
	fmt.Println(res)
}
