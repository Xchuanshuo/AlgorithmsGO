package lcw314

/**
* @Description https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
* idea: 取余dp模板题
        dp[i][j][(k+v)%K] = dp[i-1][j][k] + dp[i][j-1][k]
**/

func numberOfPaths(grid [][]int, K int) int {
	m, n := len(grid), len(grid[0])
	var M = int(1e9) + 7
	var dp = make([][][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, K+1)
		}
	}
	dp[0][1][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < K; k++ {
				var v = grid[i][j]
				dp[i+1][j+1][(k+v)%K] = (dp[i+1][j][k] + dp[i][j+1][k]) % M
			}
		}
	}
	return dp[m][n][0]
}
