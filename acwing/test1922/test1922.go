package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var arr = make([][]int, 0)
	for i := 0;i < n;i++ {
		var g, x int
		fmt.Scanf("%d %d", &g, &x)
		arr = append(arr, []int{x, g})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	l, r, sum, res := 0, 0, 0, 0
	for r < len(arr) {
		sum += arr[r][1]
		for l<r && arr[r][0] - arr[l][0] > 2*k {
			sum -= arr[l][1]
			l++
		}
		res = max(res, sum)
		r++
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b { return a}
	return b
}
