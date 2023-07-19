package test1463

// dp[i][j][k] 滚动数组优化
func cherryPickup(g [][]int) int {
	m, n := len(g), len(g[0])
	var dp = make([][]int, n+2)
	for j := 0; j <= n+1; j++ {
		dp[j] = make([]int, n+2)
		for k := 0; k < len(dp[j]); k++ {
			dp[j][k] = -int(1e9)
		}
	}
	dp[1][n] = g[0][0] + g[0][n-1]
	for i := 2; i <= m; i++ {
		var cur = make([][]int, n+2)
		for t := 0; t < len(cur); t++ {
			cur[t] = make([]int, n+2)
		}
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				var mx = -int(1e9)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						mx = max(mx, dp[j+x][k+y])
					}
				}
				cur[j][k] = mx + g[i-1][j-1] + g[i-1][k-1]
				if j == k {
					cur[j][k] -= g[i-1][j-1]
				}
			}
		}
		dp = cur
	}
	var res = 0
	for j := 1; j <= n; j++ {
		for k := 1; k <= n; k++ {
			res = max(res, dp[j][k])
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
