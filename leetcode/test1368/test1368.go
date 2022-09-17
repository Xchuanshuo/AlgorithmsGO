package test1368

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var INF = int(1e9)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dis = make([][]int, m)
	for i := 0; i < m; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = INF
		}
	}
	dis[0][0] = 0
	var q = make([][]int, 0)
	q = append(q, []int{0, 0})
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for i := 0; i < len(dirs); i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			var cost = 0
			// 方向不同才需要额外花费
			if grid[x][y] != i+1 {
				cost = 1
			}
			// 到下一个路径花费更少才需要加入下一轮
			if dis[nx][ny] > dis[x][y]+cost {
				dis[nx][ny] = dis[x][y] + cost
				q = append(q, []int{nx, ny})
			}
		}
	}
	return dis[m-1][n-1]
}
