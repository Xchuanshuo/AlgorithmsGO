package lcw361

/**
 * @Description https://leetcode.cn/contest/weekly-contest-361/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/
 * idea: dfs + lca + 树上倍增
		1.最小操作次数 => 求两节点路径的边数 - 权值次数出现最多的边数
		2.求边: 设从任意一根节点出发，到点i的深度为depth[i]。对于节点v,w, 边数为 depth[v]+depth[w]-2*depth[lca(v,w)]
		3.求权值次数: dp[i][j]表示节点i的第2^j节点，sum[i][j][k]表示节点i的第2^j节点, 权值k的出现次数。求解lca的过程
					累加权值次数，最后对所有不同类型的权值累加即可
 **/

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	var g = make(map[int][][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], e[2] - 1})
		g[e[1]] = append(g[e[1]], []int{e[0], e[2] - 1})
	}
	var depth = make([]int, n)
	var dp = make([][15]int, n)
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var sum = make([][15][26]int, n)
	// 1.dfs求深度
	var dfs func(fa, cur int)
	dfs = func(fa, cur int) {
		dp[cur][0] = fa // 预处理
		for _, nxt := range g[cur] {
			if fa == nxt[0] {
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
					cnt[k] += sum[v][j][k] + sum[w][j][k]
				}
				// 业务逻辑-------------------------------------
				v, w = dp[v][j], dp[w][j]
			}
			for k := 0; k < 26; k++ { // 累加v, w到lca的权值次数
				cnt[k] += sum[v][0][k] + sum[w][0][k]
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
