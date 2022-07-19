package test873

func lenLongestFibSubseq(arr []int) int {
	var n = len(arr)
	var dp = make([][]int, n)
	var mp = make(map[int]int)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 0)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 2
		}
		mp[arr[i]] = i
	}
	var res = 0
	for i := 2; i < n; i++ {
		for j := i - 1; j >= 1; j-- {
			var d = arr[i] - arr[j]
			if k, exist := mp[d]; exist {
				dp[i][j] = dp[j][k] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
