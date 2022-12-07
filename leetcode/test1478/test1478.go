package test1478

import "sort"

/**
* @Description https://leetcode.cn/problems/allocate-mailboxes/
* idea: 与test2463类似, 将房子位置排序后，每个邮局一定是负责连续的一批房子
		对于连续一批的房子，邮局修在这一批房子【中间】，能保证距离最短，所以可以先预处理区间和
		设dp[i][j] 表示给前j个房子修i个邮局的最小距离总和
        则dp[i][j]=min(dp[i-1][j], dp[i-1][j-k] + cost[j-k..j])
		时间复杂度 O(n^3)
**/

func minDistance(houses []int, m int) int {
	sort.Ints(houses)
	var n = len(houses)
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = 1e9
		}
	}
	var sums = make([][]int, n)
	for i := 0; i < n; i++ {
		sums[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			sums[i][j] = houses[j] - houses[i] + sums[i+1][j-1]
		}
	}
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= j; k++ {
				dp[i][j] = min(dp[i][j], dp[i-1][j-k]+sums[j-k][j-1])
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
