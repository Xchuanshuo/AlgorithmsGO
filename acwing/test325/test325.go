package main

import (
	"fmt"
)

func main() {
	var n int
	for true {
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		var edges = make([][]int, 0)
		for i := 1; i < n; i++ {
			var id, w int
			fmt.Scanf("%d %d\n", &id, &w)
			edges = append(edges, []int{i, id - 1, w})
		}
		for _, w := range solve(n, edges) {
			fmt.Println(w)
		}
		n = -1
	}
}

func solve(n int, edges [][]int) []int {
	var g = make(map[int][]int)
	var w = make([]map[int]int, n)
	for i := 0; i < n; i++ {
		w[i] = make(map[int]int)
	}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		w[e[0]][e[1]] = e[2]
		w[e[1]][e[0]] = e[2]
	}
	var dis = make([]int, n)
	var dfs1 func(f, cur int) int
	dfs1 = func(fa, cur int) int {
		var res = 0
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			res = max(res, dfs1(cur, nxt)+w[cur][nxt])
		}
		dis[cur] = res
		return res
	}
	var dp = make([]int, n)
	var dfs2 func(fa, cur int, d int)
	dfs2 = func(fa, cur int, d int) {
		// 先序遍历，保存最新结果
		dp[cur] = d
		f, s := 0, 0
		for _, nxt := range g[cur] {
			var curW = dis[nxt] + w[cur][nxt]
			if curW > f {
				s, f = f, curW
			} else if curW > s {
				s = curW
			}
		}
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			// cur->nxt 换根具体计算逻辑
			if dis[nxt]+w[cur][nxt] == f {
				dis[cur] = s
			} else {
				dis[cur] = f
			}
			dfs2(cur, nxt, max(dis[nxt], dis[cur]+w[cur][nxt]))
		}
	}
	var sum = dfs1(-1, 0)
	dfs2(-1, 0, sum)
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
