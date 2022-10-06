package lcw309

import (
	"fmt"
)

/**
 * @Description https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
 * idea: 动态规划 设dp[i][j]表示结束位置为i，走j步情况下的方案总数
 *		 	1.状态转移方程: dp[i][j] = dp[i+1][j-1] + dp[i-1][j+1]
					   表示结束位置为i，是从i-1,i+1位置再走一步而来
 		 	2.初始条件: dp[start][0]=1, 表示走0步到达初始位置的方案总数为1
		 tips: a.由于每个位置可以从左或者右边的位置过来,所以对坐标进行一个整体偏移，防止坐标出现负数
		       b.枚举只需要从初始坐标左边k步到右边k步即可
 **/

var M = int(1e9) + 7
var base = 1001

func numberOfWays(startPos int, endPos int, k int) int {
	start, end := startPos+base, endPos+base
	var dp = make([][]int, 3010)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	dp[start][0] = 1
	for j := 1; j <= k; j++ {
		for i := start - k; i <= start+k; i++ {
			dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % M
			dp[i][j] = (dp[i][j] + dp[i+1][j-1]) % M
		}
	}
	return dp[end][k]
}

var mp map[string]int

func numberOfWays1(startPos int, endPos int, k int) int {
	mp = make(map[string]int)
	var dfs func(s, k int) int
	dfs = func(s, k int) int {
		var key = fmt.Sprintf("%d_%d", s, k)
		if v, exist := mp[key]; exist {
			return v
		}
		var res = 0
		if k <= 0 {
			if s == endPos {
				res = 1
			}
		} else {
			res = (dfs(s+1, k-1) + dfs(s-1, k-1)) % M
		}
		mp[key] = res
		return res
	}
	return dfs(startPos, k)
}
