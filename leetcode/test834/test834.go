package test834

/**
 * @Description https://leetcode.cn/problems/sum-of-distances-in-tree/
 * idea: (换根)树形dp, 两遍dfs, 用dp[i]表示节点i到其它节点的距离之和,
		sz[i]表示以某个节点为根的树节点个数，sum[i]表示以i为根的子树所有节点到i的距离之和
		 1.统计以任意节点为子树的节点个数, 计算某一点到其它点的距离之和
		 2.若v,w相邻, v,w拆成两颗子树
			dp[v] = sum[v] + sum[w] + sz[w]
			dp[w] = sum[w] + sum[v] + sz[v]
		  dp[v]-dp[w] = sz[w]-sz[v] => dp[v] = dp[w] + sz[w] - (n-sz[w])
											 = dp[w] + 2*sz[w] - n
 											 = dp[w] - 2*sz[v] + n
	参考: https://leetcode.cn/problems/sum-of-distances-in-tree/solution/c-by-p8vxdbjpmc-jbgk/
 **/

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var sz = make([]int, n)
	var dfs func(f, v int) int
	dfs = func(f, v int) int {
		sz[v] = 1
		var dis = 0
		for _, w := range g[v] {
			if w == f {
				continue
			}
			// w到所有子节点的距离 + 1*sz[w]
			dis += dfs(v, w) + sz[w]
			sz[v] += sz[w]
		}
		return dis
	}
	var dp = make([]int, n)
	var sum0 = dfs(-1, 0)
	var calc func(f, v, sum int)
	calc = func(f, v int, sum int) {
		dp[v] = sum
		for _, w := range g[v] {
			if w == f {
				continue
			}
			calc(v, w, sum-2*sz[w]+n)
		}
	}
	calc(-1, 0, sum0)
	return dp
}
