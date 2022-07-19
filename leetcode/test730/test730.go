package test730

const M int64 = 1e9 + 7

func countPalindromicSubsequences(s string) int {
	var n = len(s)
	var dp = make([][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int64, 4)
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i][s[i]-'a'] = 1
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := l + i - 1
			var idx = s[i] - 'a'
			for k := 0; i < 4; i++ {
				dp[i][j][k] = dp[i+1][j-1][k]
			}
			if s[i] == s[j] {
				dp[i][j][idx] = (sum(dp[i+1][j-1]) + 2) % M
			} else {
				dp[i][j][s[i]-'a'] = (sum(dp[i+1][j]) + 1) % M
				dp[i][j][s[j]-'a'] = (sum(dp[i][j-1]) + 1) % M
			}
		}
	}
	return int(sum(dp[0][n-1]))
}

func sum(dp []int64) int64 {
	var res int64 = 0
	for i := 0; i < len(dp); i++ {
		res = (res + dp[i]) % M
	}
	return res
}
