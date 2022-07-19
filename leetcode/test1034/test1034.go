package test1034

// 1.找出联通分量并使用特殊颜色-1填充
// 2.将不满足边界的-1修改为-2
// 3.对于-1使用color填充，-2使用原来的颜色填充
func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	p := grid[row][col]
	dfs(grid, row, col, p)
	m, n := len(grid), len(grid[0])
	for i := 0;i < m;i++ {
		for j := 0;j < n;j++ {
			if grid[i][j] != -1 {
				continue
			}
			if !isValid(grid, m, n, i, j) {
				grid[i][j] = -2
			}
		}
	}
	for i := 0;i < m;i++ {
		for j := 0;j < n;j++ {
			if grid[i][j] == -1 {
				grid[i][j] = color
			} else if grid[i][j] == -2 {
				grid[i][j] = p
			}
		}
	}
	return grid
}

func isValid(grid [][]int, m, n, x, y int) bool {
	if x == 0 || y == 0 || x == m - 1 || y == n - 1 {
		return true
	}
	return grid[x+1][y] > 0 || grid[x][y+1] > 0 ||
		grid[x-1][y] > 0 || grid[x][y-1] > 0
}

func dfs(grid [][]int, x, y, last int) {
	if x < 0 || y < 0 || x >= len(grid) ||
		y >= len(grid[0]) || grid[x][y] == -1 || last != grid[x][y] {
		return
	}
	c := grid[x][y]
	grid[x][y] = -1
	dfs(grid, x + 1, y, c)
	dfs(grid, x, y + 1, c)
	dfs(grid, x - 1, y, c)
	dfs(grid, x, y - 1, c)
}
