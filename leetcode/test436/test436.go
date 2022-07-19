package test436

import "sort"

func findRightInterval(intervals [][]int) []int {
	var arr = make([][]int, 0)
	var n = len(intervals)
	for i := 0; i < n; i++ {
		arr = append(arr, []int{intervals[i][0], i})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	arr = append(arr, []int{1e9, -1})
	var res = make([]int, n)
	for _, cur := range arr {
		if cur[1] == -1 {
			continue
		}
		var pos = sort.Search(n, func(i int) bool {
			return arr[i][0] >= intervals[cur[1]][1]
		})
		res[cur[1]] = arr[pos][1]
	}
	return res
}
