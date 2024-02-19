package lcwt118

import "math"

func minimumCoins(prices []int) int {
	var n = len(prices)
	var dp = make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n+1 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		var nxt = math.MaxInt32
		for j := i + 1; j <= i+i+1; j++ {
			// j开始买
			nxt = min(nxt, dfs(j))
		}
		dp[i] = prices[i-1] + nxt
		return dp[i]
	}
	return dfs(1)
}
