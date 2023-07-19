package lcwt108

import (
	"math"
	"strconv"
)

var valid = make(map[int64]bool)

func init() {
	for i := 0; i < 30; i++ {
		valid[int64(math.Pow(5.0, float64(i)))] = true
	}
}

var INF = int(1e9)

func minimumBeautifulSubstrings(s string) int {
	var n = len(s)
	var dp = make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			if s[j] == '0' {
				continue
			}
			val, _ := strconv.ParseInt(s[j-1:i], 2, 64)
			if valid[val] && dp[j] != INF {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	if dp[n] == INF {
		return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
