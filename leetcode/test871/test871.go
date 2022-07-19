package test871

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	if len(stations) == 0 {
		if startFuel >= target {
			return 0
		}
		return -1
	}
	var n = len(stations)
	var dp = make([]int, n+1)
	dp[0] = startFuel
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			var cur = stations[i-1][0]
			if dp[j-1] >= cur {
				dp[j] = max(dp[j], dp[j-1]+stations[i-1][1])
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
