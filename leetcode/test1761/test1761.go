package test1761

import "math"

func minTrioDegree(n int, edges [][]int) int {
	var g = make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	var degree = make([]int, n)
	for _, e := range edges {
		g[e[0]-1][e[1]-1] = true
		g[e[1]-1][e[0]-1] = true
		degree[e[0]-1]++
		degree[e[1]-1]++
	}
	var res = math.MaxInt32
	var dfs func(fa, cur, layer int)
	dfs = func(fa, cur, layer int) {
		for nxt := 0; nxt < n; nxt++ {
			if nxt == fa || !g[cur][nxt] {
				continue
			}
			if fa != -1 && g[nxt][fa] {
				var d = degree[fa] + degree[cur] + degree[nxt] - 6
				res = min(res, d)
			}
			if layer < 2 {
				dfs(cur, nxt, layer+1)
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(-1, i, 1)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
