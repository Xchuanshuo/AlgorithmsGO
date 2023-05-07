package test1000

import "math"

// dp[i][j]表示区间[i,j]合并到不能合并为止的最低成本
// 不能合并: j-i+1 <= k-1，
// 由于每次合并减少k-1堆石头，区间能完全合并成一堆: (j-i)%(k-1) == 0
func mergeStones(stones []int, k int) int {
	var n = len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	var sum = make([]int, n+1)
	for i, v := range stones {
		sum[i+1] = sum[i] + v
	}
	var dp = make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
		dp[i][i] = 0
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			for p := i; p < j; p += k - 1 {
				dp[i][j] = min(dp[i][j], dp[i][p]+dp[p+1][j])
			}
			if (j-i)%(k-1) == 0 { // 可以合并成一堆
				dp[i][j] += sum[j+1] - sum[i]
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
