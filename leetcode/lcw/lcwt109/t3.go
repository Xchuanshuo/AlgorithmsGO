package lcwt109

func maxScore(nums []int, x int) int64 {
	var n = len(nums)
	var dp = make([][2]int64, n+1)
	if nums[0]%2 == 0 {
		dp[1][0] = int64(nums[0])
		dp[1][1] = int64(nums[0] - x)
	} else {
		dp[1][0] = int64(nums[0] - x)
		dp[1][1] = int64(nums[0])
	}
	for i := 2; i <= n; i++ {
		if nums[i-1]%2 == 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]-int64(x)) + int64(nums[i-1])
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-int64(x)) + int64(nums[i-1])
		}
	}

	return max(dp[n][0], dp[n][1])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
