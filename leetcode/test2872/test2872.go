package test2872

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-k-divisible-components
 * idea: 每访问到一条边，边两端的节点和能整除k，则为一个合法分割
 **/

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res = 0
	var dfs func(fa, cur int) int
	dfs = func(fa, cur int) int {
		var s = values[cur]
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			s = (s + dfs(cur, nxt)) % k
		}
		if s%k == 0 {
			res++
		}
		return s
	}
	dfs(-1, 0)
	return res
}
