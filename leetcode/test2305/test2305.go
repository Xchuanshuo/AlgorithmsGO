package test2305

// 当前孩子取sub个零食包 其它孩子取s^sub个零食包
func distributeCookies(cookies []int, k int) int {
	var n = len(cookies)
	var cal = func(s int) int {
		var total = 0
		for j := 0; j < n; j++ {
			if s&(1<<j) != 0 {
				total += cookies[j]
			}
		}
		return total
	}
	var l = 1 << n
	var dp = make([]int, l)
	for i := 1; i < len(dp); i++ {
		dp[i] = int(1e9)
	}
	for i := 0; i < k; i++ {
		for s := l - 1; s > 0; s-- {
			for sub := s; sub != 0; sub = (sub - 1) & s {
				dp[s] = min(dp[s], max(cal(sub), dp[s^sub]))
			}
		}
	}
	return dp[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
