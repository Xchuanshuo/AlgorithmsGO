package test827

var count = 0
var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func largestIsland(grid [][]int) int {
	m, n, v := len(grid), len(grid[0]), 1
	var idCntMp = make(map[int]int)
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count = 0
				dfs(grid, i, j, -v)
				idCntMp[-v] = count
				if count > res {
					res = count
				}
				v++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			var cur = 1
			var used = make(map[int]bool)
			for k := 0; k < len(dirs); k++ {
				nx, ny := i+dirs[k][0], j+dirs[k][1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				var id = grid[nx][ny]
				if used[id] {
					continue
				}
				cur += idCntMp[id]
				used[id] = true
			}
			if cur > res {
				res = cur
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y, id int) {
	m, n := len(grid), len(grid[0])
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 1 {
		return
	}
	count++
	grid[x][y] = id
	dfs(grid, x+1, y, id)
	dfs(grid, x, y+1, id)
	dfs(grid, x-1, y, id)
	dfs(grid, x, y-1, id)
}
