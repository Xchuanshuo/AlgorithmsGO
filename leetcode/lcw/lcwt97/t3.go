package lcwt97

func maximizeWin(pos []int, k int) int {
	var n = len(pos)
	var dp = make([]int, n+1)
	l, r := 0, 0
	var res = 0
	for ; r < n; r++ {
		for pos[r]-pos[l] > k {
			l++
		}
		res = max(res, r-l+1+dp[l])
		dp[r+1] = max(dp[r], r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
