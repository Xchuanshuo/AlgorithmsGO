package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var paths = make([][]int, 0)
	for i := 0;i < n;i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		paths = append(paths, []int{a, b})
	}
	sort.Slice(paths, func(i, j int) bool {
		return paths[i][0] < paths[j][0]
	})
	minR := make([]int, n)
	minR[n-1] = paths[n-1][1]
	for i := len(paths)-2;i >= 0;i-- {
		minR[i] = paths[i][1]
		if minR[i+1] < paths[i][1] {
			minR[i] = minR[i+1]
		}
	}
	res, lMax := 0, paths[0][1]
	for i := 0;i < len(paths);i++ {
		if paths[i][1]>=lMax && paths[i][1]<= minR[i] {
			res++
		}
		if lMax < paths[i][1] { lMax = paths[i][1] }
	}
	fmt.Println(res)
}
