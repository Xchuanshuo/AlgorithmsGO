package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([][]int, 0)
	for i := 0; i < n; i++ {
		var s, e int
		fmt.Scanln(&s, &e)
		arr = append(arr, []int{s, e})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	res := 0
	for i := 0; i < n; i++ {
		last, cur := -1, 0
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			start, end := arr[j][0], arr[j][1]
			if start > last {
				cur += end - start
				last = end
			} else {
				cur += max(0, end-last)
				last = max(last, end)
			}
		}
		if cur > res {
			res = cur
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
