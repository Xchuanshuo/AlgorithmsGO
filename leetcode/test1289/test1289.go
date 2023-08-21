package test1289

import "math"

func minFallingPathSum(g [][]int) int {
	var n = len(g[0])
	var dp = make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = g[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
			for k := 0; k < n; k++ {
				if k == j {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+g[i][j])
			}
		}
	}
	var res = math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
