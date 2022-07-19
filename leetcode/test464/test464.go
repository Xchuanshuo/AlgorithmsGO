package test464

func canIWin(max int, total int) bool {
	if max >= total {
		return true
	}
	if max*(max+1)/2 < total {
		return false
	}
	return canWin(max, total, make([]int, 1<<max), 0)
}

func canWin(max int, total int, dp []int, state int) bool {
	if dp[state] == 1 {
		return false
	}
	if dp[state] == 2 {
		return true
	}
	if total <= 0 {
		dp[state] = 1
		return false
	}
	for i := 1; i <= max; i++ {
		cur := 1 << (i - 1)
		if state&cur != 0 {
			continue
		}
		var win = canWin(max, total-i, dp, state|cur)
		if !win {
			dp[state] = 2
			return true
		}
	}
	dp[state] = 1
	return false
}
