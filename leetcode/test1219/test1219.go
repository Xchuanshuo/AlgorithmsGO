package test1219

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var m, n, res int

func getMaximumGold(grid [][]int) int {
	m, n, res = len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				dfs(i, j, 0, grid)
			}
		}
	}
	return res
}

func dfs(x, y, gold int, grid [][]int) {
	val := grid[x][y]
	grid[x][y] = 0
	gold += val
	if gold > res {
		res = gold
	}
	for i := 0; i < 4; i++ {
		nx, ny := x+dirs[i][0], y+dirs[i][1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] <= 0 {
			continue
		}
		dfs(nx, ny, gold, grid)
	}
	grid[x][y] = val
}
