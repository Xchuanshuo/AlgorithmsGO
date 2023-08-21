package test2318

/**
 * @Description https://leetcode.cn/problems/number-of-distinct-roll-sequences
 * idea: dp[i][j] 表示前i个数以j结尾的方案数，枚举前面i-1的数k，满足k != j 且 gcd(j, k) = 1
		 由于abs(i-j) > 2, 所以需要过滤掉i-2位j的数 即 dp[i-1][k]-dp[i-2][j], 而实际
		 计算dp[i-1][k]时，i-3中已经过滤掉了和k相等的，dp[i-2][j]中包含，所以每次都多减去一部分
		 这一部分累计结果刚好为 dp[i-2][j]。
		注: 若直接加上dp[i-3][k], 这部分又包含了和[i-2][j]相等的，所以结果不对
 **/

var M = int(1e9) + 7

func distinctSequences(n int) int {
	var dp = make([][7]int, n+1)
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				if k == j || gcd(j, k) != 1 {
					continue
				}
				dp[i][j] = (dp[i][j] + dp[i-1][k] - dp[i-2][j] + M) % M
			}
			if i > 3 { // -dp[i-2][j]的过程中多减去了计算i-1时不包含的数，累计结果为dp[i-2][j]
				dp[i][j] = (dp[i][j] + dp[i-2][j]) % M
			}
		}
	}
	var res = 0
	for i := 1; i <= 6; i++ {
		res = (res + dp[n][i]) % M
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
