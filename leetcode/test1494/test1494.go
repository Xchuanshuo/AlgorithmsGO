package test1494

import "math/bits"

/**
 * @Description https://leetcode.cn/problems/parallel-courses-ii/
 * idea: 状压dp, 课程n最多为15, 可以使用状态压缩 关键点
		1.对于枚举到的每个选课状态, 要知道它的下一个可转移状态。需满足两个条件，
		一是选择未被选修的课程；二是未被选修的i课程中，它的前置课程都已经修完，这个可以提前预处理
		2.对于下一次可选修的课程，枚举1中得到的所有可修课程的非空子集，使用二进制枚举
		整体时间复杂度为(3^n)，空间为(2^n)
 **/

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	var pre = make([]int, n)
	for _, r := range relations {
		pre[r[1]-1] |= 1 << (r[0] - 1)
	}
	var l = 1 << n
	var dp = make([]int, l)
	for i := 1; i < len(dp); i++ {
		dp[i] = n
	}
	for s := 0; s < l; s++ {
		var m = 0
		for i := 0; i < n; i++ {
			// 课程i还未选修 且课程i前置课程已经修完
			if s&(1<<i) == 0 && pre[i]&s == pre[i] {
				m |= 1 << i
			}
		}
		// 枚举二进制状态m的所有非空子集
		for nxt := m; nxt != 0; nxt = (nxt - 1) & m {
			if bits.OnesCount(uint(nxt)) <= k {
				dp[s|nxt] = min(dp[s|nxt], dp[s]+1)
			}
		}
	}
	return dp[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	// 遍历所有状态
	return b
}
