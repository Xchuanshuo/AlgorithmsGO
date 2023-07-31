package test2136

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	var arr = make([][]int, 0)
	for i, p := range plantTime {
		arr = append(arr, []int{p, growTime[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	cur, mx := 0, 0
	for _, v := range arr {
		cur += v[0]
		mx = max(mx, cur+v[1])
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
