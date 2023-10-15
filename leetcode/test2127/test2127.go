package test2127

/**
 * @Description https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting
 * idea: 内向基环树 有向连通图中k个节点有k条边，所以该连通图一定存在环。不存在环则一定有员工不能坐在喜欢的人身边，不符合题意
		分情况讨论 1.若环大小=2, 说明两个人互相喜欢，此时喜欢这两个人的其它员工都可以加入，即把所有环大小为2的连通块相加
			     2.若环大小>2, 选环以外的任何路径都不能满足题意，此时只能选环里面的所有员工，即环大小为最多人数
 **/

func maximumInvitations(favorite []int) int {
	var n = len(favorite)
	var degree = make([]int, n)
	for _, f := range favorite {
		degree[f]++
	}
	// dp数组记录到每个点的最长路径
	var dp = make([]int, n)
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		dp[i] = 1
		if degree[i] == 0 {
			q = append(q, i)
		}
	}
	// 1.拓扑排序去除叶子节点
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		var nxt = favorite[cur]
		degree[nxt]--
		if degree[nxt] == 0 {
			q = append(q, nxt)
		}
		dp[nxt] = max(dp[nxt], dp[cur]+1)
	}
	// 2.找环 入度>0的节点为环的路口
	res, sum := 0, 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			continue
		}
		ringSize, cur := 0, i
		for degree[cur] > 0 {
			degree[cur]--
			cur = favorite[cur]
			ringSize++
		}
		if ringSize == 2 {
			sum += dp[i] + dp[favorite[i]]
			res = max(res, sum)
		} else {
			res = max(res, ringSize)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
