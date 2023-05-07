package lcwt99

import (
	"sort"
)

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var pre = []int{-1, -1}
	var group = 0
	for _, cur := range ranges {
		if cur[0] > pre[1] {
			group++
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	var M = int(1e9) + 7
	var res = 1
	for i := 1; i <= group; i++ {
		res = (res * 2) % M
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
