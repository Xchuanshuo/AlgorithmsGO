package test1986



func minSessions(tasks []int, sessionTime int) int {
	var n = len(tasks)
	var l = 1 << n
	var cost = make([]int, l)
	var dp = make([]int, l)
	for i := 0; i < l; i++ {
		var t = 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				t += tasks[j]
			}
		}
		cost[i] = t
		dp[i] = int(1e9)
	}
	dp[0] = 0
	for i := 0; i < l; i++ {
		var m = 0
		for j := 0; j < n; j++ {
			if i&(1<<j) == 0 {
				m |= 1 << j
			}
		}
		for s := m; s != 0; s = (s - 1) & m {
			if cost[s] <= sessionTime {
				dp[i|s] = min(dp[i|s], dp[i]+1)
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
