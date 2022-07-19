package test1706

var m, n int

func findBall(grid [][]int) []int {
	m, n = len(grid), len(grid[0])
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = dfs(0, i, grid)
	}
	return res
}

func dfs(x, y int, grid [][]int) int {
	if x == m {
		return y
	}
	if y == 0 && grid[x][y] == -1 {
		return -1
	}
	if y == n-1 && grid[x][y] == 1 {
		return -1
	}
	if grid[x][y] == 1 && y+1 < n && grid[x][y+1] != -1 {
		return dfs(x+1, y+1, grid)
	}
	if grid[x][y] == -1 && y-1 >= 0 && grid[x][y-1] != 1 {
		return dfs(x+1, y-1, grid)
	}
	return -1
}
