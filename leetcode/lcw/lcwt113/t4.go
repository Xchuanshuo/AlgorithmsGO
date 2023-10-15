package lcwt113

/**
 * @Description https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/
 * idea: 换根dp. 建无向图，需要反转的边权重记为1, 无需反转记为0
			1.第一遍dfs选取任意一个点作为根节点计算最少反转次数
			2.第二遍dfs, 从1选取的根节点开始依次换根，计算当前次数
			 当前节点cur, 下一节点nxt, 换根后的权值为 sz[nxt]+(dp[cur]-sz[nxt]-w)+(w^1))
			 即下一节点nxt的反转次数 + 当前节点最小反转次数(需要去除nxt子树的次数) + nxt到cur的边权
 **/

func minEdgeReversals(n int, edges [][]int) []int {
	var g = make(map[int][][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], 0})
		g[e[1]] = append(g[e[1]], []int{e[0], 1})
	}
	var sz = make([]int, n)
	var dp = make([]int, n)
	var dfs1 func(fa, cur int) int
	dfs1 = func(fa, cur int) int {
		var res = 0
		for _, nxtA := range g[cur] {
			nxt, w := nxtA[0], nxtA[1]
			if nxt == fa {
				continue
			}
			res += dfs1(cur, nxt) + w
		}
		sz[cur] = res
		return res
	}
	var dfs2 func(fa, cur, sum int)
	dfs2 = func(fa, cur, sum int) {
		dp[cur] = sum
		for _, nxtA := range g[cur] {
			nxt, w := nxtA[0], nxtA[1]
			if nxt == fa {
				continue
			}
			// 换根逻辑 cur->nxt, sz[nxt] + (dp[cur]-sz[nxt]-w) + (w^1)
			dfs2(cur, nxt, sz[nxt]+(dp[cur]-sz[nxt]-w)+(w^1))
		}
	}
	var dp0 = dfs1(-1, 0)
	dfs2(-1, 0, dp0)
	return dp
}
