package test1043

func maxSumAfterPartitioning(arr []int, K int) int {
	var n = len(arr)
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		var mx = arr[i-1]
		for j := i; j >= i-K && j >= 0; j-- {
			dp[i] = max(dp[i], dp[j]+mx*(i-j))
			if j-1 >= 0 {
				mx = max(mx, arr[j-1])
			}
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
