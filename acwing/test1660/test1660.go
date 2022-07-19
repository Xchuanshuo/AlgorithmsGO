package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a = make([][]int, 0)
	for i := 0; i < n; i++ {
		var x, s int
		fmt.Scanf("%d %d\n", &x, &s)
		a = append(a, []int{x, s})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	var r = n
	for i := 1; i < len(a); i++ {
		if a[i][1] != a[i-1][1] {
			r = min(r, a[i][0]-a[i-1][0]-1)
		}
	}
	var res = 0
	for i := 0; i < len(a); i++ {
		if a[i][1] == 1 {
			res++
			i++
			for i < n && a[i][0]-a[i-1][0] <= r {
				i++
			}
			i--
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
