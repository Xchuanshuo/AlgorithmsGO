package test2538

/**
 * @Description https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum/
 * idea: 树形dp
		1.去掉任一节点 + 保留当前节点的子树的路径
		2.不去掉任何节点 + 去掉当前节点后的路径
 **/

func maxOutput(n int, edges [][]int, price []int) int64 {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res int64 = 0
	// dfs返回 不去掉根节点后的最大价值, 去掉一个后的最大价值
	var dfs func(x, fa int) (int, int)
	dfs = func(x, fa int) (int, int) {
		mx_s1, mx_s2 := price[x], 0
		for _, nxt := range g[x] {
			if nxt == fa {
				continue
			}
			s1, s2 := dfs(nxt, x)
			res = int64(max(int(res), s1+mx_s2))
			res = int64(max(int(res), s2+mx_s1))
			mx_s1 = max(mx_s1, s1+price[x])
			mx_s2 = max(mx_s2, s2+price[x])
		}
		return mx_s1, mx_s2
	}
	dfs(0, -1)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
