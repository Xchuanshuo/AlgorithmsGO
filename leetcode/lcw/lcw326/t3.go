package lcw326

import (
	"math"
	"strconv"
)

/**
 * @Description
 * idea: dp, 计算k可以存储的位数 dp[i] = dp[i-j] + 1
 **/

func minimumPartition(s string, k int) int {
	n, m := len(s), len(strconv.Itoa(k))
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := i - 1; j >= i-m && j >= 0; j-- {
			var sub = s[j:i]
			v, _ := strconv.Atoi(sub)
			if v > k {
				continue
			}
			dp[i] = min(dp[i], dp[j]+1)
		}
	}
	if dp[n] == math.MaxInt32 {
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
