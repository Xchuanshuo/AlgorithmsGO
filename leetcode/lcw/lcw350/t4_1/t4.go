package t4_1

// 0-1背包问题
// dp[i][j]表示付费工人对于前i面墙，与免费工人总共涂了j面的最小总开销
// 状态转移, 不涂i面墙, dp[i][j]=dp[i-1][j]
//		     涂i面墙, dp[i][j]=dp[i-1][j-time[i]-1] + cost[i]
func paintWalls(cost []int, time []int) int {
	var n = len(cost)
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = int(1e9)
		}
	}
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			dp[i][j] = min(dp[i][j], min(dp[i-1][j], dp[i-1][max(j-time[i-1]-1, 0)]+cost[i-1]))
		}
	}
	return dp[n][n]
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
