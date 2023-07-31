package lcwt109

import "math"

var M = int(1e9) + 7

func numberOfWays(n int, x int) int {
	var dp = make([]int, n+1)
	var a = make([]int, 0)
	for i := 1; i <= n; i++ {
		var v = int(math.Pow(float64(i), float64(x)))
		a = append(a, v)
	}
	dp[0] = 1
	for _, v := range a {
		for i := n; i >= 1; i-- {
			if i >= v {
				dp[i] = (dp[i] + dp[i-v]) % M
			}
		}
	}
	return dp[n]
}
