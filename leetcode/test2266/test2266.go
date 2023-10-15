package test2266

var M = int(1e9) + 7

func countTexts(s string) int {
	var n = len(s)
	var dp = make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		var ci = int(s[i-1] - '0')
		for j := i; j >= max(1, i-getCnt(ci)+1); j-- {
			var cj = int(s[j-1] - '0')
			if ci != cj {
				break
			}
			dp[i] = (dp[i] + dp[j-1]) % M
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getCnt(a int) int {
	if (a >= 2 && a <= 6) || a == 8 {
		return 3
	}
	return 4
}
