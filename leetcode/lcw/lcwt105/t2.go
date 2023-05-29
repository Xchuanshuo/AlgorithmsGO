package lcwt105

func minExtraChar(s string, dictionary []string) int {
	var mp = make(map[string]bool)
	for _, d := range dictionary {
		mp[d] = true
	}
	var n = len(s)
	var dp = make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n
		for j := 1; j <= i; j++ {
			if mp[s[j-1:i]] {
				dp[i] = min(dp[i], dp[j-1])
			} else {
				dp[i] = min(dp[i], dp[j-1]+i-j+1)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
