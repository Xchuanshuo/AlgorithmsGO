package test1639

// dp[i][j]表示target的前i个字符由words的前j个字符构建的方案数
// dp[i][j]= (dp[i-1][j-1] * cnts[i][v]) + dp[i][j-1]
// 即选第j个字符构成前i个字符的方案数 + 不选字符j构成前i个字符的方案数
// 初始条件: target的前0个字符由words的前j个字符构成的方案数为1

var M = int(1e9) + 7

func numWays(words []string, target string) int {
	var n = len(target)
	var m = len(words[0])
	var cnts = make([][26]int, m)
	for _, w := range words {
		for i, v := range w {
			var idx = int(v - 'a')
			cnts[i][idx]++
		}
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		var idx = int(target[i-1] - 'a')
		for j := i; j <= m; j++ {
			dp[i][j] = (dp[i-1][j-1]*cnts[j-1][idx] + dp[i][j-1]) % M
		}
	}
	return dp[n][m]
}
