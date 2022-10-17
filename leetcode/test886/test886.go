package test886

/**
 * @Description  https://leetcode.cn/problems/possible-bipartition/
 * idea: 图染色问题 二分图的检测
 **/

func possibleBipartition(n int, dislikes [][]int) bool {
	var visited = make([]int, n+1)
	for i := 0; i < len(visited); i++ {
		visited[i] = -1
	}
	var g = make(map[int][]int)
	for _, d := range dislikes {
		g[d[0]] = append(g[d[0]], d[1])
		g[d[1]] = append(g[d[1]], d[0])
	}
	var dfs func(cur, color int) bool
	dfs = func(cur, color int) bool {
		visited[cur] = color
		for _, nxt := range g[cur] {
			if visited[nxt] == -1 {
				if !dfs(nxt, getOther(color)) {
					return false
				}
			} else if visited[cur] == visited[nxt] {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if visited[i] == -1 && !dfs(i, 0) {
			return false
		}
	}
	return true
}

func getOther(color int) int {
	if color == 0 {
		return 1
	}
	return 0
}
