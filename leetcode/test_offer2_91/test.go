package test_offer2_91

func minCost(costs [][]int) int {
	var n = len(costs)
	var dp = make([][]int, n+1)
	dp[0] = make([]int, 3)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 3)
		dp[i][0] = costs[i-1][0] + min(dp[i-1][1], dp[i-1][2])
		dp[i][1] = costs[i-1][1] + min(dp[i-1][0], dp[i-1][2])
		dp[i][2] = costs[i-1][2] + min(dp[i-1][0], dp[i-1][1])
	}
	return min(dp[n][0], min(dp[n][1], dp[n][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
