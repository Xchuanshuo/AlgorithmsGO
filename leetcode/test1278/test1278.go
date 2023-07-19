package test1278

func palindromePartition(s string, k int) int {
	var n = len(s)
	var cost = make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			var v = 0
			if s[i] != s[j] {
				v = 1
			}
			cost[i][j] = cost[i+1][j-1] + v
		}
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = int(1e9)
		}
	}
	dp[0][0] = 0
	for l := 1; l <= k; l++ {
		for i := 1; i <= n; i++ {
			for j := i; j >= 1; j-- {
				dp[i][l] = min(dp[i][l], dp[j-1][l-1]+cost[j-1][i-1])
			}
		}
	}
	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
