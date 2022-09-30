package test1420

/**
 * @Description https://leetcode.cn/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
 * idea: 动态规划 设dp[i][j][k]表示取i个数最大值为j花费为k的情况下生成数组的方案数
 *  求当前的方案数，只需看前i-1个数, 分为两种情况
 *     1.若当前数也为前i-1个数最大值, 当前数可以取[1..j] 此时方案数为 dp[i-1][j][k]*j
 *     2.若当前数不为前i-1个数最大值, 此时方案数为 dp[i-1][1..j-1][k-1]
 *  初始值: dp[1][1..m][1] = 1 取1个数, 最大值为1..m，花费为1时的方案数为1
 *   答案: dp[n][1..m][c] 即 取n个数，最大值为1..m, 花费为c时的方案数
 *  前缀和优化: 可以在里层枚举j，方便记录花费k时，1..j-1的前缀和
 *  整体复杂度: O(N*M*K)
 **/

const M = int(1e9) + 7

func numOfArrays(n int, m int, c int) int {
	var dp = make([][][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, c+1)
		}
	}
	for i := 1; i <= m; i++ {
		dp[1][i][1] = 1
	}
	for i := 1; i <= n; i++ {
		for k := 1; k <= c; k++ {
			var preSum = 0
			for j := 1; j <= m; j++ {
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k]*j) % M
				dp[i][j][k] = (dp[i][j][k] + preSum) % M
				preSum += dp[i-1][j][k-1]
			}
		}
	}
	var res = 0
	for i := 1; i <= m; i++ {
		res = (res + dp[n][i][c]) % M
	}
	return res
}
