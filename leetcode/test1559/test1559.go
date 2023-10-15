package test1559

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func containsCycle(g [][]byte) bool {
	m, n := len(g), len(g[0])
	var dfs func(fa, cur int) bool
	var vis = make([]bool, m*n)
	dfs = func(fa, cur int) bool {
		x, y := cur/n, cur%n
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || g[nx][ny] != g[x][y] {
				continue
			}
			var np = nx*n + ny
			if vis[np] && np != fa {
				return true
			}
			if vis[np] {
				continue
			}
			vis[np] = true
			if dfs(cur, np) {
				return true
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i*n+j] {
				continue
			}
			vis[i*n+j] = true
			if dfs(-1, i*n+j) {
				return true
			}
		}
	}
	return false
}
