package test1878

import (
	"sort"
)

func getBiggestThree1(g [][]int) []int {
	m, n := len(g), len(g[0])
	var dp = make([][][2]int, m+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, n+2)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var v = g[i-1][j-1]
			dp[i][j][0] = dp[i-1][j+1][0] + v
			dp[i][j][1] = dp[i-1][j-1][1] + v
		}
	}
	var mn = min(m, n)
	var set = make(map[int]bool)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for w := 0; w <= mn/2; w++ {
				if i-w <= 0 || i+w > m || j-w <= 0 || j+w > n {
					break
				}
				var r1 = dp[i][j-w][0] - dp[i-w][j][0] + g[i-w-1][j-1] // 右上1
				var l1 = dp[i][j+w][1] - dp[i-w][j][1]                 // 左上1
				var r2 = max(dp[i+w][j][0]-dp[i][j+w][0], 0)           // 右上2
				var l2 = max(dp[i+w-1][j-1][1]-dp[i][j-w][1], 0)       // 左上2
				var cur = l1 + l2 + r1 + r2
				set[cur] = true
			}
		}
	}
	var a = make([]int, 0)
	for k := range set {
		a = append(a, k)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	return a[0:min(len(a), 3)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
