package test837

// 设dp[i]表示当前分为i分时, 不超过n分的概率
// 1.当前分i为[k,n]时，肯定不超过n，概率为1
// 2.当前分>n时，概率为0
// 3.当前分数为[0,k-1], 此时概率为sum(dp[k..,k+2])/w
func new21Game(n int, k int, w int) float64 {
	var dp = make([]float64, n+k+w+1)
	for i := k; i <= n; i++ {
		dp[i] = 1.0
	}
	var total = 0.0
	for i := k; i < k+w; i++ {
		total += dp[i]
	}
	for i := k - 1; i >= 0; i-- {
		dp[i] = total / float64(w)
		total += dp[i] - dp[i+w]
	}
	return dp[0]
}
