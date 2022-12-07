package test1235

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/maximum-profit-in-job-scheduling/submissions/
 * idea: 按结束时间排序
 **/

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var n = len(startTime) + 1
	var arr = make([][]int, 0)
	arr = append(arr, []int{0, 0, 0})
	for i := 0; i < len(profit); i++ {
		arr = append(arr, []int{startTime[i], endTime[i], profit[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	var dp = make([]int, n)
	for i := 1; i < n; i++ {
		var j = sort.Search(n, func(k int) bool {
			return arr[k][1] > arr[i][0]
		})
		dp[i] = max(dp[i-1], dp[j-1]+arr[i][2])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
