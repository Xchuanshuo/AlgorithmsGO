package lwc_autox

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/contest/autox2023/problems/BjAFy9/
 * idea: 动态规划 + 二分查找左边界 根据当前天数 枚举所有票 算能覆盖到的最大天数的位置
		dp[j] = dp[i+t_cover(days)] + t_price
	    tips: 注意边界条件 max_price = 1e9, days = 1e5, 票价初始化为 1e9*1e6
 **/

func minCostToTravelOnDays(days []int, tickets [][]int) int64 {
	var n = len(days)
	var dp = make([]int64, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = int64(1e9 * 1e6)
	}
	for i := 0; i < n; i++ {
		for _, t := range tickets {
			var j = sort.Search(n, func(p int) bool {
				return days[p] >= days[i]+t[0]
			})
			dp[j] = min(dp[j], dp[i]+int64(t[1]))
		}
	}
	return dp[n]
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
