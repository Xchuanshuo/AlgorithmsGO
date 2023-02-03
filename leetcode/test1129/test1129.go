package test1129

// 权值为1红, 2蓝
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	var g = make(map[int][][]int)
	var visited = make([][][]bool, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < len(visited[i]); j++ {
			visited[i][j] = make([]bool, 3)
		}
	}
	for _, r := range redEdges {
		if _, exist := g[r[0]]; !exist {
			g[r[0]] = make([][]int, 0)
		}
		g[r[0]] = append(g[r[0]], []int{r[1], 1})
	}
	for _, r := range blueEdges {
		if _, exist := g[r[0]]; !exist {
			g[r[0]] = make([][]int, 0)
		}
		g[r[0]] = append(g[r[0]], []int{r[1], 2})
	}
	var res = make([]int, n)
	for i := 1; i < n; i++ {
		res[i] = -1
	}
	var q = make([][]int, 0)
	q = append(q, []int{0, 0, 0})
	visited[0][0][0] = true
	var level = 0
	for len(q) > 0 {
		var sz = len(q)
		level++
		for i := 0; i < sz; i++ {
			var cur = q[0]
			q = q[1:]
			for _, v := range g[cur[1]] {
				if visited[cur[1]][v[0]][v[1]] || cur[2] == v[1] {
					continue
				}
				visited[cur[1]][v[0]][v[1]] = true
				q = append(q, []int{cur[1], v[0], v[1]})
				if res[v[0]] == -1 {
					res[v[0]] = level
				}
			}
		}
	}
	return res
}
