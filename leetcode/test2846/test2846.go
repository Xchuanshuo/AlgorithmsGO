package test2846

/**
 * @Description https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree
 * idea: 1.求v、w最近公共祖先lca
		 2.v、w到lca的长度即为路径长度 dep[v]+dep[w]-2*dep[lca]
		 3.求出v,w到lca的所有路径中出现次数最大的边权
 **/

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	var g = make(map[int][][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], e[2] - 1})
		g[e[1]] = append(g[e[1]], []int{e[0], e[2] - 1})
	}
	var dp = make([][15]int, n) // 记录节点i的第2^j个节点
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var depth = make([]int, n) // 记录节点i的深度
	var sum = make([][15][26]int, n)
	// 1.计算深度与倍增数组
	var dfs func(fa, cur int)
	dfs = func(fa, cur int) {
		dp[cur][0] = fa
		for _, nxt := range g[cur] {
			if nxt[0] == fa {
				continue
			}
			sum[nxt[0]][0][nxt[1]] = 1
			depth[nxt[0]] = depth[cur] + 1
			dfs(cur, nxt[0])
		}
	}
	dfs(-1, 0)
	// 2.树上倍增 计算节点的kth祖先节点以及节点权值次数
	for j := 1; j < 15; j++ {
		for i := 0; i < n; i++ {
			if dp[i][j-1] == -1 {
				continue
			}
			dp[i][j] = dp[dp[i][j-1]][j-1]
			// 业务逻辑-------------------------------------
			for k := 0; k < 26; k++ {
				sum[i][j][k] = sum[i][j-1][k] + sum[dp[i][j-1]][j-1][k]
			}
			// 业务逻辑-------------------------------------
		}
	}
	// 3.计算lca
	var lca = func(v, w int) int {
		var pathLen = depth[v] + depth[w]
		var cnt = make([]int, 26)
		if depth[v] > depth[w] {
			v, w = w, v
		}
		var d = depth[w] - depth[v]
		// 将节点w提升到与v同一高度 即求w的第d个祖先节点
		for j := 0; j < 15; j++ {
			if d&(1<<j) == 0 {
				continue
			}
			// 业务逻辑-------------------------------------
			for k := 0; k < 26; k++ {
				cnt[k] += sum[w][j][k]
			}
			// 业务逻辑-------------------------------------
			w = dp[w][j]
		}
		var lca = v
		if w != v { // w == v, 则当前w即为最近公共祖先
			for j := 14; j >= 0; j-- {
				if dp[v][j] == dp[w][j] {
					continue
				}
				// 业务逻辑-------------------------------------
				for k := 0; k < 26; k++ {
					cnt[k] += sum[w][j][k] + sum[v][j][k]
				}
				// 业务逻辑-------------------------------------
				v, w = dp[v][j], dp[w][j]
			}
			for k := 0; k < 26; k++ {
				cnt[k] += sum[w][0][k] + sum[v][0][k]
			}
			lca = dp[v][0] // dp[v][0]即为lca
		}
		var mx = 0
		for k := 0; k < 26; k++ {
			mx = max(mx, cnt[k])
		}
		pathLen -= 2 * depth[lca]
		return pathLen - mx
	}
	var res = make([]int, 0)
	for _, q := range queries {
		res = append(res, lca(q[0], q[1]))
	}
	return res
}
