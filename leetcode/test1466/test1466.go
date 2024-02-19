package test1466

func minReorder(n int, edges [][]int) int {
	var g = make(map[int][][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], []int{e[1], -1})
		g[e[1]] = append(g[e[1]], []int{e[0], 1})
	}
	var cnt = 0
	var dfs func(fa, cur int)
	dfs = func(fa, cur int) {
		for _, nxt := range g[cur] {
			if nxt[0] == fa {
				continue
			}
			if nxt[1] == -1 {
				cnt++
			}
			dfs(cur, nxt[0])
		}
	}
	dfs(-1, 0)
	return cnt
}
