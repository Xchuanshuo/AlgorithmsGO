package test2218

/**
 * @Description https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles
 * idea: 分组背包 dp[i][j] 前i个栈取j个硬币，枚举当前取的硬币数l, dp[i-1][j-l] + sum(1, l)
 **/
func maxValueOfCoins(piles [][]int, k int) int {
	var n = len(piles)
	var sum = make([][]int, n)
	for i := 0; i < n; i++ {
		var m = len(piles[i])
		sum[i] = make([]int, m+1)
		for j, p := range piles[i] {
			sum[i][j+1] = sum[i][j] + p
		}
	}
	var dp = make([]int, k+1)
	for i := 0; i < n; i++ {
		var m = len(piles[i])
		for j := k; j >= 1; j-- {
			for l := 1; l <= m; l++ {
				if j-l < 0 {
					break
				}
				dp[j] = max(dp[j], dp[j-l]+sum[i][l])
			}
		}
	}
	return dp[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
