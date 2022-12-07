package lcw320

/**
 * @Description https://leetcode.cn/problems/number-of-beautiful-partitions/
 * idea: dp[i][j] 前i个字符 分成j段
		 dp[i][j] = sum(dp[0..i-l][j-1]) 枚举i，最终复杂度为O(N*N*K)
		 考虑前缀和优化，过程中累加dp[0..i][j-1]即可
		 过程中判断切分点是否合法，最终复杂度O(N*K)
 **/

// 1.边界是否过滤 // 2.初始条件是否考虑清楚
func beautifulPartitions(s string, k int, l int) int {
	n, M := len(s), int(1e9)+7
	var isPrim = func(a byte) bool {
		return a == '2' || a == '3' || a == '5' || a == '7'
	}
	var isValid = func(i int) bool {
		if i < 0 {
			return true
		}
		return (i == 0 && isPrim(s[i])) || (i == n-1 && !isPrim(s[i])) ||
			(!isPrim(s[i]) && isPrim(s[i+1]))
	}
	if !isValid(0) || !isValid(n-1) {
		return 0
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	if l == 1 {
		l += 1
	}
	dp[0][0] = 1
	for j := 1; j <= k; j++ {
		var sum = 0
		for i := 0; i <= n-l; i++ {
			if isValid(i - 1) {
				sum = (sum + dp[i][j-1]) % M
			}
			if isValid(i + l - 1) {
				dp[i+l][j] = sum
			}
		}
	}
	return dp[n][k]
}
