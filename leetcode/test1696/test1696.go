package test1696

/**
 * @Description https://leetcode.cn/problems/jump-game-vi/
 * idea: 线性dp + 队列
		1.dp[i]表示以i为终点时的最大得分
		2.dp[i] = max(dp[i-k]..dp[i-1]), 单调队列维护k大小窗口即可
 **/

func maxResult(nums []int, k int) int {
	var n = len(nums)
	var dp = make([]int, n)
	var q = make([]int, 0)
	dp[0] = nums[0]
	q = append(q, 0)
	for i := 1; i < n; i++ {
		if q[0]+k < i {
			q = q[1:]
		}
		dp[i] = dp[q[0]] + nums[i]
		for len(q) > 0 && dp[i] >= dp[q[len(q)-1]] {
			q = q[0 : len(q)-1]
		}
		q = append(q, i)
	}
	return dp[n-1]
}
