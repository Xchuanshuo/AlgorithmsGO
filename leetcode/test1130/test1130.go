package test1130

import (
	"math"
)

func mctFromLeafValues(arr []int) int {
	var n = len(arr)
	var mx = make([][]int, n)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		mx[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
		dp[i][i] = 0
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if l == 1 {
				mx[i][j] = arr[i]
			} else {
				mx[i][j] = max(mx[i+1][j], arr[i])
			}
		}
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			if l == 2 {
				dp[i][j] = arr[i] * arr[j]
				continue
			}
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+mx[i][k]*mx[k+1][j])
				dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k][j]+mx[i][k-1]*mx[k][j])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
