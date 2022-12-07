package lcw316

import (
	"math"
	"sort"
)

/**
 * @Description https://leetcode.cn/contest/weekly-contest-316/problems/minimum-cost-to-make-array-equal/
 * idea: 前缀和 公式转换即可  pre=(x-ai)*bi next=(ai-x)*bi
 **/

func minCost(nums []int, cost []int) int64 {
	var n = len(nums)
	var arr = make([][]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, []int{nums[i], cost[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	var sum = make([]int64, n+1)
	var mul = make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(arr[i][1])
		mul[i+1] = mul[i] + int64(arr[i][0])*int64(arr[i][1])
	}
	var res int64 = math.MaxInt64
	for i := 0; i < n; i++ {
		var pre = int64(arr[i][0])*sum[i] - mul[i]
		var next = mul[n] - mul[i] - int64(arr[i][0])*(sum[n]-sum[i])
		res = minInt64(res, pre+next)
	}
	return res
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
