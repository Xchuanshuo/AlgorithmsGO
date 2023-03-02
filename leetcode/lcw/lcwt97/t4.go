package lcwt97

var dirs = [][]int{{1, 0}, {0, 1}}

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var visited = make([]bool, m*n)
	var ord = make([]int, m*n)
	var low = make([]int, m*n)
	var cnt = 0
	var dfs func(v int, fa int)
	var cutP = 0
	dfs = func(v int, fa int) {
		x, y := v/n, v%n
		visited[v] = true
		ord[v], low[v] = cnt, cnt
		cnt += 1
		for i := 0; i < len(dirs); i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx >= m || ny >= n {
				continue
			}
			idx := nx*n + n
			if !visited[idx] {
				dfs(idx, v)
				low[v] = min(low[idx], low[v])
				if low[idx] > ord[v] {
					cutP++
				}
			} else if idx != fa {
				low[v] = min(low[idx], low[v])
			}
		}
	}
	return cutP <= 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
