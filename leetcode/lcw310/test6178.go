package main

import "sort"

func minGroups(intervals [][]int) int {
	var arr = make([][]int, 0)
	for _, inv := range intervals {
		arr = append(arr, []int{inv[0], 1})
		arr = append(arr, []int{inv[1] + 1, -1})
	}
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	res, cnt := 0, 0
	for _, a := range arr {
		cnt += a[1]
		if cnt > res {
			res = cnt
		}
	}
	return res
}
