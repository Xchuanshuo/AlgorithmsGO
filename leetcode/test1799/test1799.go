package test1799

/**
 * @Description https://leetcode.cn/problems/maximize-score-after-n-operations/
 * idea: 状压dp，1.枚举所有合法状态 2.对于当前状态，枚举可以由前面哪些状态转移过来
 **/

import "math/bits"

func maxScore(nums []int) int {
	var n = len(nums)
	var g = make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			g[i][j] = gcd(nums[i], nums[j])
		}
	}
	var m = 1 << n
	var dp = make([]int, m)
	for s := 0; s < m; s++ {
		var cnt = bits.OnesCount(uint(s))
		if cnt%2 != 0 {
			continue
		}
		for i := 0; i < n; i++ {
			if (1<<i)&s == 0 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if (1<<j)&s == 0 {
					continue
				}
				var pre = s ^ (1 << i) ^ (1 << j)
				dp[s] = max(dp[s], dp[pre]+cnt/2*g[i][j])
			}
		}
	}
	return dp[m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
