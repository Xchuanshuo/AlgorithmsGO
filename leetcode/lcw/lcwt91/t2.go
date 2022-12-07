package lcwt91

func countGoodStrings(low int, high int, zero int, one int) int {
	var dp = make([][]int, high+1)
	for i := 0; i <= high; i++ {
		dp[i] = make([]int, 2)
	}
	var M = int(1e9) + 7
	if zero > one {
		zero, one = one, zero
	}
	dp[0][1] = 1
	dp[zero][0] = 1
	for i := zero; i <= high; i++ {
		dp[i][0] = (dp[i-zero][0] + dp[i-zero][1]) % M
		if i-one >= 0 {
			dp[i][1] = (dp[i-one][0] + dp[i-one][1]) % M
		}
	}
	var res = 0
	for i := low; i <= high; i++ {
		res = (res + dp[i][0]) % M
		res = (res + dp[i][1]) % M
	}
	return res
}
