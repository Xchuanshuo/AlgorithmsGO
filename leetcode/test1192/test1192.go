package test1192

func criticalConnections(n int, connections [][]int) [][]int {
	var g = make(map[int][]int)
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], c[1])
		g[c[1]] = append(g[c[1]], c[0])
	}
	return findBridge(n, g)
}

var findBridge = func(n int, g map[int][]int) [][]int {
	// ord, low分别记录dfs顺序号，往后遍历过程中能访问到的最小序号
	ord, low := make([]int, n), make([]int, n)
	var res = make([][]int, 0)
	var visited = make([]bool, n)
	var pos = 0
	var dfs func(fa, cur int)
	dfs = func(fa, cur int) {
		visited[cur] = true
		ord[cur], low[cur] = pos, pos
		pos++
		for _, nxt := range g[cur] {
			if nxt == fa {
				continue
			}
			if !visited[nxt] {
				dfs(cur, nxt)
				if ord[cur] < low[nxt] {
					res = append(res, []int{cur, nxt})
				}
			}
			low[cur] = min(low[cur], low[nxt])
		}
	}
	dfs(-1, 0)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
