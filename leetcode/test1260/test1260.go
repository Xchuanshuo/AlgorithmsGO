package test1260

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	var M = m * n
	k = k % M
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nx, ny := ((i*n+j-k+M)/n)%m, (i*n+j-k+M)%n
			res[i][j] = grid[nx][ny]
		}
	}
	return res
}
