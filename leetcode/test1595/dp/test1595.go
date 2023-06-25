package dp

import "math"

func connectTwoGroups(cost [][]int) int {
	m, n := len(cost), len(cost[0])
	var minCost = make([]int, n)
	for i := 0; i < n; i++ {
		var cur = cost[0][i]
		for j := 0; j < m; j++ {
			cur = min(cur, cost[j][i])
		}
		minCost[i] = cur
	}
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32 / 2
		}
	}
	for j := 0; j < 1<<n; j++ {
		var sum = 0
		for k := 0; k < n; k++ {
			if j&(1<<k) > 0 {
				sum += minCost[k]
			}
		}
		dp[0][j] = sum
	}
	for i := 1; i <= m; i++ {
		for j := 0; j < 1<<n; j++ {
			var cur = math.MaxInt32 / 2
			for k := 0; k < n; k++ {
				cur = min(cur, dp[i-1][j&^(1<<k)]+cost[i-1][k])
			}
			dp[i][j] = cur
		}
	}
	// A集合剩余m个元素, B集合全部未选时的最小方案数
	return dp[m][1<<n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
