package test410

import "math"

/**
 * @Description https://leetcode.cn/problems/split-array-largest-sum
 * idea: 动态规划 + 前缀和
		1.预处理区间前缀和，dp[i][k]表示前i个数切分成k个子数组的最大和的最小值.
		2.枚举区间[j..i], 则dp[i][k] = min(dp[i][k], max(dp[j][k-1], sum[j..i]))
 **/

func splitArray(nums []int, K int) int {
	var n = len(nums)
	var sum = make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for k := 1; k <= K; k++ {
		for i := 1; i <= n; i++ {
			for j := i; j >= k; j-- {
				dp[i][k] = min(dp[i][k], max(dp[j-1][k-1], sum[i]-sum[j-1]))
			}
		}
	}
	return dp[n][K]
}
