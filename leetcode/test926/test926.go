package test926

// dp[i][0] = 当前以0结尾满足递增需要翻转的次数
// dp[i][1] = 当前以1结尾满足递增需要翻转的次数
func minFlipsMonoIncr(s string) int {
	var n = len(s)
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= n; i++ {
		if s[i-1] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[n][0], dp[n][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
