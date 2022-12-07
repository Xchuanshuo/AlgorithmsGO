package test813

import "math"

/**
 * @Description https://leetcode.cn/problems/largest-sum-of-averages/
 * idea: dp[i][j]表示前i个元素分成j段, 最大分数是多少
		 枚举切分点x, dp[i][j] = max(dp[i][j], dp[x-1][j-1] + avg([x..i])
		 初始条件,dp[0..i][1]=sum([0..i])/(i+1)表示前i个元素切分成1段的评价值
 **/

func largestSumOfAverages(nums []int, k int) float64 {
	var n = len(nums)
	var dp = make([][]float64, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, k+1)
	}
	var t = 0
	for i, v := range nums {
		t += v
		dp[i+1][1] = float64(t) / float64(i+1)
	}
	for j := 2; j <= k; j++ {
		for i := 1; i <= n; i++ {
			var sum = 0
			dp[i][j] = dp[i][j-1]
			for x := i; x >= 1; x-- {
				sum += nums[x-1]
				dp[i][j] = math.Max(dp[i][j], dp[x-1][j-1]+float64(sum)/float64(i-x+1))
			}
		}
	}
	return dp[n][k]
}
