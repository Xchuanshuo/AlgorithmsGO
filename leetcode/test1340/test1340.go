package test1340

/**
 * @Description https://leetcode.cn/problems/jump-game-v/
 * idea:	线性dp, 因为只能从高走向矮的位置, 所以将所有点从矮到高排序
			设dp[i]表示从点i出发能走的最大步数, 初始化dp[i]=1
			从小点开始依次察看左右边点j能否跳跃，能跳则当前点i出发可走点数为 max(dp[i], dp[j]+1)
 **/

import "sort"

func maxJumps(arr []int, d int) int {
	var n = len(arr)
	var a = make([][]int, 0)
	for i, v := range arr {
		a = append(a, []int{v, i})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	var dp = make([]int, n)
	var res = 1
	for i := 0; i < n; i++ {
		var idx = a[i][1]
		dp[idx] = 1
		for j := idx - 1; j >= idx-d && j >= 0; j-- {
			if a[i][0] <= arr[j] {
				break
			}
			dp[idx] = max(dp[idx], dp[j]+1)
		}
		for j := idx + 1; j <= idx+d && j < n; j++ {
			if a[i][0] <= arr[j] {
				break
			}
			dp[idx] = max(dp[idx], dp[j]+1)
		}
		res = max(res, dp[idx])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
