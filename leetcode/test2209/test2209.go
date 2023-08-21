package test2209

/**
 * @Description https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/
 * idea: dp[i][j]表示前i块砖块总共使用j条毯子能覆盖的最大白色(变成黑色)砖块数,
		 当前覆盖的结束位置为i，则dp[i][j] = max(dp[i-1][j]+w, dp[i-l][j-1] + l)
		 初始条件，dp[i][0] = sum(黑色砖块)，结果: n - dp[n][k]
 **/

func minimumWhiteTiles(f string, k int, l int) int {
	var n = len(f)
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		var w = 1 ^ int(f[i-1]-'0')
		dp[i][0] = dp[i-1][0] + w
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			var w = 1 ^ int(f[i-1]-'0')
			if i <= l {
				dp[i][j] = i
			} else {
				dp[i][j] = max(dp[i-1][j]+w, dp[i-l][j-1]+l)
			}
		}
	}
	return n - dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}