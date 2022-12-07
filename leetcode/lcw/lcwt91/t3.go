package lcwt91

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var dfs1 func(v int)
	var parent = make(map[int]int)
	parent[bob] = -1
	var visited = make(map[int]bool)
	dfs1 = func(v int) {
		for _, nxt := range g[v] {
			if visited[nxt] {
				continue
			}
			visited[nxt] = true
			parent[nxt] = v
			if nxt == 0 {
				return
			}
			dfs1(nxt)
		}
	}
	visited[bob] = true
	dfs1(bob)
	var s = 0
	var nodes = make([]int, 0)
	for parent[s] != -1 {
		nodes = append(nodes, parent[s])
		s = parent[s]
	}
	var levelMap = make(map[int]int)
	for i := len(nodes) - 1; i >= 0; i-- {
		levelMap[nodes[i]] = len(nodes) - i - 1
	}
	visited = make(map[int]bool)
	var dfs2 func(v, level, score int)
	var res = -int(1e9)
	dfs2 = func(v, level int, score int) {
		var hasNxt = false
		for _, nxt := range g[v] {
			if visited[nxt] {
				continue
			}
			visited[nxt] = true
			hasNxt = true
			var nxtScore = amount[nxt]
			if preLevel, exist := levelMap[nxt]; exist {
				// fmt.Println(preLevel, nxtScore)
				// 同一层级 共享分数
				if preLevel == level {
					nxtScore /= 2
				} else if level > preLevel {
					nxtScore = 0
				}
			}
			dfs2(nxt, level+1, score+nxtScore)
		}
		if !hasNxt {
			res = max(res, score)
		}
	}
	visited[0] = true
	dfs2(0, 1, amount[0])
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
