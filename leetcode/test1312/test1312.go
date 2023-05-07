package test1312

func minInsertions(s string) int {
	var n = len(s)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return n - dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
