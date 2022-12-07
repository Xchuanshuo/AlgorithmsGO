package test1751

import "sort"

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/
 * idea: 类似题 t2008 t1235
 **/

func maxValue(events [][]int, k int) int {
	var n = len(events)
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	for j := 0; j < k; j++ {
		for i := 0; i < n; i++ {
			// 找到第一个大于等于给定值的元素，则它的右边为小于给定值的最大元素
			var idx = sort.Search(i, func(j int) bool {
				return events[j][1] >= events[i][0]
			})
			dp[i+1][j+1] = max(dp[i][j+1], dp[idx][j]+events[i][2])
		}
	}
	return dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
