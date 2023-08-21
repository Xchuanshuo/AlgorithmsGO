package lcw359

import (
	"sort"
)

func longestEqualSubarray(nums []int, k int) int {
	var group = make(map[int][]int)
	for i, v := range nums {
		group[v] = append(group[v], i)
	}
	var sums = make([][]int, 0)
	for _, g := range group {
		var sum = make([]int, len(g)+1)
		for i := range g {
			if i > 0 {
				sum[i+1] = sum[i] + g[i] - g[i-1] - 1
			} else {
				sum[i+1] = 0
			}
		}
		sum = sum[1:]
		sums = append(sums, sum)
	}
	var res = 0
	for _, sum := range sums {
		for i := 0; i < len(sum); i++ {
			var j = sort.SearchInts(sum, sum[i]-k)
			res = max(res, i-j+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
