package test1020

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt, total := 0, 0
	q := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				total++
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					grid[i][j] = 2
					cnt++
					q = append(q, []int{i, j})
				}
			}
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < len(dirs); i++ {
			nx, ny := cur[0]+dirs[i][0], cur[1]+dirs[i][1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != 1 {
				continue
			}
			cnt++
			grid[nx][ny] = 2
			q = append(q, []int{nx, ny})
		}
	}
	return total - cnt
}
