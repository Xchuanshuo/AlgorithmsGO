package test2290

var dirs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
var INF = int(1e9)

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dis = make([][]int, m)
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = INF
		}
	}
	dis[0][0] = 0
	var q0 = make([][]int, 0)
	var q1 = make([][]int, 0)
	q0 = append(q0, []int{0, 0})
	for len(q0) > 0 || len(q1) > 0 {
		var x, y int
		if len(q0) > 0 {
			x, y = q0[0][0], q0[0][1]
			q0 = q0[1:]
		} else {
			x, y = q1[0][0], q1[0][1]
			q1 = q1[1:]
		}
		if x == m-1 && y == n-1 {
			return dis[m-1][n-1]
		}
		for i := 0; i < len(dirs); i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			var cost = 0
			if grid[nx][ny] == 1 {
				cost = 1
			}
			if dis[nx][ny] > dis[x][y]+cost {
				dis[nx][ny] = dis[x][y] + cost
				if cost == 1 {
					q1 = append(q1, []int{nx, ny})
				} else {
					q0 = append(q0, []int{nx, ny})
				}
			}
		}
	}
	return dis[m-1][n-1]
}
