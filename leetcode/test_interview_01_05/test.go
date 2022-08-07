package test_interview_01_05

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	return dp[m][n] <= 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
