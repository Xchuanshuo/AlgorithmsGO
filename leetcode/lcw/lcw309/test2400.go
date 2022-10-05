package lcw309

import "fmt"

func numberOfWays(startPos int, endPos int, k int) int {
	var n = endPos - startPos + 1
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for j := 1; j <= k; j++ {
		for i := 1; i <= n; i++ {
			dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % M
			dp[i][j] = (dp[i][j] + dp[i+1][j-1]) % M
		}
	}
	return dp[n][k]
}

var mp map[string]int
var M = int(1e9) + 7

func numberOfWays1(startPos int, endPos int, k int) int {
	mp = make(map[string]int)
	var dfs func(s, k int) int
	dfs = func(s, k int) int {
		var key = fmt.Sprintf("%d_%d", s, k)
		if v, exist := mp[key]; exist {
			return v
		}
		var res = 0
		if k <= 0 {
			if s == endPos {
				res = 1
			}
		} else {
			res = (dfs(s+1, k-1) + dfs(s-1, k-1)) % M
		}
		mp[key] = res
		return res
	}
	return dfs(startPos, k)
}
