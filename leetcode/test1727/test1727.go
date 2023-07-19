package test1727

import "sort"

func largestSubmatrix(g [][]int) int {
	m, n := len(g), len(g[0])
	var h = make([]int, n)
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 0 {
				h[j] = 0
			} else {
				h[j] += 1
			}
		}
		var t = make([]int, n)
		copy(t, h)
		sort.Ints(t)
		for j := 0; j < n; j++ {
			if t[j] != 0 {
				res = max(res, t[j]*(n-j))
			}
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
