package test1462

func checkIfPrerequisite(n int, ps [][]int, qs [][]int) []bool {
	var degree = make([]int, n)
	var g = make(map[int][]int)
	for _, p := range ps {
		g[p[0]] = append(g[p[0]], p[1])
		degree[p[1]]++
	}
	var q = make([]int, 0)
	var fa = make([][]bool, n)
	for i := 0; i < n; i++ {
		fa[i] = make([]bool, n)
		if degree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		for _, nxt := range g[cur] {
			fa[cur][nxt] = true
			for i := 0; i < n; i++ {
				fa[i][nxt] = fa[i][nxt] || fa[i][cur]
			}
			degree[nxt]--
			if degree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}
	var res = make([]bool, 0)
	for _, qv := range qs {
		res = append(res, fa[qv[0]][qv[1]])
	}
	return res
}
