package test1220


const (
	a = iota
	e
	i
	o
	u
)

func countVowelPermutation(n int) int {
	var dp = make([][5]int, n+1)
	for i := 0;i < 5;i++ {
		dp[1][i] = 1
	}
	res, MOD := 0, int(1e9)+7
	for k := 2;k <= n;k++ {
		dp[k][a] = (dp[k-1][u] + dp[k-1][e] + dp[k-1][i])%MOD
		dp[k][e] = (dp[k-1][a] + dp[k-1][i])%MOD
		dp[k][i] = (dp[k-1][e] + dp[k-1][o])%MOD
		dp[k][o] =  dp[k-1][i]
		dp[k][u] = (dp[k-1][i] + dp[k-1][o])%MOD
	}
	for i := 0;i < 5;i++ {
		res = (res + dp[n][i])%MOD
	}
	return res
}