package test940

// dp[i] 以字符i结尾的子序列个数
func distinctSubseqII(s string) int {
	var dp = make([]int, 26)
	res, M := 0, int(1e9)+7
	for _, v := range s {
		dp[v-'a'] += 1
		for i := 0; i < len(dp); i++ {
			if i != int(v-'a') {
				dp[v-'a'] = (dp[v-'a'] + dp[i]) % M
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		res = (res + dp[i]) % M
	}
	return res
}
